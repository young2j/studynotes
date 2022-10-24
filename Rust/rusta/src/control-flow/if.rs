fn main() {
    let number = 3;
    if number > 5 {
        println!("number {number} is greater than 5");
    } else if number < 3 {
        println!("number {number} is less than 3");
    } else {
        println!("number {number} is in range [3,5] ");
    }

    let eq3 = if number == 3 { true } else { false };
    println!("number eq3: {eq3}")
}
