# A Static Load Balancer written in Go
This project implements a static load balancer in **Go** based on the **Round Robin** load balancing algorithm. The project features **active health checks** to the three custom backends every 15 seconds, **health check request timeouts** and **skipping over offline backends.** In addition, the project also features **Prometheus** instrumentation and monitoring in custom **Grafana dashboards**.
## Table of Contents


1. [The motivation behind the project](#The-motivation-behind-the-project)
2. [What I strived to achieve and learn through this project](#What-I-strived-to-achieve-and-learn-through-this-project)
3. [Project Features](#Project-Features)
4. [Tech Stack](#Tech-Stack)
5. [Project Architecture](#Project-Architecture)
6. [Project Structure](#Project-Structure)
7. [Grafana and Prometheus](#Grafana-and-Prometheus)
8. [Getting Started](#Getting-Started)


# The motivation behind the project
I wanted some more novelty and challenge in the down months between getting my bachelor's degree and starting my master's degree. I had the idea of writing my own load balancer for quite a while now, so I finally got to work. I decided to also use this project as a learning ground, and I picked Go as the programming language because it was created exactly for networking and infra-related software. I also wanted to have a portfolio project where I showcase some Prometheus and Grafana skills.

# What I strived to achieve and learn through this project
The goals of the project were to build a loadbalancer with the Round Robin algorithm and have most of the core load balancer functionalities preset, i.e. active health checks, skipping over down backends, request time outs, etc. I also wanted to pick up on some Go programming skills, which I believe I achieved. In addition, I wanted to dabble around with concurrency, which I found Go and this project to be perfect for. Finally, I wanted to practice my Prometheus and Grafana skills.

# Project Features
1. Active Health Checks  
2. Round Robin Routing
3. Request Timeouts
4. Health Checks Cooldown
5. Skipping Down Backends
6. Atomic variables and operations
7. Mutexes

# Tech Stack 
The tech stack consists of **Go**, **Prometheus**, **Grafana**, and **Docker Compose**. 

# Project Architecture

# Project Structure
The project consists of the endpoints, HelperFunctions, Grafana, Prometheus, Server1, Server2, Server3, and Structs directories. 

**The endpoints Directory** contains the function that returns the specified response upon a successful Health Check, otherwise status code 503 is returned.

**The GlobalVariables Directory** contains the file from which the global variables are "taken".

**The Grafana and Prometheus Directories** contain the usual Grafana and Prometheus configuration yml files.

**The HelperFunctions Directory** contains all of the functions that the Load Balancer needs to execute during runtime. I tried to modularize the Load Balancer file as much as I could, while also keeping everything scalable and maintainable.

**The LoadBalancer Directory** contains the Load Balancer source and the Dockerfile form which the docker image for the docker container is created.  

**The Servers 1, 2, and 3 directories** are all identical with the exception that each server returns a different message to the user, i.e., "Server 1", "Server 2", or "Server 3", depending on which server the user is routed to.

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

# Getting Started
