// 生命周期注解三条规则:
// 1. 每一个引用参数都有它自己的生命周期;
// 2. 如果只有一个输入生命周期，那么它被赋予所有输出生命周期;
// 3. 如果有多个输入生命周期且其中一个参数是&self(或&mut self), 那么self的生命周期被赋予所有输出生命周期;
//    此条规则使得方法更易读写, 往往可以省略生命周期注解。所以也称为生命周期省略规则。

// 静态生命周期
// 静态生命周期能够存活于整个程序期间;
// 所有的字符串字面值都拥有'static 生命周期

use std::fmt::Display;

// 引用的泛型生命周期
fn longest<'a, T>(x: &'a str, y: &'a str, z: T) -> &'a str
where
    T: Display,
{
    println!("z: {}", z);
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

// 结构体中定义生命周期注解
struct Annotation<'a> {
    x: &'a str,
    y: &'a str,
}

impl<'a> Annotation<'a> {
    fn get_x(&self) -> &str {
        self.x
    }
    fn get_y(&self) -> &str {
        self.y
    }
    // 这里有两个输入生命周期, 具有参数&self, 因此返回值被赋予了&self的生命周期
    fn x_eq(&self, z: &str) -> bool {
        self.x == z
    }
}

pub fn run() {
    let comment: Vec<&str> =
        "x is a param for rust lifetime demo.y is another param for rust lifetime demo"
            .split(".")
            .collect();
    let x = comment[0];
    let y = comment[1];
    let z: &'static str = "all string literal have a static lifetime";
    let anno = Annotation { x, y };

    println!(
        "the longest: \"{}\"",
        longest(anno.get_x(), anno.get_y(), z)
    );
    println!("anno.x == anno.y : {}", anno.x_eq(z));
}
