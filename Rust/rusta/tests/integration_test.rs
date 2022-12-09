// 二进制 crate 的集成测试
// 二进制 crate 不能在 tests 目录创建集成测试。
// 只有库 crate 才会向其他 crate 暴露可供调用和使用的函数；二进制 crate 只意在单独运行。
// 这就是许多 Rust 二进制项目使用一个简单的 src/main.rs 调用 src/lib.rs 中的逻辑的原因之一。

mod common;

#[test]
fn demo_test() {
    common::setup();
}
