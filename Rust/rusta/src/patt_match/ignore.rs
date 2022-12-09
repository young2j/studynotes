fn use_lodash() {
    let x = (1, 2, 3, 4, 5);
    match x {
        (first, _, third, _, five) => println!("{}", first + third + five),
    }
    let _y = 5;
}

struct Point {
    x: i32,
    y: i32,
    z: i32,
}
fn use_dotdot() {
    let p = Point { x: 1, y: 2, z: 3 };
    match p {
        Point { x, .. } => println!("{}", x),
    }

    let nums = (1, 2, 3, 4, 5);
    let (_first, .., _last) = nums;
}

