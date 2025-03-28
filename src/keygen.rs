use rand::Rng;
use serde::Serialize;
use std::error::Error;

#[derive(Serialize)]
pub struct KeyResp {
    key_hex: String,
    length: usize,
}

pub fn generate_key(length: usize) -> Result<KeyResp, Box<dyn Error>> {
    if length > 256 {
        return Err(format!("Key length request over limit of 256, requested: {}", length).into());
    }

    let key: Vec<u8> = (0..length).map(|_| rand::thread_rng().gen()).collect();
    let key_hex = hex::encode(&key);

    Ok(KeyResp {
        key_hex,
        length,
    })
}
