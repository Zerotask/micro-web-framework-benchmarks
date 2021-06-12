use actix_web::{get, App, HttpRequest, HttpResponse, HttpServer, Responder};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Product {
    id: i32,
    name: String,
    price: f32,
}

fn add(value1: i32, value2: i32) -> i32 {
    value1 + value2
}

#[get("/")]
async fn index_endpoint(_req: HttpRequest) -> impl Responder {
    "Hello world"
}

#[get("/products")]
async fn products_endpoint(_req: HttpRequest) -> impl Responder {
    let products: [Product; 5] = [
        Product {
            id: 1,
            name: String::from("Fridge1"),
            price: 200.99,
        },
        Product {
            id: 2,
            name: String::from("Fridge2"),
            price: 300.99,
        },
        Product {
            id: 3,
            name: String::from("Fridge3"),
            price: 400.99,
        },
        Product {
            id: 4,
            name: String::from("Fridge4"),
            price: 500.99,
        },
        Product {
            id: 5,
            name: String::from("Fridge5"),
            price: 600.99,
        },
    ];
    HttpResponse::Ok().json(products)
}

#[get("/add")]
async fn add_endpoint(_req: HttpRequest) -> impl Responder {
    format!("{}", add(3, 8))
}

// http://127.0.0.1:13005
#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(index_endpoint)
            .service(add_endpoint)
            .service(products_endpoint)
    })
    .bind(("127.0.0.1", 13005))?
    .run()
    .await
}
