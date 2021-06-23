use serde::{Serialize, Deserialize};
#[derive(Serialize, Deserialize, Debug)]
struct ServerConfig {
    workers: u64,
    ignore: bool,
    auth_server: Option<String>,
}

fn main() {
    let config = ServerConfig {
        workers: 100,
        ignore: false,
        auth_server: Some(String::from("auth_server.io")),
    };
    {
        println!("To and from Json");
        let serialized = serde_json::to_string(&config).unwrap();
        println!("serialized: {}", serialized);
        println!("");

        let deserialized: ServerConfig = serde_json::from_str(&serialized).unwrap();
        println!("deserialized: {:#?}", deserialized);
    }
    {
        println!("To and from Yaml");
        let serialized = serde_yaml::to_string(&config).unwrap();
        println!("serialized: {}", serialized);
        println!("");

        let deserialized: ServerConfig = serde_yaml::from_str(&serialized).unwrap();
        println!("deserialized: {:#?}", deserialized);
    }

}
