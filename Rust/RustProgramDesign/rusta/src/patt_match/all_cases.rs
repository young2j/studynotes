// case 1
fn use_match() {
    let x = 5;
    match x {
        1 | 2 => println!("match or"),
        3..=4 => println!("match range"),
        _ => println!("{}", x),
    }
}

// case 2
fn use_if_let() {
    let opt = Option::Some(5);
    let num: Result<u8, _> = "5".parse();
    if let Some(v) = opt {
        println!("{}", v);
    } else if let Ok(v) = num {
        println!("{}", v);
    }
}

// case 3
fn use_while_let() {
    let mut stack = vec![1, 2, 3];
    while let Some(v) = stack.pop() {
        println!("{}", v);
    }
}

// case 4
fn use_for() {
    let list = vec![1, 2, 3];
    for (index, value) in list.iter().enumerate() {
        println!("index={} value={}", index, value);
    }
}

// case 5
fn use_let() {
    let (_a, _b, _c) = (1, "2", String::from("3"));
}

// case 6
fn use_func() {
    let f = |&(x, y)| (x, y);
    f(&(1, 2));
}
