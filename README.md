# A Static Load Balancer written in Go
This project implements a static load balancer in **Go** based on the **Round Robin** load balancing algorithm. The project features **active health checks** to the three custom backends every 15 seconds, **health check request timeouts** and **skipping over offline backends.** In addition, the project also features **Prometheus** instrumentation and monitoring in custom **Grafana dashboards**.

# A short demo

The GIF below shows the load balancer routing to the three servers that are up. Next, *ServerOne* is turned off. I wait approximately 15 seconds until I am sure that a health check happened after the server went down, and the load balancer skips it over after it fails a health check and routes to servers 2 and 3 only.
<img width="800" height="416" alt="LoadBalancerGif" src="https://github.com/user-attachments/assets/3108b360-3294-41db-b3a8-31a2a968cdc4" />

**Note:** Refer to the [Future Works Section](#Future-Works) for the project caveats.

## Table of Contents


1. [The motivation behind the project](#The-motivation-behind-the-project)
2. [What I strived to achieve and learn through this project](#What-I-strived-to-achieve-and-learn-through-this-project)
3. [Project Features](#Project-Features)
4. [Tech Stack](#Tech-Stack)
5. [Project Architecture](#Project-Architecture)
6. [Project Structure](#Project-Structure)
7. [Grafana and Prometheus](#Grafana-and-Prometheus)
8. [Getting Started](#Getting-Started)
9. [Future Works](#Future-Works)

**Important: The terms "server" and "backend" are used interchangeably in this project.**

# The motivation behind the project
I wanted some more novelty and challenge in the down months between getting my bachelor's degree and starting my master's degree. I had the idea of writing my own load balancer for quite a while now, so I finally got to work. I decided to also use this project as a learning ground, and I picked Go as the programming language because it was created exactly for networking and infra-related software. I also wanted to have a portfolio project where I showcase some Prometheus and Grafana skills.

# What I strived to achieve and learn through this project
The goals of the project were to build a load balancer with the Round Robin algorithm and have most of the core load balancer functionalities preset, i.e. active health checks, skipping over down backends, request time outs, etc. I also wanted to pick up on some Go programming skills, which I believe I achieved. In addition, I wanted to dabble around with concurrency, which I found Go and this project to be perfect for. Finally, I wanted to practice my Prometheus and Grafana skills.

# Project Features
1. Active Health Checks and Health Check Cooldowns.
    - Health Checks for all backend are performed every 15 seconds. Upon a backend failing a health check five consecutive times, it is put on cooldown for 15 minutes until further health checks are performed on it. This saves computational power from going to backends that are down and are not likely to come back up within the next health check cycle(s).
2. Round Robin Routing
    - Round Robin is utilized to select the next server that the Load Balancer is routing to. The first request that the load balancer receives is sent to the first backend on the list of healthy backends, then the next request is routed to the next backend in line, etc.
3. Request Timeouts
    - If a health check endpoint response takes more than 5 seconds to arrive, the request is cancelled, and the function goes on without it. This ensures that the load balancer does not get stuck waiting for a response from an unresponsive backend.
4. Skipping Down Backends
    - Backends that have failed 5 consecutive health checks are counted as down and not included in the list of healthy backends until their cooldown of 15 minutes expires **and** they successfully pass a health check.
5. Atomic Variables and Monotonic Clocks
    - Atomic Variables are used as a way of avoiding race conditions since it is possible that more than one goroutine is trying to read and/or write to the same atomic variable at the same time.
    - Monotonic Clocks are used because they measure time regardless of the system time, thus, a change of system time could not affect the cooldown duration.
6. Mutexes
    - Mutexes were used where using atomic variables was either not possible or not the best solution when it came to avoiding race conditions and deadlocks.

# Tech Stack 
The tech stack consists of **Go**, **Prometheus**, **Grafana**, and **Docker Compose**. 

# Project Architecture
The project architecture consists of a reverse proxy that routes client requests to the specified backends and the backends themselves. The reverse proxy also functions as a load balancer because it calls to the functions that implement the load balancer functionality, i.e., FindNextServerIdx, SetServerCooldown, etc. 

The diagram below shows a high-level overview of the architecture. The dashed lines are intended to indicate backend that are down and failing health checks, whereas the solid lines indicate a successful health check and thus routing to the respective backend.

<img width="1185" height="710" alt="LoadBalancerArchitecture" src="https://github.com/user-attachments/assets/747f38df-e9d4-42ab-8c2b-83f32ca327fc" />

Apart from handling user requests and routing users to a specific backend, the load balancer also performs health checks every 15 seconds for every backend it is configured to route to. The health checks are at the `/health` endpoint. They either return a status code of `200`, indicating success, or they do not return anything, whereas a status code of `503` is assumed. 

There are two ways that a health check fails:
1. The health check HTTP request times out.
2. The backend that the health check is performed for is down.

The diagram below shows a high-level overview of how health checks happen:

<img width="779" height="703" alt="TimeOutDiagram" src="https://github.com/user-attachments/assets/a4205e36-ace5-4c10-8f72-03344947594a" />

It is important to note that I spent a lot of time preventing race conditions and deadlocks when handling concurrent operations and processes. I haven't caught either of those after debugging the source code and implementing various solutions such as mutexes and using atomic variables where applicable. If you run the project and notice such, however, **please let me know**. I am interested to see where and how they occur.

# Project Structure
The project consists of the endpoints, HelperFunctions, Grafana, Prometheus, Server1, Server2, Server3, and Structs directories. 

**The endpoints Directory** contains the function that returns the specified response upon a successful Health Check, otherwise status code 503 is returned.

**The GlobalVariables Directory** contains the file from which the global variables are "taken".

**The Grafana and Prometheus Directories** contain the usual Grafana and Prometheus configuration yml files.

**The HelperFunctions Directory** contains all of the functions that the Load Balancer needs to execute during runtime. I tried to modularize the Load Balancer file as much as I could, while also keeping everything scalable and maintainable.

**The LoadBalancer Directory** contains the Load Balancer source and the Dockerfile form which the docker image for the docker container is created.  

**The Servers 1, 2, and 3 directories** are all identical with the exception that each server returns a different message to the user, i.e., "Server 1", "Server 2", or "Server 3", depending on which backend the user is routed to.

**The structs folder** contains all of the structs that are referenced in the rest of the code.

The full tree structure is shown below:
```
├── docker-compose.yml
├── endpoints
│   └── healthcheck.go
├── GlobalVariables
│   └── GlobalVariables.go
├── go.mod
├── go.sum
├── Grafana
│   └── datasources.yml
├── HelperFunctions
│   ├── FindNextServerIdx.go
│   ├── HandleServerDownCounter.go
│   ├── InitializeServers.go
│   ├── LoadHealthyServersList.go
│   ├── ParseJSONResponse.go
│   ├── SetServerCooldown.go
│   ├── SetUpCustomPrometheusMetrics.go
│   └── UpdateServerHealthState.go
├── loadbalancer
│   ├── Dockerfile
│   └── loadbalancer.go
├── prometheus
│   └── prometheus.yml
├── server1
│   ├── Dockerfile
│   └── server1.go
├── server2
│   ├── Dockerfile
│   └── server2.go
├── server3
│   ├── Dockerfile
│   └── server3.go
└── Structs
    └── Structs.go
```
# Grafana and Prometheus
## Prometheus
The project makes use of 10 custom Prometheus metrics defined in the *PrometheusMetrics* struct below:
```
type PrometheusMetrics struct {
	TotalRequests prometheus.Counter
	LoadBalancerResponseLatency prometheus.Histogram
	BackendsCount prometheus.Gauge
	HealthyBackendsCount prometheus.Gauge
	BackendDowntimeDuration *prometheus.HistogramVec
	TotalHealthCheckFailures prometheus.Counter
	BackendHealthCheckFailures *prometheus.CounterVec
	BackendsOnCooldown prometheus.Gauge
	TotalSuccessfulHealthChecks prometheus.Counter
	BackendCooldownsCounter *prometheus.GaugeVec
}
```
You can find the Struct in the Structs directory.

## Grafana
The project employs a single Grafana dashboard that consist of 18 total panels:
```
        A p99 load balancer response latency time series panel
        A p95 load balancer response latency time series panel
        A p50 load balancer response latency time series panel
        Requests per second toward the load balancer
        Healthy Backends
        Total Backends
        Healthy / Total Backends Ratio
        Total Health Check Failures
        Backends on Cooldown
        Backend One Cooldown Counter
        Backend Two Cooldown Counter
        Backend Three Cooldown Counter
        Backend One Health Check Failures
        Backend Two Health Check Failures
        Backend Three Health Check Failures
        Backend One p99/95/50 Downtime Duration
        Backend Two p99/95/50 Downtime Duration
        Backend Three p99/95/50 Downtime Duration
```

I created those panels because they make it easy to triangulate issues across the project and pinpoint the cause faster than it would be to manually go through the code and debug. Even though I am hosting this project locally, it still helped me pinpoint an issue with the downtime duration that I did not catch in the code because the diagrams had clear issues that did not make any mathematical sense when looking at them, i.e., having a downtime of 40+ years, while the actual downtime was no more than a minute.

Screenshots from the Dashboard can be seen below.

<img width="1914" height="927" alt="Screenshot From 2026-08-03 08-26-49" src="https://github.com/user-attachments/assets/10ca61dd-0dbc-4abe-89dc-c2ec212e94f9" />
<img width="1914" height="927" alt="Screenshot From 2026-08-03 08-26-56" src="https://github.com/user-attachments/assets/310e2ac4-b08c-434c-9b7c-2466bb4ae104" />
<img width="1914" height="927" alt="Screenshot From 2026-08-03 08-27-00" src="https://github.com/user-attachments/assets/8ec719a7-3ccc-475a-a03d-483570e96d5b" />


# Getting Started

You only need to have Docker and Docker Compose installed. 

Once you have them installed, everything else will be taken care of automatically. The steps to running the project are as follows:

First, navigate to your terminal and clone the project on your machine:

```git clone https://github.com/MartinYordanov374/Golang-LoadBalancer.git```

Once the project is cloned, navigate to the repository:

```cd ./GoLang-LoadBalancer ```

Run the Docker Compose command:

```docker compose up --build```

You should have the project running on localhost now. Keep in mind that it takes about 30 seconds after having all containers running that the load balancer starts accepting and re-routing requests.

The Load Balancer is exposed on port 8000. 

The backend HTTP servers are on ports 8080, 8081, and 8082. 

Grafana is on port 3000, and Prometheus is on port 9090.
Note that I have not set up any custom credentials for Grafana, so you should be able to log in with the default credentials.

# Future Works
Currently, the project is quite simplistic. A couple of improvements that I thought of are:

1. Implement other routing algorithms, i.e. IP Hash or Weighted Round Robin
2. Implement Caching
   	- Currently, if a server goes down before a health check has established that it is down and a client visits the reverse proxy at port 8000, an error would be shown to the user, stating that the server could not be found. In such cases, I thought of caching the response of a generic endpoint from the server **or** having the load balancer inject some sort of a message to the user, i.e., *"The server seems to be down, please try refreshing this page."*
3. Use an actual backend rather than a simplistic HTTP response writer, i.e., a web application with an authentication system and a database. 
