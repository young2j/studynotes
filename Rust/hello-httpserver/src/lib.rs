mod text;
mod thread_pool;

use std::{
    io::{Read, Write},
    net::{TcpListener, TcpStream},
    thread,
    time::Duration,
};

pub fn create_server(addr: &str) {
    let listener = TcpListener::bind(addr).unwrap();
    let pool = thread_pool::ThreadPool::new(4);
    for stream in listener.incoming() {
        let mut buf = [0; 1024];
        let mut stream = stream.unwrap();
        stream.read(&mut buf).unwrap();

        // println!("======Received Request======\n{}", String::from_utf8_lossy(&buf));
        // 单线程版
        // thread::spawn(move || handle_response(stream, &buf));
        // 多线程版-线程池
        pool.execute(move || handle_response(stream, &buf))
    }
}

fn handle_response(mut stream: TcpStream, buf: &[u8]) {
    thread::sleep(Duration::from_secs(5)); // 模拟慢请求

    let get = b"GET / HTTP/1.1\r\n";
    let (status_line, content) = if buf.starts_with(get) {
        ("HTTP/1.1 200 OK", text::CONTENT200)
    } else {
        ("HTTP/1.1 404 NOT FOUND", text::CONTENT404)
    };

    let resp = format!(
        "{}\r\nContent-Length: {}\r\n\r\n{}",
        status_line,
        content.len(),
        content,
    );

    stream.write(resp.as_bytes()).unwrap();
    stream.flush().unwrap();
}
