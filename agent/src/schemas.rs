use serde::Deserialize;

#[derive(Debug)]
pub struct AgentConfig {
    pub name: String,
    pub server_url: String,
    pub duration: u64,
    pub jitter_percent: f64,
}

#[derive(Debug, Deserialize)]
pub struct BeaconResponse {
    pub commands: Vec<Command>,
}

#[derive(Debug, Deserialize, PartialEq, Eq)]
pub struct Command {
    pub uid: u64,
    pub command_id: u64,
    pub command: String,
}