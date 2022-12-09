
pub fn run() {
    let arr = [1; 5];
    let mut i = 0;
    while i < 5 {
        let elem = arr[i];
        println!("arr[i]={elem}");
        i += 1;
    }
}
