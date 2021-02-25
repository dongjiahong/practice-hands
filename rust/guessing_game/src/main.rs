use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn hello() -> i32 {
    let a = [1, 2, 3, 4, 5];
    for e in a.iter() {
        println!("hello => {}", e)
    }

    for n in (1..6).rev() {
        println!(" {}!", n);
    }
    1 + 4
}

fn main() {
    println!("请猜一个数字!");

    let secret_number = rand::thread_rng().gen_range(1, 101);
    println!("the secret number is: {}", secret_number);

    loop {
        println!("请输入一个数字!");
        let mut guess = String::new(); // mut 关键字可变 let guess = String::new(); // 不可变

        io::stdin().read_line(&mut guess).expect("读取失败");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("你猜的数字：{}", guess);

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("too small!"),
            Ordering::Greater => println!("too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }

    println!("hello {}: ", hello())
}
