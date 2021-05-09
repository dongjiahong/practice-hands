fn main() {
    let s = String::from("hello world");

    let word = first_word(&s);

    //s.clear();

    println!("word: {}", word);
    println!("s: {}", s);
    //println!("s2: {}", s[2]); // s已经被清空了,这里panic
    
     println!("new s: {}",new_first_word(&s));
}

// 返回下标
fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }
    s.len()
}

// 返回单词
fn new_first_word(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[..i];
        }
    }
    &s[..]
}
