use std::sync::mpsc;
use std::thread;
use std::time::Duration;

fn main() {
    let handle = thread::spawn(||{
        for i in 0..10 {
            println!("hi number {} from swawned thread!", i);
            thread::sleep(Duration::from_millis(100));
        }
    });

    for i in 0..5 {
        println!("hi number {} from main thread!", i);
        thread::sleep(Duration::from_millis(100));
    }
    // join表示等待线程执行完
    handle.join().unwrap();

    let v = vec![1, 2, 3];
    // move转移所有权，不然线程李访问不到v
    let handler = thread::spawn(move ||{
        println!("here's a vector: {:?}", v);
    });
    handler.join().unwrap();

    let (tx, rx) = mpsc::channel();
    thread::spawn(move ||{
        let val = String::from("hi");
        tx.send(val).unwrap();
    });
    let received = rx.recv().unwrap();
    println!("Got: {}", received);
}
