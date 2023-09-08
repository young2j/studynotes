// mpsc: mutiple producer, single consumer
use std::{sync::mpsc, thread};

#[allow(unused)]
pub fn run() {
    let (tx, rx) = mpsc::channel();

    let tx1 = tx.clone();
    let tx2 = tx.clone();

    thread::spawn(move || {
        let val = "send from tx";
        tx.send(val).unwrap();
    });

    thread::spawn(move || {
        let val = "send from tx1";
        tx1.send(val).unwrap();
    });

    thread::spawn(move || {
        let val = "send from tx2";
        tx2.send(val).unwrap();
    });

    for val in rx {
        println!("{}", val);
    }
}
