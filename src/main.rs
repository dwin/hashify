use actix_web::{web, App, HttpServer};
use crate::controller::{get_index, get_status, list_methods, keygen, compute_hash, hash_highway, hash_highway64, hash_highway128};

mod controller;
mod hash;
mod highway;
mod keygen;
mod methods;
mod status;
mod tests;

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .route("/", web::get().to(get_index))
            .route("/status", web::get().to(get_status))
            .route("/methods", web::get().to(list_methods))
            .route("/keygen/{length}", web::get().to(keygen))
            .service(
                web::scope("/hash")
                    .route("/{algo}/{format}", web::get().to(compute_hash))
                    .route("/{algo}/{format}", web::post().to(compute_hash))
            )
            .service(
                web::scope("/highway")
                    .route("", web::get().to(hash_highway))
                    .route("", web::post().to(hash_highway))
                    .route("64", web::get().to(hash_highway64))
                    .route("64", web::post().to(hash_highway64))
                    .route("128", web::get().to(hash_highway128))
                    .route("128", web::post().to(hash_highway128))
            )
    })
    .bind("127.0.0.1:1313")?
    .run()
    .await
}
