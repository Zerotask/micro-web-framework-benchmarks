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
