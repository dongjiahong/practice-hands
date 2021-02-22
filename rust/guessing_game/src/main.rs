use std::io;

fn main() {
    println!("请猜一个数字!");
    println!("请输入一个数字!");

    let mut guess = String::new(); // mut 关键字可变
                                   //let guess = String::new(); // 不可变

    io::stdin().read_line(&mut guess).expect("读取失败");

    println!("你猜的数字：{}", guess);
}
