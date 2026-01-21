mod systeminfo;

use systeminfo::{ get_system_info };

fn main() ->  Result<(), Box<dyn std::error::Error>>  {
    let x = get_system_info()?;

    println!("{:?}", x);

    Ok(())
}
