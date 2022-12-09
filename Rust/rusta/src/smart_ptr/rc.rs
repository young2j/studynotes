// Rc<T> 允许相同数据有多个所有者

enum List {
    Cons(i32, Rc<List>),
    Nil,
}

use self::List::{Cons, Nil};
use std::rc::Rc;

pub fn run() {
    let c = Rc::new(Cons(2, Rc::new(Cons(3, Rc::new(Cons(4, Rc::new(Nil)))))));
    println!("count(c)={}", Rc::strong_count(&c));
    {
        // Rc::clone会增加引用计数
        let _b = Cons(2, Rc::clone(&c));
        println!("count(c)={}", Rc::strong_count(&c));
    }
    let _a = Cons(1, Rc::clone(&c));
    println!("count(c)={}", Rc::strong_count(&c));
}
