// Mutex<T> 是一个智能指针
// Mutex<T>.lock() -> Result<MutexGuard<T>>
// MutexGuard<T> 也是一个智能指针
//  - 可视为一个内部数据可变的引用;
//  - 其实现了Deref和Drop trait, 在离开作用域时自动执行清理, 可自动释放锁;

// Rc<T>  -- 非并发安全的引用计数
// Arc<T> -- 并发安全引用计数, atomic reference count

use std::{
    sync::{Arc, Mutex},
    thread,
};

#[allow(unused)]
pub fn run() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            // lock()会阻塞当前线程，直到拥有锁为止
            // num视为一个内部数据的可变引用
            let mut num = counter.lock().unwrap();
            *num += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("counter={:?}", *counter.lock().unwrap());
}
