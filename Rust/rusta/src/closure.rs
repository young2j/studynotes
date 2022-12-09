// 闭包捕获环境的方式与函数获取参数的三种方式相同:
// 1. 获取所有权-- 通过FnOnce trait 实现
//    闭包捕获变量必须获取其所有权，被捕获的变量会移动进闭包。Once代表捕获只能发生一次，闭包只能被调用一次。
// 2. 可变借用--通过FnMut trait实现
//    获取可变借用值，闭包的调用可以改变捕获的变量值。
// 3. 不可变借用--通过Fn trait实现
//    获取不可变借用值，闭包的调用不能改变捕获的变量值

struct Cacher<T, U>
where
    T: Fn(u32) -> u32,
    U: Fn(Vec<u32>) -> u32,
{
    calculation: T,
    value: Option<u32>,
    sumer: U,
    sum: Option<u32>,
}

impl<T, U> Cacher<T, U>
where
    T: Fn(u32) -> u32,
    U: Fn(Vec<u32>) -> u32,
{
    fn new(calculation: T, sumer: U) -> Cacher<T, U> {
        Cacher {
            calculation,
            value: None,
            sumer,
            sum: None,
        }
    }

    fn value(&mut self, arg: u32) -> u32 {
        match self.value {
            Some(v) => v,
            None => {
                let v = (self.calculation)(arg);
                self.value = Some(v);
                v
            }
        }
    }
    fn sum(&mut self, arg: Vec<u32>) -> u32 {
        match self.sum {
            Some(v) => v,
            None => {
                let v = (self.sumer)(arg);
                self.sum = Some(v);
                v
            }
        }
    }
}

pub fn run() {
    let y = 32;

    // 闭包会捕获环境--捕获y的值然后保存起来
    let f = |x: u32| {
        println!("do: {}+{}", x, y);
        x + y
    };

    let ff = move |v: Vec<u32>| v.iter().sum();

    let mut cache = Cacher::new(f, ff);
    let v = cache.value(y);
    println!("result: {}", v);

    let vect = vec![1, 2, 3];
    let sum = cache.sum(vect);
    println!("sum result: {}", sum);
}
