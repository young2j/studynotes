// 函数指针实现了所有三个闭包 trait（Fn、FnMut 和 FnOnce），
// 所以函数调用总是可以传递函数指针替代闭包作为参数。
fn func_do(arg: i32) {
    println!("receive arg={}", arg)
}
fn use_func_do(f: fn(i32), arg: i32) {
    f(arg)
}
fn use_closure_do(f: Box<dyn Fn(i32)>, arg: i32) {
    f(arg)
}

// 返回闭包
#[allow(unused)]
fn returns_closure() -> Box<dyn Fn(i32) -> i32> {
    Box::new(|x| x + 1)
}

#[allow(unused)]
pub fn run() {
    use_func_do(func_do, 1);
    use_closure_do(Box::new(|x| func_do(x)), 2);

    // 使用闭包
    let list_of_numbers1 = vec![1, 2, 3];
    let list_of_strings1: Vec<String> = list_of_numbers1.iter().map(|i| i.to_string()).collect();

    // 使用函数
    let list_of_numbers2 = vec![1, 2, 3];
    let list_of_strings2: Vec<String> = list_of_numbers2.iter().map(ToString::to_string).collect();

    // 枚举示例
    enum Status {
        Value(u32),
        Stop,
    }

    let list_of_statuses: Vec<Status> = (0u32..20).map(Status::Value).collect();
}
