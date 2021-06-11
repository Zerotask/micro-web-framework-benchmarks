const express = require('express');
const app = express();

const add = (value1, value2) => value1 + value2;

app.get('/', function (req, res) {
  res.send('Hello World');
});

app.get('/products', function (req, res) {
  const products = [
    {id: 1, name: "Fridge1", price:200.99},
    {id: 2, name: "Fridge2", price:300.99},
    {id: 3, name: "Fridge3", price:400.99},
    {id: 4, name: "Fridge4", price:500.99},
    {id: 5, name: "Fridge5", price:600.99},
  ]
  res.send(products);
});

app.get('/add', function (req, res) {
  res.json(add(3, 8));
});

// http://127.0.0.1:13000
app.listen(13000, () => console.log('Server is up and running on http://127.0.0.1:13000'));