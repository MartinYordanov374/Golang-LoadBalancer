package main

import (
	"bufio"
	_ "embed"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"golang-loadbalancer/GlobalVariables"
	"golang-loadbalancer/HelperFunctions"
	"golang-loadbalancer/Structs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Every decision of the load balancer is made in loadbalancer.bf. This file only lends it a network,
// a ticker, a clock and the Prometheus metrics, and runs the commands it prints.
//
//go:embed loadbalancer.bf
var Program string

type LoadBalancer struct {
	Mutex sync.Mutex
	Stdin io.Writer
	Stdout *bufio.Reader
	Metrics *Structs.PrometheusMetrics
	Timers map[string]time.Time
}

func main(){

	PrometheusRegistry := prometheus.NewRegistry()
	CustomMetrics := HelperFunctions.SetUpCustomMetrics(PrometheusRegistry)
	http.Handle("/metrics", promhttp.HandlerFor(PrometheusRegistry, promhttp.HandlerOpts{}))

	StdinReader, StdinWriter := io.Pipe()
	StdoutReader, StdoutWriter := io.Pipe()
	go HelperFunctions.RunBrainfuck(Program, StdinReader, StdoutWriter)
	LB := &LoadBalancer{Stdin: StdinWriter, Stdout: bufio.NewReader(StdoutReader), Metrics: CustomMetrics, Timers: map[string]time.Time{}}
	LB.Ask("")

	go func() {
		OriginURL, Error := url.Parse("http://127.0.0.1:0")
		if Error != nil{
			log.Println(Error)
		}
		ReverseProxy := httputil.NewSingleHostReverseProxy(OriginURL)
		OriginalDirector := ReverseProxy.Director
		ReverseProxy.Director = func(req *http.Request){
			OriginalDirector(req)
			for _, Reply := range LB.Ask("r") {
				Command, Argument, _ := strings.Cut(Reply, " ")
				switch Command {
				case "proxy":
					NextServer, _ := url.Parse(Argument)
					req.URL.Scheme = NextServer.Scheme
					req.URL.Host = NextServer.Host
					req.Host = req.URL.Host
				case "panic":
					panic(Argument)
				}
			}
		}
		http.Handle("/", ReverseProxy)
		http.ListenAndServe(":8000", nil)
	}()

	ticker := time.NewTicker(15 * time.Second)
	for range ticker.C {
		HealthChecks := LB.Ask("t")
		Results := make([]string, len(HealthChecks))
		wg := new(sync.WaitGroup)
		wg.Add(len(HealthChecks))
		for Idx, HealthCheck := range HealthChecks {
			go SendHealthCheck(strings.TrimPrefix(HealthCheck, "get "), &Results[Idx], wg)
		}
		wg.Wait()
		LB.Ask("h" + strings.Join(Results, ""))
	}
}

// SendHealthCheck writes down what happened to one health check request as a line for the Brainfuck program:
// ! if the request failed, ? and the error if the body could not be read, otherwise the StatusCode in the JSON body.
func SendHealthCheck(uri string, Result *string, WaitGroup *sync.WaitGroup){
	defer WaitGroup.Done()
	resp, err := GlobalVariables.HttpClient.Get(uri)
	if err != nil {
		*Result = "!\n"
		return
	}
	defer resp.Body.Close()
	jsonbody, err := io.ReadAll(resp.Body)
	if err != nil {
		*Result = "?" + err.Error() + "\n"
		return
	}
	ParsedJSON := HelperFunctions.ParseJSONResponse(jsonbody)
	*Result = strconv.Itoa(ParsedJSON.StatusCode) + "\n"
}

// Ask sends one event to the Brainfuck program and reads its answer up to the empty line.
// Metric, timer and log commands are run right away, the others are returned to the caller.
func (LB *LoadBalancer) Ask(Event string) []string {
	LB.Mutex.Lock()
	defer LB.Mutex.Unlock()
	if Event != "" {
		// written from a goroutine so that a long event never waits on the answer being read
		go io.WriteString(LB.Stdin, Event)
	}
	Replies := []string{}
	for {
		Line, _ := LB.Stdout.ReadString('\n')
		Line = strings.TrimSuffix(Line, "\n")
		if Line == "" {
			return Replies
		}
		Command, Argument, _ := strings.Cut(Line, " ")
		Fields := strings.Fields(Argument)
		switch Command {
		case "inc":
			LB.Metric(Fields).(interface{ Inc() }).Inc()
		case "dec":
			LB.Metric(Fields).(prometheus.Gauge).Dec()
		case "set":
			Value, _ := strconv.ParseFloat(Fields[1], 64)
			LB.Metric(Fields[:1]).(prometheus.Gauge).Set(Value)
		case "start":
			LB.Timers[Argument] = time.Now()
		case "observe":
			LB.Metric(Fields).(prometheus.Observer).Observe(time.Since(LB.Timers[Argument]).Seconds())
		case "log":
			log.Println(Argument)
		default:
			Replies = append(Replies, Line)
		}
	}
}

// Metric looks up the field of Structs.PrometheusMetrics that the Brainfuck program named, picking its label if it gave one.
func (LB *LoadBalancer) Metric(Name []string) any {
	Metric := reflect.ValueOf(LB.Metrics).Elem().FieldByName(Name[0])
	if len(Name) > 1 {
		Metric = Metric.MethodByName("WithLabelValues").Call([]reflect.Value{reflect.ValueOf(Name[1])})[0]
	}
	return Metric.Interface()
}
