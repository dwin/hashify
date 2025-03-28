use sha2::{Sha256, Sha512, Digest};
use blake2::{Blake2b, Blake2s};
use base64;
use hex;
use std::error::Error;

pub enum HashAlgorithm {
    Sha256,
    Sha512,
    Blake2b,
    Blake2s,
}

impl HashAlgorithm {
    pub fn from_str(algo: &str) -> Option<HashAlgorithm> {
        match algo.to_uppercase().as_str() {
            "SHA256" => Some(HashAlgorithm::Sha256),
            "SHA512" => Some(HashAlgorithm::Sha512),
            "BLAKE2B" => Some(HashAlgorithm::Blake2b),
            "BLAKE2S" => Some(HashAlgorithm::Blake2s),
            _ => None,
        }
    }

    pub fn to_string(&self) -> String {
        match self {
            HashAlgorithm::Sha256 => "SHA256".to_string(),
            HashAlgorithm::Sha512 => "SHA512".to_string(),
            HashAlgorithm::Blake2b => "Blake2b".to_string(),
            HashAlgorithm::Blake2s => "Blake2s".to_string(),
        }
    }
}

pub enum HashFormat {
    Base32,
    Base64,
    Base64Url,
    Hex,
}

impl HashFormat {
    pub fn from_str(format: &str) -> Option<HashFormat> {
        match format.to_lowercase().as_str() {
            "base32" => Some(HashFormat::Base32),
            "base64" => Some(HashFormat::Base64),
            "base64url" => Some(HashFormat::Base64Url),
            "hex" => Some(HashFormat::Hex),
            _ => None,
        }
    }

    pub fn to_string(&self) -> String {
        match self {
            HashFormat::Base32 => "base32".to_string(),
            HashFormat::Base64 => "base64".to_string(),
            HashFormat::Base64Url => "base64url".to_string(),
            HashFormat::Hex => "hex".to_string(),
        }
    }
}

pub fn compute_hash(
    algorithm: HashAlgorithm,
    format: HashFormat,
    value: &str,
    key: Option<&str>,
) -> Result<(String, Option<String>), Box<dyn Error>> {
    let mut hasher: Box<dyn Digest> = match algorithm {
        HashAlgorithm::Sha256 => Box::new(Sha256::new()),
        HashAlgorithm::Sha512 => Box::new(Sha512::new()),
        HashAlgorithm::Blake2b => Box::new(Blake2b::new()),
        HashAlgorithm::Blake2s => Box::new(Blake2s::new()),
    };

    hasher.update(value.as_bytes());

    let result = hasher.finalize();
    let digest = match format {
        HashFormat::Base32 => base32::encode(base32::Alphabet::RFC4648 { padding: false }, &result),
        HashFormat::Base64 => base64::encode(&result),
        HashFormat::Base64Url => base64::encode_config(&result, base64::URL_SAFE_NO_PAD),
        HashFormat::Hex => hex::encode(&result),
    };

    Ok((digest, key.map(|k| k.to_string())))
}
