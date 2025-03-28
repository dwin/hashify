use highway::{HighwayHasher, Key};
use std::error::Error;
use std::fmt;

#[derive(Debug)]
pub struct HighwayHashError {
    details: String,
}

impl HighwayHashError {
    fn new(msg: &str) -> HighwayHashError {
        HighwayHashError {
            details: msg.to_string(),
        }
    }
}

impl fmt::Display for HighwayHashError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.details)
    }
}

impl Error for HighwayHashError {
    fn description(&self) -> &str {
        &self.details
    }
}

pub fn hash_highway(value: &str, key: Option<&str>) -> Result<(String, String), Box<dyn Error>> {
    let key = match key {
        Some(k) => {
            let key_bytes = hex::decode(k)?;
            if key_bytes.len() != 32 {
                return Err(Box::new(HighwayHashError::new(
                    "HighwayHash key parameter must be 32 bytes",
                )));
            }
            Key::from(key_bytes.as_slice())
        }
        None => Key::default(),
    };

    let mut hasher = HighwayHasher::new(&key);
    hasher.append(value.as_bytes());
    let result = hasher.finalize256();
    let digest = hex::encode(result);

    Ok((digest, hex::encode(key.as_bytes())))
}

pub fn hash_highway64(value: &str, key: Option<&str>) -> Result<(String, String), Box<dyn Error>> {
    let key = match key {
        Some(k) => {
            let key_bytes = hex::decode(k)?;
            if key_bytes.len() != 32 {
                return Err(Box::new(HighwayHashError::new(
                    "HighwayHash key parameter must be 32 bytes",
                )));
            }
            Key::from(key_bytes.as_slice())
        }
        None => Key::default(),
    };

    let mut hasher = HighwayHasher::new(&key);
    hasher.append(value.as_bytes());
    let result = hasher.finalize64();
    let digest = hex::encode(result.to_le_bytes());

    Ok((digest, hex::encode(key.as_bytes())))
}

pub fn hash_highway128(value: &str, key: Option<&str>) -> Result<(String, String), Box<dyn Error>> {
    let key = match key {
        Some(k) => {
            let key_bytes = hex::decode(k)?;
            if key_bytes.len() != 32 {
                return Err(Box::new(HighwayHashError::new(
                    "HighwayHash key parameter must be 32 bytes",
                )));
            }
            Key::from(key_bytes.as_slice())
        }
        None => Key::default(),
    };

    let mut hasher = HighwayHasher::new(&key);
    hasher.append(value.as_bytes());
    let result = hasher.finalize128();
    let digest = hex::encode(result);

    Ok((digest, hex::encode(key.as_bytes())))
}
