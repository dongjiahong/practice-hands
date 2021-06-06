use List::*;
// `#[allow(dead_code)]` 属性可以禁用`dead_code`lint
// 即用来隐藏未使用代码的警告
#[allow(dead_code)]
struct Pair(i32, f32);

struct Point {
    x: f32,
    y: f32,
}

#[allow(dead_code)]
struct Rectangle {
    p1: Point,
    p2: Point,
}

impl Rectangle {
    // 这是一个实例方法
    // `&self`是`self: &Self`的语法糖，其中`Self`是方法的调用者
    // 类型， 在这个例子中`Self`=`Rectangle`
    fn rect_area(&self) -> f32 {
        // `self`通过点运算符来访问结构体字段
        // 这里使用了解构，来获取self.p1的参数信息
        let Point { x: x1, y: y1 } = self.p1;
        let Point { x: x2, y: y2 } = self.p2;

        // `abs`是一个`f32`类型的方法，返回调用者的绝对值
        ((x1 - x2) * (y1 - y2)).abs()
    }
}

// 定义一个枚举变量，里面可能有两种类型，分别是：
#[allow(dead_code)]
enum List {
    // Cons: 元组结构体，包含一个元素和一个指向下一个节点的指针
    Cons(u32, Box<List>),
    // Nil：末节点，表明链表结果
    Nil,
}

impl List {
    // 创建一个空列表
    fn new() -> List {
        Nil
    }

    // 处理一个列表，得到一个头部带上一个新元素的同样类型的列表并返回此值
    fn prepend(self, elem: u32) -> List {
        Cons(elem, Box::new(self))
    }

    // 返回列表的长度
    fn len(&self) -> u32 {
        // `self`必须匹配，因为这个方法的行为取决于`self`的变化类型
        // `self`为`&List`类型， `*self`为`List`类型，一个具体的`T`类型匹配
        // 要参考引用`&T`的匹配
        match *self {
            // 不能得到tail的所有权，因为`self`是借用的
            // 而是得到一个tail引用
            Cons(_, ref tail) => 1 + tail.len(),
            // 基本情形，空列表的长度为0
            Nil => 0,
        }
    }

    fn stringify(&self) -> String {
        match *self {
            Cons(head, ref tail) => {
                format!("{}, {}", head, tail.stringify())
            }
            Nil => {
                format!("Nil")
            }
        }
    }
}

fn main() {
    let point: Point = Point { x: 1.3, y: 4.4 };
    println!("point coordinares: ({}, {})", point.x, point.y);

    // 使用let绑定来结构point, 这种也叫做模式解构，就是跟构造相反，
    // 举个简单的例子：
    /* let t = (1, '2', 3.3);
     * let (one, two, three) = t;
     * println!("one: {}, two: {}, three: {}", one, two, three);
     * # one: 1, two: 2, three: 3.3
     */
    let Point { x: my_x, y: my_y } = point;

    let _rectangle = Rectangle {
        p1: Point { x: my_y, y: my_x },
        p2: point,
    };

    println!("rectangle area is: {}", _rectangle.rect_area());

    let pair = Pair(1, 0.2);

    println!("pair contains {:?} and {:?}", pair.0, pair.1);

    let Pair(integer, decimal) = pair;
    println!("pair contains {:?} and {:?}", integer, decimal);

    println!("------------ 华丽的分隔符 -------------");
    let mut list = List::new();
    list = list.prepend(1);
    list = list.prepend(2);
    list = list.prepend(3);
    list = list.prepend(4);

    // 显示链表的最后状态
    print!("linked list has length: {}", list.len());
    println!("{}", list.stringify());
}
