// 此声明会查找名为`trait_bli.rs`或者`trait_lib/mod.rs`文件，
// 并将该未见得内容插入到此作用域名为`trait_lib`得模块里
mod trait_lib;

#[derive(Debug)]
struct Point<T,U> {
    x: T,
    y: U,
}

impl<T,U> Point<T,U>{
    fn x(&self) -> &T{
        &self.x
    }
}

fn main() {
    let both_integer = Point{x: 5, y: 10};
    let inter_float = Point{x:5, y: 10.0};
    println!("both integer {:?}", both_integer);
    println!("integer float {:?}", inter_float);
    println!("integer float {}", inter_float.x());

    let tweet = trait_lib::Tweet{
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probable already know, people"),
        reply: false,
        retweet: false,
    };
    println!("1 new tweet: {}", tweet.summarise());
}
