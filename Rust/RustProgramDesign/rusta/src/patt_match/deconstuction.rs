struct Point {
    x: i32,
    y: i32,
}
fn dec_struct() {
    let p = Point { x: 5, y: 10 };
    let Point { x, y: b } = p;
    println!("x={} b={}", x, b);

    match p {
        Point { x: 0, y: _ } => println!("if x==0"),
        Point { x: _, y: 1 } => println!("if y==1"),
        Point { x: 5, y: _ } => println!("x matched"),
        _ => println!("no match"),
    }
}

fn dec_enum() {
    enum Message {
        Quit,
        Move { x: i32, y: i32 },
        Single(String),
        Mutiple(i32, i32, i32),
    }

    let x = Message::Mutiple(1, 2, 3);
    match x {
        Message::Quit => println!("match quit"),
        Message::Move { x, y } => println!("match x={}, y={}", x, y),
        Message::Single(s) => println!("match str = {}", s),
        Message::Mutiple(a, b, c) => println!("match a={} b={} c={}", a, b, c),
    }
}

fn dec_tuple() {
    let ((_a, _b), Point { x: _, y: _ }) = (("a", "b"), Point { x: 1, y: 2 });
}
