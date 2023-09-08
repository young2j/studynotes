// Box::leak可以用于全局变量，例如用作运行期初始化的全局动态配置
// 它可以将一个变量从内存中泄漏,然后将其变为'static生命周期，最终该变量将和程序活得一样久

#[derive(Debug)]
struct Config {
    a: String,
    b: String
}
static mut CONFIG: Option<&mut Config> = None;

pub fn run() {
    let c = Box::new(Config {
        a: "A".to_string(),
        b: "B".to_string(),
    });

    unsafe {
        // 将`c`从内存中泄漏，变成`'static`生命周期
        CONFIG = Some(Box::leak(c));
        println!("{:?}", CONFIG);
    }
}