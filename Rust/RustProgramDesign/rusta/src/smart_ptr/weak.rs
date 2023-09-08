// Weak(T): 弱引用通常用于双向引用的场景，能够避免引用循环和内存泄漏
// Rc<T>使用weak_count来记录其存在多少个Weak<T>引用。
// 当strong_count为0时即使weak_count不为0，Rc<T>实例也会被清理。

// Rc::downgrade(Rc<T>) -> Weak<T>, weak_count+1
// Rc::upgrade<Weak<T>> -> Option<Rc<T>>, 当Rc<T>实例被清理后返回None, 否则返回Some<Rc<T>>

use std::{
    cell::RefCell,
    rc::{Rc, Weak},
};

#[derive(Debug)]
struct Node {
    value: i32,
    parent: RefCell<Weak<Node>>,
    children: RefCell<Vec<Rc<Node>>>,
}

pub fn run() {
    let leaf = Rc::new(Node {
        value: 1,
        parent: RefCell::new(Weak::new()),
        children: RefCell::new(vec![]),
    });

    // 1, 0
    println!(
        "leaf strong_count={} weak_count={}",
        Rc::strong_count(&leaf),
        Rc::weak_count(&leaf)
    );

    {
        let branch = Rc::new(Node {
            value: 0,
            parent: RefCell::new(Weak::new()),
            children: RefCell::new(vec![Rc::clone(&leaf)]),
        });

        // 2, 0
        println!(
            "leaf strong_count={} weak_count={}",
            Rc::strong_count(&leaf),
            Rc::weak_count(&leaf),
        );
        // 1, 0
        println!(
            "branch strong_count={} weak_count={}",
            Rc::strong_count(&branch),
            Rc::weak_count(&branch),
        );

        // Rc<Node> -> Weak<Node>
        *leaf.parent.borrow_mut() = Rc::downgrade(&branch);

        // 2, 0
        println!(
            "leaf strong_count={} weak_count={}",
            Rc::strong_count(&leaf),
            Rc::weak_count(&leaf),
        );
        // 1, 1
        println!(
            "branch strong_count={} weak_count={}",
            Rc::strong_count(&branch),
            Rc::weak_count(&branch),
        );

        //
        let branch = leaf.parent.borrow();
        match branch.upgrade() {
            Some(v) => println!("branch:{:?}", v),
            None => println!("branch 被清理"),
        }
    }

    // branch离开作用域, 被清理
    // 1, 0
    println!(
        "leaf strong_count={} weak_count={}",
        Rc::strong_count(&leaf),
        Rc::weak_count(&leaf),
    );

    let branch = leaf.parent.borrow();
    match branch.upgrade() {
        Some(v) => println!("branch:{:?}", v),
        None => println!("branch被清理"),
    }
}
