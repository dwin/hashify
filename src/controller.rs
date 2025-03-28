use actix_web::{web, HttpResponse};
use serde::{Deserialize, Serialize};
use crate::hash::{compute_hash, HashAlgorithm, HashFormat};
use crate::highway::{hash_highway, hash_highway64, hash_highway128};
use crate::keygen::generate_key;
use crate::methods::list_methods;
use crate::status::get_status;

#[derive(Serialize)]
struct BasicError {
    error: String,
}

#[derive(Serialize)]
struct HashResp {
    digest: String,
    digest_enc: String,
    r#type: String,
    key: Option<String>,
}

#[derive(Deserialize)]
struct HashRequest {
    value: Option<String>,
    key: Option<String>,
}

pub async fn get_index() -> HttpResponse {
    HttpResponse::Found()
        .header("Location", "https://hashify.net")
        .finish()
}

pub async fn get_status() -> HttpResponse {
    let status = get_status();
    HttpResponse::Ok().json(status)
}

pub async fn list_methods() -> HttpResponse {
    let methods = list_methods();
    HttpResponse::Ok().json(methods)
}

pub async fn keygen(web::Path(length): web::Path<usize>) -> HttpResponse {
    match generate_key(length) {
        Ok(key) => HttpResponse::Ok().json(key),
        Err(err) => HttpResponse::BadRequest().json(BasicError { error: err.to_string() }),
    }
}

pub async fn compute_hash(
    web::Path((algo, format)): web::Path<(String, String)>,
    web::Query(query): web::Query<HashRequest>,
    body: web::Bytes,
) -> HttpResponse {
    let algorithm = match HashAlgorithm::from_str(&algo) {
        Some(algo) => algo,
        None => {
            return HttpResponse::NotFound().json(BasicError {
                error: "Path does not match available endpoint (algorithm), see API documentation".to_string(),
            });
        }
    };

    let format = match HashFormat::from_str(&format) {
        Some(format) => format,
        None => {
            return HttpResponse::NotFound().json(BasicError {
                error: "Path does not match available endpoint (output), see API documentation".to_string(),
            });
        }
    };

    let value = query.value.unwrap_or_else(|| String::from_utf8_lossy(&body).to_string());
    let key = query.key;

    match compute_hash(algorithm, format, &value, key.as_deref()) {
        Ok((digest, key)) => HttpResponse::Ok().json(HashResp {
            digest,
            digest_enc: format.to_string(),
            r#type: algorithm.to_string(),
            key,
        }),
        Err(err) => HttpResponse::BadRequest().json(BasicError { error: err.to_string() }),
    }
}

pub async fn hash_highway(
    web::Query(query): web::Query<HashRequest>,
    body: web::Bytes,
) -> HttpResponse {
    let value = query.value.unwrap_or_else(|| String::from_utf8_lossy(&body).to_string());
    let key = query.key;

    match hash_highway(&value, key.as_deref()) {
        Ok((digest, key)) => HttpResponse::Ok().json(HashResp {
            digest,
            digest_enc: "hex".to_string(),
            r#type: "HighwayHash-256".to_string(),
            key: Some(key),
        }),
        Err(err) => HttpResponse::BadRequest().json(BasicError { error: err.to_string() }),
    }
}

pub async fn hash_highway64(
    web::Query(query): web::Query<HashRequest>,
    body: web::Bytes,
) -> HttpResponse {
    let value = query.value.unwrap_or_else(|| String::from_utf8_lossy(&body).to_string());
    let key = query.key;

    match hash_highway64(&value, key.as_deref()) {
        Ok((digest, key)) => HttpResponse::Ok().json(HashResp {
            digest,
            digest_enc: "hex".to_string(),
            r#type: "HighwayHash-64".to_string(),
            key: Some(key),
        }),
        Err(err) => HttpResponse::BadRequest().json(BasicError { error: err.to_string() }),
    }
}

pub async fn hash_highway128(
    web::Query(query): web::Query<HashRequest>,
    body: web::Bytes,
) -> HttpResponse {
    let value = query.value.unwrap_or_else(|| String::from_utf8_lossy(&body).to_string());
    let key = query.key;

    match hash_highway128(&value, key.as_deref()) {
        Ok((digest, key)) => HttpResponse::Ok().json(HashResp {
            digest,
            digest_enc: "hex".to_string(),
            r#type: "HighwayHash-128".to_string(),
            key: Some(key),
        }),
        Err(err) => HttpResponse::BadRequest().json(BasicError { error: err.to_string() }),
    }
}
