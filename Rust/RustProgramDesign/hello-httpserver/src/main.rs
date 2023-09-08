use std::thread;

use hello_httpserver::create_server;

fn main() {
    let addr = "0.0.0.0:7878";
    let handle = thread::spawn(|| create_server(addr));
    println!("server listen on {}", addr);
    handle.join().unwrap();
}
