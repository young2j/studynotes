// 最简单直接的智能指针是 box，其类型是 Box<T>。
// box 允许你将一个值放在堆上而不是栈上, 留在栈上的则是指向堆数据的指针。

use std::{fmt::Display, ops::Deref};

// 自定义Box
struct MyBox<T: Display>(T);

impl<T> MyBox<T>
where
    T: Display,
{
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Drop for MyBox<T>
where
    T: Display,
{
    fn drop(&mut self) {
        println!("drop: {}", self.0)
    }
}

impl<T> Deref for MyBox<T>
where
    T: Display,
{
    type Target = T;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

pub fn run() {
    // 当b离开作用域时会清理栈指针和堆数据
    let b = Box::new(5);
    println!("box demo: {}", b);

    let myb = MyBox::new(5);
    // *myb 等同于 *(myb.deref())
    assert_eq!(5, *myb);

    // 使用std::mem::drop在变量离开作用域前主动清理掉
    drop(myb);

    let myb_str = MyBox::new(String::from("自定义box"));
    let handle_str = |x: &str| println!("x={}", x);

    // Deref 隐式强制转换: 均发生在编译时
    // &MyBox<String> -> &String -> &str
    handle_str(&myb_str);
}
