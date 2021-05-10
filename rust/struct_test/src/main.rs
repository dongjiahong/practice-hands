#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
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
    println!("Can rect1 hold rect3? {} ", rect1.can_hold(&rect3))
}
