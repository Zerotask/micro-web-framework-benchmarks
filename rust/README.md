# Rust related benchmarks

## Actix (3.3.2)

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

## Rocket (0.5.0-rc.1)

URL: https://github.com/SergioBenitez/Rocket

TODO after 0.5.0 release has come out.
