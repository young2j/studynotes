// 1. 希望在某一个 Future 报错后就立即停止所有 Future 的执行，可以使用 try_join!，特别是当 Future 返回 Result 时

// 2. 传给 try_join! 的所有 Future 都必须拥有相同的错误类型。如果错误类型不同，
//    可以使用来自 futures::future::TryFutureExt 模块的 map_err和err_info方法将错误进行转换

// 3. 如果想同时等待多个 Future ，且任何一个 Future 结束就进行处理，可以使用 futures::select!

use futures::executor::block_on;

async fn do_something() {
    let hi = say_hi().await;
    println!("async doing something! —— say {}", hi)
}

async fn say_hi() -> String {
    String::from("Hi")
}

async fn reply_hello() -> String {
    String::from("Hello")
}

async fn do_another_thing() {
    let hello = reply_hello().await;
    println!("async doing another thing! —— reply {}", hello)
}

pub fn run() {
    let future1 = do_something();
    let future2 = do_another_thing();
    let futs = futures::future::join(future1, future2);
    // let futs = futures::join!(future1, future2); // 直接返回结果而非Future, 内部调用了.await
    // let futs = std::future::join!(future1, future2); // 非稳定版

    block_on(futs);
}
