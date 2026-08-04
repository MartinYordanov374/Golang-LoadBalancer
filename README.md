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

# Grafana and Prometheus

# Getting Started
