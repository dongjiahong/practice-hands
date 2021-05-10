#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    // 成员函数里面第一个参数是self, 成员函数使用.来调用eg: a.b
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
    // 关联函数,关联函数用::来调用eg: a::b
    fn square(size: u32) -> Rectangle {
        Rectangle{width: size, height: size}
    }
}

#[derive(Debug)]
enum UsState {
    Alabama,
    Alaska,
}

enum Coin {
    Dime,
    Quarter(UsState),
}

fn value_in_cents(coin: Coin) -> u8 {
    match coin{
        // match 控制流程，使用=>来表达分支，如果有多个表达用{}包裹起来
        Coin::Dime => {
            println!("Lucky Dime");
            10
        },
        Coin::Quarter(state) =>{
            println!("State quarter from {:?}!", state);
            25
        },
    }
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };
    let rect3 = Rectangle {
        width: 24,
        height: 22,
    };
    println!("rect3: {:#?}", rect3);
    println!(
        "The area of the rectangle is {} square pixels.",
        rect1.area()
    );
    println!("Can rect1 hold rect3? {} ", rect1.can_hold(&rect3));
    println!("create square {:#?}",Rectangle::square(11));
    
    println!("coin: {}",  value_in_cents(Coin::Dime));
    println!("coin: {}",  value_in_cents(Coin::Quarter(UsState::Alabama)));
    println!("coin: {}",  value_in_cents(Coin::Quarter(UsState::Alaska)));
}
