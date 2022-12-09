// 1. 类型别名 type alias
// 类型别名的主要用途是减少重复
type _Thunk = Box<dyn Fn() + Send + 'static>;

// 2. 从不返回的类型 never type !
// fn never_return() -> ! {
//    ...
// }
// - 在函数从不返回的时候充当返回值
// - continue、 panic! 、loop表达式的值是 !
// - never type 可以强转为任何其他类型

// 3. 动态大小类型(DST, dynamically sized types)和 Sized trait
// - str 是一个 DST, 直到运行时都不知道字符串有多长，大小是多少; 只能创建&str类型的变量，&str存储了 str 的地址和其长度
// - trait 也是一个DST, 每一个 trait 都是一个可以通过 trait 名称来引用的动态大小类型
// - Sized trait: 让编译器在编译时知道类型的大小
// - 泛型函数默认只能用于在编译时已知大小的类型
fn _generic_func<T>(_: T) {} // 等同于
fn _sized_generic_func<T: Sized>(_: T) {}
// 语法?Sized 只能用于Sized trait
// 因为T肯能是Sized也可能不是Sized，所以arg类型只能是&T
fn _may_sized_generic_func<T: ?Sized>(_: &T) {}
