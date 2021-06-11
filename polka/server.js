const polka = require('polka');
const app = polka();

const add = (value1, value2) => value1 + value2;

app.get('/', (req, res) => {
  res.end('Hello world!');
});

app.get('/products', function (req, res) {
  const products = [
    {id: 1, name: "Fridge1", price:200.99},
    {id: 2, name: "Fridge2", price:300.99},
    {id: 3, name: "Fridge3", price:400.99},
    {id: 4, name: "Fridge4", price:500.99},
    {id: 5, name: "Fridge5", price:600.99},
  ];
  const jsonResponse = JSON.stringify(products);
  res.end(jsonResponse);
});

app.get('/add', function (req, res) {
    const jsonResponse = JSON.stringify(add(3, 8));
  res.end(jsonResponse);
});

app.listen(13002, () => console.log('Server is up and running on http://127.0.0.1:13002'));