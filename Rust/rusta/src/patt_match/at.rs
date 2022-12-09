struct Point {
    x: i32,
    y: i32,
    z: i32,
}
fn use_at() {
    let p = Point { x: 1, y: 2, z: 3 };
    match p {
        Point {
            x: x_value @ 0..=1,
            y: 1..=2,
            ..
        } => {
            println!("x saved to x_value={}", x_value);
        }
        _ => println!("other match cases"),
    }
}
