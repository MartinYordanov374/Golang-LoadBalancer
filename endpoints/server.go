package endpoints

import (
	"bufio"
	_ "embed"
	"io"
	"net/http"
	"net/textproto"
	"strings"
	"golang-loadbalancer/HelperFunctions"
)

//go:embed server.bf
var Program string

// Server answers every request with what server.bf prints for this server number and request path:
// the headers, an empty line, then the body.
func Server(Number string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		Output := new(strings.Builder)
		HelperFunctions.RunBrainfuck(Program, strings.NewReader(Number+r.URL.Path), Output)
		Response := textproto.NewReader(bufio.NewReader(strings.NewReader(Output.String())))
		Header, _ := Response.ReadMIMEHeader()
		for Key, Values := range Header {
			w.Header()[Key] = Values
		}
		io.Copy(w, Response.R)
	}
}
