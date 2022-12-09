use std::{thread, time::Duration};

#[allow(unused)]
pub fn run() {
    let handle = thread::spawn(|| {
        for i in 1..10 {
            println!("from spawn thead: {}", i);
            thread::sleep(Duration::from_millis(1));
        }
    });

    for i in 1..5 {
        println!("from main tread: {}", i);
        thread::sleep(Duration::from_millis(2));
    }

    // join使得主线程等待子线程执行完毕
    handle.join().unwrap();
}
