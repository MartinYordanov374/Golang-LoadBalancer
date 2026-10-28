# A Static Load Balancer written in Brainfuck

The Round Robin load balancer, rewritten in Brainfuck.

`loadbalancer/loadbalancer.bf` holds the backend list, the health state of every backend (down counter, 15 minute cooldown, downtime), the list of healthy backends and the round robin index. It decides which backends get a health check, where every request goes and which metric moves. `endpoints/server.bf` is the three backends: `Server N` on every path, the JSON health check on `/health`.

Brainfuck has no sockets, so Go stays on as the bridge. `loadbalancer/loadbalancer.go` sends the program one byte per event (`r` request, `t` ticker, `h` plus the health check results) and runs the commands it prints back (`proxy`, `get`, `inc`, `set`, `log`, ...). `endpoints/server.go` pipes the request path in and the response out. Both use the interpreter in `HelperFunctions/RunBrainfuck.go`.

## Running

```
docker compose up --build
```

As before, the load balancer is on port 8000, the backends on 8080 to 8082, Prometheus on 9090 and Grafana on 3000.

Any 8-bit Brainfuck interpreter runs the programs in a terminal too:

```
printf '2/health' | bf endpoints/server.bf
printf 'th200\n!\n200\nrrr' | bf loadbalancer/loadbalancer.bf
```
