# Micro web framework benchmarks

## Express (node.js)

### rewrk -c 100 -t 10 -d 30s -h http://127.0.0.1:13000

```
Benchmarking 100 connections @ http://127.0.0.1:13000 for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    134.80ms  17.28ms  3.00ms   246.95ms
  Requests:
    Total:  22256  Req/Sec: 740.82
  Transfer:
    Total: 5.05 MB Transfer Rate: 172.18 KB/Sec
```

### rewrk -c 200 -t 10 -d 30s -h http://127.0.0.1:13000/products

```
Benchmarking 200 connections @ http://127.0.0.1:13000/products for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    299.83ms  26.73ms  99.05ms  576.91ms
  Requests:
    Total:  20112  Req/Sec: 666.41
  Transfer:
    Total: 8.48 MB Transfer Rate: 287.65 KB/Sec
```

### rewrk -c 500 -t 10 -d 30s -h http://127.0.0.1:13000/add

```
Benchmarking 500 connections @ http://127.0.0.1:13000/add for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    684.23ms  358.57ms  43.11ms  6735.71ms
  Requests:
    Total:  22046  Req/Sec: 727.70
  Transfer:
    Total: 4.94 MB Transfer Rate: 167.00 KB/Sec
```

## Fastify (node.js)

### rewrk -c 100 -t 10 -d 30s -h http://127.0.0.1:13001

```
Benchmarking 100 connections @ http://127.0.0.1:13001 for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    119.65ms  13.71ms  76.25ms  558.43ms
  Requests:
    Total:  25104  Req/Sec: 835.03
  Transfer:
    Total: 4.19 MB Transfer Rate: 142.71 KB/Sec
```

### rewrk -c 200 -t 10 -d 30s -h http://127.0.0.1:13001/products

```
Benchmarking 200 connections @ http://127.0.0.1:13001/products for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    291.30ms  81.24ms  100.09ms  1797.23ms
  Requests:
    Total:  20661  Req/Sec: 685.84
  Transfer:
    Total: 7.43 MB Transfer Rate: 252.50 KB/Sec
```

### rewrk -c 500 -t 10 -d 30s -h http://127.0.0.1:13001/add

```
Benchmarking 500 connections @ http://127.0.0.1:13001/add for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    677.81ms  513.85ms  45.79ms  7835.59ms
  Requests:
    Total:  22294  Req/Sec: 736.48
  Transfer:
    Total: 3.64 MB Transfer Rate: 122.99 KB/Sec
```

## Polka (node.js)

`rewrk` is returning 404 errors at the moment. No tests were possible.

## Go stdlb (go)

### rewrk -c 100 -t 10 -d 30s -h http://127.0.0.1:13003

```
Benchmarking 100 connections @ http://127.0.0.1:13003 for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    114.19ms  19.80ms  7.65ms   426.32ms
  Requests:
    Total:  26278  Req/Sec: 874.42
  Transfer:
    Total: 3.21 MB Transfer Rate: 109.30 KB/Sec
```

### rewrk -c 200 -t 10 -d 30s -h http://127.0.0.1:13003/products

```
Benchmarking 200 connections @ http://127.0.0.1:13003/products for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    268.54ms  29.82ms  65.78ms  457.79ms
  Requests:
    Total:  22424  Req/Sec: 744.40
  Transfer:
    Total: 6.93 MB Transfer Rate: 235.53 KB/Sec
```

### rewrk -c 500 -t 10 -d 30s -h http://127.0.0.1:13003/add

```
Benchmarking 500 connections @ http://127.0.0.1:13003/add for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    597.83ms  156.53ms  9.30ms   3704.70ms
  Requests:
    Total:  24969  Req/Sec: 824.96
  Transfer:
    Total: 2.81 MB Transfer Rate: 95.06 KB/Sec
```
