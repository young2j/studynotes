//! # rusta
//! `//!`风格的注释通常用于crate或模块根文件, 为 crate或模块整体提供文档。

/// A demo function for doc comment
/// # Panics
/// Panics：这个函数可能会 panic! 的场景。并不希望程序崩溃的函数调用者应该确保他们不会在这些情况下调用此函数。
/// # Errors
/// Errors：如果这个函数返回 Result，此部分描述可能会出现何种错误以及什么情况会造成这些错误，这有助于调用者编写代码来采用不同的方式处理不同的错误。
/// # Safety
/// Safety：如果这个函数使用 unsafe 代码（这会在第十九章讨论），这一部分应该会涉及到期望函数调用者支持的确保 unsafe 块中代码正常工作的不变条件（invariants）。
/// # Examples
/// ```
///  let func_name = "demo_func";
///  demo_func(func_name);
/// ```
pub fn demo_func(arg: &str) {
    println!("{}: just a demo for  doc comment introduction", arg)
}