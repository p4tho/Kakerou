use cryptify::encrypt_string;
use hostname::get;
use local_ip_address::local_ip;

#[derive(Debug)]
pub struct SystemInfo {
    pub os: String,
    pub arch: String,
    pub local_ip: String,
    pub hostname: String,
    pub user: String,
}

pub fn get_hostname() -> Result<String, Box<dyn std::error::Error>> {
    Ok(get()?.into_string().unwrap_or("localhost".to_string()))
}

pub fn get_ip_address() -> Result<String, Box<dyn std::error::Error>> {
    Ok(local_ip()?.to_string())
}

pub fn get_system_info() -> Result<SystemInfo, Box<dyn std::error::Error>> {
    Ok(SystemInfo {
        os: std::env::consts::OS.to_string(),
        arch: std::env::consts::ARCH.to_string(),
        local_ip: get_ip_address()
            .unwrap_or_else(|_| encrypt_string!("unknown")),
        hostname: get_hostname()
            .unwrap_or_else(|_| encrypt_string!("unknown").to_string()),
        user: std::env::var("USER") // Check for Linux
            .or_else(|_| std::env::var("USERNAME")) // Check for Windows
            .unwrap_or_else(|_| encrypt_string!("unknown").to_string()),
    })
}