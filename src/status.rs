use serde::Serialize;
use std::time::{Duration, SystemTime};

#[derive(Serialize)]
pub struct Status {
    status: String,
    uptime: String,
    hostname: String,
    hashes_generated: usize,
    keys_generated: usize,
}

pub fn get_status() -> Status {
    let start_time = SystemTime::now();
    let uptime = match start_time.elapsed() {
        Ok(elapsed) => format!("{:?}", elapsed),
        Err(_) => "unknown".to_string(),
    };

    let hostname = match hostname::get() {
        Ok(name) => name.to_string_lossy().into_owned(),
        Err(_) => "unknown".to_string(),
    };

    Status {
        status: "OK".to_string(),
        uptime,
        hostname,
        hashes_generated: 0, // Replace with actual value
        keys_generated: 0,   // Replace with actual value
    }
}
