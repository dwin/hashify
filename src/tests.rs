use actix_web::{test, App};
use crate::controller::{get_index, get_status, list_methods, keygen, compute_hash, hash_highway, hash_highway64, hash_highway128};

#[actix_rt::test]
async fn test_get_index() {
    let mut app = test::init_service(App::new().route("/", web::get().to(get_index))).await;
    let req = test::TestRequest::get().uri("/").to_request();
    let resp = test::call_service(&mut app, req).await;
    assert_eq!(resp.status(), http::StatusCode::FOUND);
}

#[actix_rt::test]
async fn test_get_status() {
    let mut app = test::init_service(App::new().route("/status", web::get().to(get_status))).await;
    let req = test::TestRequest::get().uri("/status").to_request();
    let resp = test::call_service(&mut app, req).await;
    assert_eq!(resp.status(), http::StatusCode::OK);
}

#[actix_rt::test]
async fn test_list_methods() {
    let mut app = test::init_service(App::new().route("/methods", web::get().to(list_methods))).await;
    let req = test::TestRequest::get().uri("/methods").to_request();
    let resp = test::call_service(&mut app, req).await;
    assert_eq!(resp.status(), http::StatusCode::OK);
}

#[actix_rt::test]
async fn test_keygen() {
    let mut app = test::init_service(App::new().route("/keygen/{length}", web::get().to(keygen))).await;
    let req = test::TestRequest::get().uri("/keygen/32").to_request();
    let resp = test::call_service(&mut app, req).await;
    assert_eq!(resp.status(), http::StatusCode::OK);
}

#[actix_rt::test]
async fn test_compute_hash() {
    let mut app = test::init_service(App::new().service(
        web::scope("/hash")
            .route("/{algo}/{format}", web::get().to(compute_hash))
            .route("/{algo}/{format}", web::post().to(compute_hash))
    )).await;

    let req = test::TestRequest::get().uri("/hash/SHA256/hex?value=hello").to_request();
    let resp = test::call_service(&mut app, req).await;
    assert_eq!(resp.status(), http::StatusCode::OK);
}

#[actix_rt::test]
async fn test_hash_highway() {
    let mut app = test::init_service(App::new().service(
        web::scope("/highway")
            .route("", web::get().to(hash_highway))
            .route("", web::post().to(hash_highway))
    )).await;

    let req = test::TestRequest::get().uri("/highway?value=hello").to_request();
    let resp = test::call_service(&mut app, req).await;
    assert_eq!(resp.status(), http::StatusCode::OK);
}

#[actix_rt::test]
async fn test_hash_highway64() {
    let mut app = test::init_service(App::new().service(
        web::scope("/highway")
            .route("64", web::get().to(hash_highway64))
            .route("64", web::post().to(hash_highway64))
    )).await;

    let req = test::TestRequest::get().uri("/highway/64?value=hello").to_request();
    let resp = test::call_service(&mut app, req).await;
    assert_eq!(resp.status(), http::StatusCode::OK);
}

#[actix_rt::test]
async fn test_hash_highway128() {
    let mut app = test::init_service(App::new().service(
        web::scope("/highway")
            .route("128", web::get().to(hash_highway128))
            .route("128", web::post().to(hash_highway128))
    )).await;

    let req = test::TestRequest::get().uri("/highway/128?value=hello").to_request();
    let resp = test::call_service(&mut app, req).await;
    assert_eq!(resp.status(), http::StatusCode::OK);
}
