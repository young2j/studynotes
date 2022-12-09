// 1. 关联类型（associated types）
// 用于在trait中指定类型占位符，这样 trait 的方法签名中就可以使用这些占位符类型。
pub trait Iterator {
    // 类型占位，类似泛型，但泛型在具体实现和使用时需要指定具体类型，类型占位符则不用
    type Item;

    fn next(&mut self) -> Option<Self::Item>;
}

// 2. 默认类型参数和运算符重载
// !默认+号右边的类型是实现Add的类型Self, 可以自定义实现默认的类型参数
// pub trait Add<Rhs = Self> {
//     type Output;
//     fn add(self, rhs: Rhs) -> Self::Output;
// }
use std::ops::Add;
struct Point {
    x: i32,
    y: i32,
}
impl Add for Point {
    type Output = Point;
    fn add(self, rhs: Self) -> Self::Output {
        Point {
            x: self.x + rhs.x,
            y: self.y + rhs.y,
        }
    }
}
struct FPoint {
    x: f32,
    y: f32,
}
impl Add<FPoint> for Point {
    type Output = Point;
    fn add(self, rhs: FPoint) -> Self::Output {
        Point {
            x: self.x + rhs.x as i32,
            y: self.y + rhs.y as i32,
        }
    }
}

// 3. 使用完全限定语法消除歧义
// <Type as Trait>::function(self_if_method, args, ...)
pub trait Animal {
    fn fly();
    fn name(&self) {
        println!("animal");
    }
}

struct Dog;
impl Dog {
    fn fly() {
        println!("dog that can fly");
    }
    fn name(&self) {
        println!("dog");
    }
}

impl Animal for Dog {
    fn fly() {
        println!("animal that can fly");
    }
}
// 4. super trait -  给一个trait绑定另一个trait
trait StarPrint: std::fmt::Display {
    fn print(&self) {
        let output = self.to_string();
        let len = output.len();
        println!("{}", "*".repeat(len + 4));
        println!("*{}*", " ".repeat(len + 2));
        println!("* {} *", output);
        println!("*{}*", " ".repeat(len + 2));
        println!("{}", "*".repeat(len + 4));
    }
}
impl StarPrint for Point {} // 实现StartPrint必须先实现Display trait
impl std::fmt::Display for Point {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "({},{})", self.x, self.y)
    }
}

// 5. newtype 模式用以在外部类型上实现外部 trait
// 例如无法为Vec实现Display trait，因为二者均定义在当前crate外部
// 可以使用元组结构体定义一个新类型(包装一层结构)，再为其实现trait
struct Wrapper(Vec<String>);
impl std::fmt::Display for Wrapper {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        writeln!(f, "[ {} ]", self.0.join(", "))
    }
}

#[allow(unused)]
pub fn run() {
    let dog = Dog {};
    dog.name(); // 等同于Animal::name(&dog); Dog::name(&dog);
    Dog::fly();

    // Animal::fly(); X
    <Dog as Animal>::fly(); // 使用完全限定语法才能访问到Animal的fly函数

    let p = Point { x: 1, y: 2 };
    p.print();

    let v = Wrapper(vec![String::from("a"), String::from("b")]);
    println!("{}", v);
}
