use serde::Serialize;

#[derive(Serialize)]
pub struct HashMethod {
    name: String,
    endpoint: String,
    min_key_length: usize,
    max_key_length: usize,
}

pub fn list_methods() -> Vec<HashMethod> {
    vec![
        HashMethod {
            name: "Blake2b-256".to_string(),
            endpoint: "/hash/BLAKE2B-256".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "Blake2b-384".to_string(),
            endpoint: "/hash/BLAKE2B-384".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "Blake2b-512".to_string(),
            endpoint: "/hash/BLAKE2B-512".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "Blake2s-128".to_string(),
            endpoint: "/hash/BLAKE2s-128".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "Blake2s-256".to_string(),
            endpoint: "/hash/BLAKE2s-256".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "HighwayHash-256".to_string(),
            endpoint: "/hash/HIGHWAY".to_string(),
            min_key_length: 32,
            max_key_length: 32,
        },
        HashMethod {
            name: "HighwayHash-64".to_string(),
            endpoint: "/hash/HIGHWAY-64".to_string(),
            min_key_length: 32,
            max_key_length: 32,
        },
        HashMethod {
            name: "HighwayHash-128".to_string(),
            endpoint: "/hash/HIGHWAY-128".to_string(),
            min_key_length: 32,
            max_key_length: 32,
        },
        HashMethod {
            name: "MD4".to_string(),
            endpoint: "/hash/MD4".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "MD5".to_string(),
            endpoint: "/hash/MD5".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "SHA1".to_string(),
            endpoint: "/hash/SHA1".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "SHA256".to_string(),
            endpoint: "/hash/SHA256".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "SHA384".to_string(),
            endpoint: "/hash/SHA384".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "SHA512".to_string(),
            endpoint: "/hash/SHA512".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "SHA512-256".to_string(),
            endpoint: "/hash/SHA512-256".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "SHA3-256".to_string(),
            endpoint: "/hash/SHA3-256".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "SHA3-384".to_string(),
            endpoint: "/hash/SHA3-384".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
        HashMethod {
            name: "SHA3-512".to_string(),
            endpoint: "/hash/SHA3-512".to_string(),
            min_key_length: 0,
            max_key_length: 0,
        },
    ]
}
