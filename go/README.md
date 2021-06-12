# Go related benchmarks

## stdlib

URL: https://golang.org/pkg/net/http/

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

## Fiber

URL: https://github.com/gofiber/fiber

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
