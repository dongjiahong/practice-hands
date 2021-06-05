use std::fmt;

#[derive(Debug)]
struct Point2D {
    x: f32,
    y: f32,
}

impl fmt::Display for Point2D {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "Display: {} + {}i", self.x, self.y)
    }
}

struct List(Vec<i32>);

impl fmt::Display for List {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let List(ref vec) = *self;

        write!(f, "[");
        for (count, v) in vec.iter().enumerate() {
            if count != 0 {
                write!(f, ", ");
            }

            write!(f, "{}", v);
        }
        write!(f, "]")
    }
}

fn main() {
    println!("Hello, world!");
    // `{}`占位符，会被任意变量替换
    println!("{} day", 31);

    // 位置参数
    println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");

    // 使用赋值语句
    println!(
        "{subject} {verb} {object}",
        object = "the lazy dog",
        subject = "the quick brown fox",
        verb = "jumps over"
    );

    // 特色格式在后面加`:`符号
    println!("{} of {:b} people know binary, the other half don't", 1, 2);

    // 指定宽度右对齐文本
    // 下面语句输出"     1",5个空格连着1
    println!("{number:>with$}", number = 1, with = 6);
    // 前面补0而不是空格
    println!("{number:>0with$}", number = 1, with = 6);

    // println!会检查使用到的参数是否正确
    println!("My name is {0}, {1} {0}", "Bond", "Tom");

    #[derive(Debug)]
    struct Structure(i32);

    // 但是像结构体这个样自定义的类型需要更复杂的方式处理。
    println!("This struct `{:?}` won't print...", Structure(3));

    let pi = 3.141592;
    println!("Pi is roughly {pi:.with$}", pi = pi, with = 4);

    let point = Point2D { x: 3.3, y: 7.2 };
    println!("{}", point);
    println!("Debug: {:?}", point);

    let v = List(vec![1, 2, 3]);
    println!("{}", v);
}
