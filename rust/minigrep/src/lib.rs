use std::error::Error;
use std::fs;

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn one_result() {
        let query = "duct";
        let contents = "\
Rust:
safe, fast, productive.
Pick three.";
        assert_eq!(vec!["safe, fast, productive."], search(query, contents));
    }
}

// 这里得'a 就是指定声明周期，这里指定返回得vec中应该包含引用参数contents slice得字符串slice.
// 意思就是说我们告诉rust函数search返回得数据将于search函数中得参数contents得数据存在得一样久。
pub fn search<'a>(query: &str, contents: &'a str) -> Vec<&'a str> {
    let mut results = Vec::new();
    for line in contents.lines() {
        if line.contains(query) {
            results.push(line);
        }
    }
    results
}

pub struct Config {
    pub query: String,
    pub filename: String,
}

impl Config {
    pub fn new(args: &Vec<String>) -> Result<Config, &'static str> {
        if args.len() < 2 {
            return Err("not enough arguments");
        }
        let query = args[1].clone();
        let filename = args[2].clone();

        Ok(Config { query, filename })
    }
}

pub fn run(config: Config) -> Result<(), Box<dyn Error>> {
    // 下面一行最后得？表示如果出错就返回给他得调用者，让调用者处理
    let contents = fs::read_to_string(config.filename)?;
    for line in search(&config.query, &contents) {
        println!("result => {}", line);
    }
    Ok(())
}
