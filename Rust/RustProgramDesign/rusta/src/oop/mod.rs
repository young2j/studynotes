// --Trait 对象要求对象安全
// 对象安全对于 trait 对象是必须的
// 如果一个 trait 中所有的方法有如下属性时，则该 trait 是对象安全的：
//  - 返回值类型不为 Self
//  - 方法没有任何泛型类型参数

// 一个 trait 的方法不是对象安全的例子是标准库中的 Clone trait
// pub trait Clone {
//     fn clone(&self) -> Self;
// }

mod button;
mod checkbox;
mod draw;
pub mod screen;
