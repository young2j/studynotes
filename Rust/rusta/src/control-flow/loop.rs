fn main() {
    let mut number = 3;
    let mut counter = 2;
    // loop标签
    'outer: loop {
        println!("number={number}");
        number -= 1;
        if number <= 0 {
            break;
        } else {
            let result = loop {
                counter -= 1;
                if counter == 0 {
                    break counter + 100; // loop-break返回值
                } else if counter < 0 {
                    break 'outer; // break外层循环
                }
            };
            println!("result={result}");
        }
    }
}
