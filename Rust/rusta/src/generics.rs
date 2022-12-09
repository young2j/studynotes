use std::fmt::Display;

struct Point<T> {
    x: T,
    y: T,
}

impl<T> Point<T> {
    fn new(x: T, y: T) -> Point<T> {
        Point { x, y }
    }
    fn get_x(&self) -> &T {
        &self.x
    }
    fn get_y(&self) -> &T {
        &self.y
    }
    fn mix_t<T2>(self, other: Point<T2>) -> (T, T2) {
        (self.x, other.x)
    }
}

// 只对f32类型的Point实现方法
impl Point<f32> {
    fn distance_from_origin(&self) -> f32 {
        (self.x.powi(2) + self.y.powi(2)).sqrt()
    }
}

// 使用 trait bound 有条件地实现方法
impl<T: Display + PartialOrd> Point<T> {
    fn cmp_display(&self) {
        if self.x > self.y {
            println!("x is greater than y")
        } else {
            println!("x is less than y")
        }
    }
}

pub fn run() {
    let p = Point::new(11, 22);
    println!("p.x={} p.y={}", p.get_x(), p.get_y());

    let fp: Point<f32> = Point::new(3.3, 4.4);
    let distance = fp.distance_from_origin();
    println!("x={} y={} distance={}", fp.x, fp.y, distance);

    let (x1, x2) = p.mix_t(fp);
    println!("p.x={} fp.x={}", x1, x2);
}
