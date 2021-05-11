use std::fs::File;
use std::io::{self,Read};

fn main() {
    let v = vec![1, 2, 3, 4, 5, 6];

    let third: &i32 = &v[2];
    println!("The third element is {}", third);

    match v.get(2) {
        Some(third) => println!("The third element is {}", third),
        None => println!("There is no third element"),
    }

    let mut s1 = String::from("foo");
    let s2 = "bar";
    s1.push_str(s2);
    println!("S2 is {}", s2);
    println!("S1 is {}", s1);
    println!("S2 is {}", s2);
    // error & panics
}

fn read_username_from_file() -> Result<String,io::Error> {
    let f = File::open("hello.txt");

    let mut f = match f{
        Ok(file) => file,
        Err(e) => return Err(e),
    };

    let mut s = String::new();
    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    };
}

fn easy_read_username_from_fil() -> Result<String, io::Error> {
    let mut s = String::new();
    // ?操作符，是帮我们处理match等一堆错误的写法
    File::open("hello.txt")?.read_to_string(&mut s)?;
    Ok(s)
}
