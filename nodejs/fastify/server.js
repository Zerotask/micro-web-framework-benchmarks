const fastify = require('fastify');
const app = fastify();

const add = (value1, value2) => value1 + value2;

app.get('/', (request, reply) => {
  reply.send('Hello world')
});

app.get('/products', function (request, reply) {
  const products = [
    {id: 1, name: "Fridge1", price:200.99},
    {id: 2, name: "Fridge2", price:300.99},
    {id: 3, name: "Fridge3", price:400.99},
    {id: 4, name: "Fridge4", price:500.99},
    {id: 5, name: "Fridge5", price:600.99},
  ]
  reply.send(products);
});

app.get('/add', function (request, reply) {
  reply.send(add(3, 8));
});

// http://127.0.0.1:13001
app.listen(13001, (err, address) => {
  if (err) throw err
  console.log('Server is up and running on http://127.0.0.1:13001');
});