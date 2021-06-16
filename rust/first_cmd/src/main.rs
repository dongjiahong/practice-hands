extern crate clap;
use clap::{Arg, App};
use std::path::Path;
use std::process;
use std::fs::File;
use std::io::{Read, Write};

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
        println!("Value for file argument: {}", file);
        if Path::new(&file).exists() {
            match File::open(file) {
                Ok(mut f) => {
                    let mut data = String::new();
                    f.read_to_string(&mut data).expect("[kt Error] unable to read the file.");
                    let stdout = std::io::stdout(); // 获取全局stdout对象
                    let mut handle = std::io::BufWriter::new(stdout); // 可选项：将handle包装在缓冲区中
                    match writeln!(handle, "{}", data) {
                        Ok(_res) => {},
                        Err(err) => {
                            eprintln!("[kt Error] Unable to display the file contents. {:?}", err);
                            process::exit(1);
                        },
                    }
                }
                Err(err) => {
                    eprintln!("[kt Error] Unable to read the file. {:?}", err);
                    process::exit(1);
                },
            }
        } else {
            eprintln!("[kt Error] No such file or directory.");
            process::exit(1);
        }
    }
}