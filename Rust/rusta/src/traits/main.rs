use super::bound::{notify1, notify2};
use super::implfor::Summary;
use super::implret::return_summarizable;

pub fn use_trait() {
    let tweet = return_summarizable();
    println!("tweet summary: {}", tweet.summarize());
    println!("tweet summary default: {}", tweet.summarize_default());

    notify1(&tweet);
    notify2(&tweet);
}
