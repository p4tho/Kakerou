use cryptify::encrypt_string;
use hostname::get;
use local_ip_address::local_ip;
use sysinfo::{
    Disks, 
    System,
};

#[derive(Debug)]
pub struct SystemInfo {
    pub os: String,
    pub arch: String,
    pub kernel: String,
    pub local_ip: String,
    pub hostname: String,
    pub user: String,
    pub cpu_brand: String,
    pub num_cores: usize,
    pub memory_total_gib: f64,
    pub disk_total_gib: f64,
    pub disk_available_gib: f64,
}

pub fn bytes_to_gib(amount: u64) -> f64 {
    (amount as f64) / 1024.0 / 1024.0 / 1024.0
}

pub fn get_hostname() -> Result<String, Box<dyn std::error::Error>> {
    Ok(get()?.into_string().unwrap_or("localhost".to_string()))
}

pub fn get_ip_address() -> Result<String, Box<dyn std::error::Error>> {
    Ok(local_ip()?.to_string())
}

pub fn get_system_info() -> Result<SystemInfo, Box<dyn std::error::Error>> {
    let mut sys = System::new_all();
    sys.refresh_all();

    // Get total disk and available disk
    let disks = Disks::new_with_refreshed_list();
    let mut disk_total_b = 0u64;
    let mut disk_available_b = 0u64;
    for disk in disks.iter() {
        disk_total_b += disk.total_space();
        disk_available_b += disk.available_space();
    }

    // Return all system information
    Ok(SystemInfo {
        os: std::env::consts::OS
            .to_string(),
        arch: std::env::consts::ARCH
            .to_string(),
        kernel: System::kernel_version()
            .unwrap_or_else(|| encrypt_string!("unknown")),
        local_ip: get_ip_address()
            .unwrap_or_else(|_| encrypt_string!("unknown")),
        hostname: get_hostname()
            .unwrap_or_else(|_| encrypt_string!("unknown")),
        user: std::env::var("USER") // Check for Linux
            .or_else(|_| std::env::var("USERNAME")) // Check for Windows
            .unwrap_or_else(|_| encrypt_string!("unknown")),
        cpu_brand: sys
            .cpus()
            .first()
            .map(|cpu| cpu.brand().to_string())
            .unwrap_or_else(|| encrypt_string!("unknown")),
        num_cores: sys
            .cpus()
            .len(),
        memory_total_gib: bytes_to_gib(sys.total_memory()),
        disk_total_gib: bytes_to_gib(disk_total_b), 
        disk_available_gib: bytes_to_gib(disk_available_b),
    })
}