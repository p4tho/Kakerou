#[derive(Debug)]
pub struct AgentConfig {
    pub name: String,
    pub server_url: String,
    pub duration: u64,
    pub jitter_percent: f64,
}