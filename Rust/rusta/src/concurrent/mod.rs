// --安全传递:
// Send 标记 trait 表明类型的所有权可以在多个线程间传递。
// 几乎所有基本类型都是 Send 的, Rc<T>、裸指针除外;
// 完全由 Send 的类型组成的类型也会自动被标记为 Send;

// --安全访问:
// Sync 标记 trait 表明类型可以安全的在多个线程中拥有其值的引用。
// 基本类型是 Sync 的，完全由 Sync 的类型组成的类型也是 Sync 的。

// 通常并不需要手动实现 Send 和 Sync trait, 手动实现 Send 和 Sync 是不安全的

pub mod channel;
pub mod mutex;
pub mod thread;
pub mod condvar;
mod atomic;
// pub mod semaphore;