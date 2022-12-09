use std::fmt::Display;

use super::implfor::Summary;

// ====================================
// trait bound
pub fn notify1<T: Summary>(item: &T) {
    println!("Notice-> {}", item.summarize());
}

// notify1的语法糖
pub fn notify2(item: &impl Summary) {
    println!("Notice-> {}", item.summarize());
}

// ====================================
// 类型相同的trait bound
pub fn notify3<T: Summary>(item1: &T, item2: &T) {
    println!("Notice-> {}", item1.summarize());
    println!("Notice-> {}", item2.summarize());
}
// 同notify3，但是更冗长
pub fn notify4(item1: &impl Summary, item2: &impl Summary) {
    println!("Notice-> {}", item1.summarize());
    println!("Notice-> {}", item2.summarize());
}

// ====================================
// 使用 + 指定多个trait bound
pub fn notify5<T: Summary + Display>(item: &T) {
    println!("Notice-> {}", item.summarize());
}
// 同notify5
pub fn notify6(item: &(impl Summary + Display)) {
    println!("Notice-> {}", item.summarize());
}
// 使用where，同notify5，notify6
pub fn notify<T>(item: &T)
where
    T: Summary + Display,
{
    println!("Notice-> {}", item.summarize());
}
