
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

    handle.join().unwrap();

    let v = vec![1, 2, 3];
    let handler = thread::spawn(||{
        println!("here's a vector: {:?}", v);
    });
    handler.join().unwrap();
}
