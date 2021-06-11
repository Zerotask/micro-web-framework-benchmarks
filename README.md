# Micro web framework benchmarks

Micro web framework benchmarks created with [rewrk](https://github.com/ChillFish8/rewrk).

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

## Go stdlib (go)

### rewrk -c 100 -t 10 -d 30s -h http://127.0.0.1:13003

```
Benchmarking 100 connections @ http://127.0.0.1:13003 for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    111.43ms  47.04ms  2.48ms   1166.70ms
  Requests:
    Total:  26958  Req/Sec: 896.95
  Transfer:
    Total: 3.29 MB Transfer Rate: 112.12 KB/Sec
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

## Fiber (go)

### rewrk -c 100 -t 10 -d 30s -h http://127.0.0.1:13004

```
Benchmarking 100 connections @ http://127.0.0.1:13004 for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    119.56ms  29.55ms  3.46ms   629.52ms
  Requests:
    Total:  25106  Req/Sec: 835.25
  Transfer:
    Total: 3.11 MB Transfer Rate: 106.04 KB/Sec
```

### rewrk -c 200 -t 10 -d 30s -h http://127.0.0.1:13004/products

```
Benchmarking 200 connections @ http://127.0.0.1:13004/products for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    278.27ms  160.02ms  6.39ms   2314.31ms
  Requests:
    Total:  21618  Req/Sec: 717.54
  Transfer:
    Total: 6.49 MB Transfer Rate: 220.73 KB/Sec
```

### rewrk -c 500 -t 10 -d 30s -h http://127.0.0.1:13004/add

```
Benchmarking 500 connections @ http://127.0.0.1:13004/add for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    609.23ms  96.70ms  14.36ms  1283.02ms
  Requests:
    Total:  24669  Req/Sec: 811.27
  Transfer:
    Total: 2.56 MB Transfer Rate: 86.36 KB/Sec
```

## Actix (rust)

### rewrk -c 100 -t 10 -d 30s -h http://127.0.0.1:13005

#### Debug

```
Benchmarking 100 connections @ http://127.0.0.1:13005 for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    134.57ms  30.42ms  12.60ms  661.99ms
  Requests:
    Total:  22315  Req/Sec: 742.29
  Transfer:
    Total: 2.72 MB Transfer Rate: 92.79 KB/Sec
```

#### Production

```
Benchmarking 100 connections @ http://127.0.0.1:13005 for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    118.11ms  14.91ms  4.79ms   426.10ms
  Requests:
    Total:  25397  Req/Sec: 844.84
  Transfer:
    Total: 3.10 MB Transfer Rate: 105.61 KB/Sec
```

### rewrk -c 200 -t 10 -d 30s -h http://127.0.0.1:13005/products

#### Debug

```
Benchmarking 200 connections @ http://127.0.0.1:13005/products for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    330.60ms  37.53ms  163.78ms  842.58ms
  Requests:
    Total:  18227  Req/Sec: 604.35
  Transfer:
    Total: 5.48 MB Transfer Rate: 185.91 KB/Sec
```

#### Production

```
Benchmarking 200 connections @ http://127.0.0.1:13005/products for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    300.00ms  33.55ms  99.26ms  483.09ms
  Requests:
    Total:  20084  Req/Sec: 666.21
  Transfer:
    Total: 6.03 MB Transfer Rate: 204.94 KB/Sec
```

### rewrk -c 500 -t 10 -d 30s -h http://127.0.0.1:13005/add

#### Debug

```
Benchmarking 500 connections @ http://127.0.0.1:13005/add for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    742.27ms  350.27ms  74.63ms  5293.28ms
  Requests:
    Total:  19312  Req/Sec: 635.60
  Transfer:
    Total: 2.17 MB Transfer Rate: 73.24 KB/Sec
```

#### Production

```
Benchmarking 500 connections @ http://127.0.0.1:13005/add for 30 seconds
  Latencies:
    Avg      Stdev    Min      Max
    675.64ms  292.69ms  144.44ms  5462.66ms
  Requests:
    Total:  21544  Req/Sec: 710.71
  Transfer:
    Total: 2.42 MB Transfer Rate: 81.90 KB/Sec
```
