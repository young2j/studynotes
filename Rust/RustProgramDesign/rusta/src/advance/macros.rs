// 宏: 元编程, 生成 Rust 代码的 Rust 代码
// 1. 使用 macro_rules! 声明宏
#[macro_export] // 指定宏可以被引入作用域
macro_rules! vec {
    // 匹配表达式 "expr," 0次或多次, 并将表达式存入x
    ( $( $x:expr ),* ) => {
        {
            let mut temp_vec = Vec::new();
            $(
                temp_vec.push($x);
            )*
            temp_vec
        }
    };
}

// 2. 三种类型的过程宏（自定义派生（derive），类属性和类函数）
// 自定义derive宏 -- derive 只能用于结构体和枚举
use hello_macro_derive::HelloMacro;
pub trait HelloMacro {
    fn hello_macro();
}
#[derive(HelloMacro)]
struct Point;

pub fn run() {
    Point::hello_macro();
}

// 类属性宏--允许创建新的属性
// #[route(GET, "/")]
// fn index() {}
//
// #[proc_macro_attribute] // attr 匹配`GET, "/"`部分，item匹配`fn index() {}`部分
// pub fn route(attr: TokenStream, item: TokenStream) -> TokenStream {}

// 函数宏-可以接受未知数量的参数
// let sql = sql!(SELECT * FROM posts WHERE id=1);
// #[proc_macro]
// pub fn sql(input: TokenStream) -> TokenStream {}
