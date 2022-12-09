// 自定义迭代器--实现Iterator trait的next方法
struct Counter {
    count: u32,
}

impl Counter {
    fn new() -> Counter {
        Counter { count: 0 }
    }
}

impl Iterator for Counter {
    type Item = u32;
    fn next(&mut self) -> Option<Self::Item> {
        if self.count < 5 {
            self.count += 1;
            Some(self.count)
        } else {
            None
        }
    }
}
// ---

pub fn run() {
    // 不可变引用迭代器
    let v1 = vec![1, 2, 3];
    let immut_iter = v1.iter();
    for v in immut_iter {
        println!("{v}")
    }
    println!();

    // 可变引用迭代器
    let mut v2 = vec![1, 2, 3];
    let mut_iter = v2.iter_mut();
    for v in mut_iter {
        println!("{v}")
    }
    println!();

    // 移动迭代器，v3失去所有权
    let v3 = vec![1, 2, 3];
    let into_iter = v3.into_iter();
    for v in into_iter {
        println!("{v}")
    }
    println!();

    // 生成新的迭代器
    let v4 = vec![1, 2, 3];
    let new_iter = v4.iter().map(|x| x * x).filter(|y| y % 2 == 0);
    for v in new_iter {
        println!("{}", v)
    }
    println!();

    // 自定义迭代器
    let counter = Counter::new();
    let count_iter = counter.into_iter();
    for v in count_iter {
        println!("{v}")
    }
    println!();

    // [1,2,3,4,5] [1,2,3,4,5]
    // [(1,2),(2,3),(3,4),(4,5)]
    // [3,5,7,9]
    // 7+9 = 16
    let sum: u32 = Counter::new()
        .zip(Counter::new().skip(1))
        .map(|(x, y)| x + y)
        .filter(|z| z > &5)
        .sum();
    println!("{sum}")
}
