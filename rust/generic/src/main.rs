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
}
