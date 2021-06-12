# Micro web framework benchmarks

Micro web framework benchmarks created with [rewrk](https://github.com/ChillFish8/rewrk).

The benchmarks are located in language specific folders.

## Explanation

Every test consists of 3 benchmarks:

- a simple text response
- a simple JSON response
- a simple function call to add 2 numbers

## Run tests

Before you run your tests, ensure that the web servers are running.

the root package.json provides some scripts to start them:

- `npm run start:rust:actix`
- `npm run start:nodejs:express`
- `npm run start:nodejs:fastify`

## Alternative HTTP Benchmark tools
- [ali](https://github.com/nakabonne/ali) - graphical tool written in Go
- [bombardier](https://github.com/codesenberg/bombardier) - written in Go
- [cassowary](https://github.com/rogerwelin/cassowary) - written in Go
- [hey](https://github.com/rakyll/hey) - written in Go
- [httpit](https://github.com/gonetx/httpit) - written in Go
- [oha](https://github.com/hatoo/oha) - some graphic, written in Rust
- [vegeta](https://github.com/tsenart/vegeta) - graphical options, written in Go