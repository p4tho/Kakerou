mod systeminfo;

use systeminfo::{ SystemInfo, get_system_info };

pub struct C2Agent {
    system_info: SystemInfo
}

impl C2Agent {
    pub fn new() -> Result<Self, Box<dyn std::error::Error>> {
        Ok(C2Agent {
            system_info: get_system_info()?,
        })
    }
}

fn main() ->  Result<(), Box<dyn std::error::Error>> {
    let x = C2Agent::new()?;

    println!("{:?}", x.system_info);

    Ok(())
}
