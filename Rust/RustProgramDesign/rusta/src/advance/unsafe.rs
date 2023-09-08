// 五类unsafe操作: 不会被编译器检查内存安全, 由开发者确保代码以有效的方式访问内存
// 1. 解引用裸指针
//    a. 裸指针(raw pointer)
//      - 不可变裸指针: *const T
//      - 可变罗指针: *mut T

//    b. 裸指针与引用和智能指针的区别在于：
//      - 允许忽略借用规则，可以同时拥有多个不可变和可变的指针
//      - 不保证指向有效的内存
//      - 允许为空
//      - 不能实现任何自动清理功能

//!   c. 不能在安全的代码块中解引用裸指针

//    d. 裸指针的使用场景
//      - 调用 C 代码接口
extern "C" {
    // extern 块中声明的函数在 Rust 代码中总是不安全的
    fn abs(input: i32) -> i32;
}
#[no_mangle] //告诉 Rust 编译器不要 mangle 此函数的名称
pub extern "C" fn call_rust_in_c() {
    println!(r#"call rust function in c program"#);
}

//      - 构建借用检查器无法理解的不安全代码的安全抽象
use std::slice;
fn safe_split() {
    let mut v = vec![1, 2, 3, 4, 5, 6];
    let r = &mut v[..];
    let (v1, v2) = r.split_at_mut(3); // std提供的
    assert_eq!(v1, &mut [1, 2, 3]);
    assert_eq!(v2, &mut [4, 5, 6]);
}
fn unsafe_split() {
    let mut v = vec![1, 2, 3, 4, 5, 6];
    let ptr = v.as_mut_ptr();
    let mid = v.len() / 2;
    let (v1, v2) = unsafe {
        // 自行实现的
        (
            slice::from_raw_parts_mut(ptr, mid),
            slice::from_raw_parts_mut(ptr.add(mid), v.len() - mid),
        )
    };
    assert_eq!(v1, &mut [1, 2, 3]);
    assert_eq!(v2, &mut [4, 5, 6]);
}

// 2. 调用不安全的函数或方法
unsafe fn dangerous() {}

// 3. 访问或修改可变静态变量--都是不安全的
static mut WORD_COUNT: u32 = 0;
fn use_mut_static() {
    let incr = 10;
    unsafe {
        WORD_COUNT += incr;
    }
    unsafe {
        println!("WORD_COUNT={}", WORD_COUNT);
    }
}
//  常量和不可变静态变量
//     - 常量是不可变的；常量可以在任意作用域内声明；常量只能设置为常量表达式；常量类型必须注明。
//     - 静态变量可以是可变的；静态变量是全局的；静态变量也必须注明类型。
//  区别:
//     - 可变性
//     - 静态变量值有一个固定的内存地址, 使用这个值总是会访问相同的地址。常量允许在任何被用到的时候复制其数据。
const _CONST_COUNTER: u32 = 1;
static _STATIC_COUNTER: u32 = 1;

// 4. 实现不安全的 trait
unsafe trait Foo {}
unsafe impl Foo for i32 {}

// 5. 访问 union 的字段
// 联合体主要用于和 C 代码中的联合体交互。访问联合体的字段是不安全的。

#[allow(unused)]
pub fn run() {
    let mut num = 5;
    // as强转不可变裸指针
    let r1 = &num as *const i32;
    // as强转可变裸指针
    let r2 = &mut num as *mut i32;

    // 不能确定有效性的裸指针, 可能产生片段错误(segmentation fault)
    let addr = 0x012345usize;
    let _r3 = addr as *const i32;

    unsafe {
        // 只能在unsafe代码块中解引用裸指针
        println!("r1 value = {}", *r1);
        println!("r2 value = {}", *r2);

        // 调用不安全的函数或方法
        dangerous();
    }

    // 不安全代码的安全抽象
    safe_split();
    unsafe_split();

    // 调用c代码
    unsafe {
        println!("call abs(-100) from c: {}", abs(-100));
    }

    // 访问、修改可变静态变量
    use_mut_static();
}
