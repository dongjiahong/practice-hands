extern crate clap;
use clap::{Arg, App};


fn main() {
    let matches = App::new("kt")
        .version("0.1.0")
        .author("Jiahong dong")
        .about("A drop in cat replacement written in Rust")
        .arg(Arg::with_name("FILE")
             .help("File to print.")
             .empty_values(false)
             )
        .get_matches();

    if let Some(file) = matches.value_of("FILE") {
        println!("Value for file argument: {}", file)
    }

}
