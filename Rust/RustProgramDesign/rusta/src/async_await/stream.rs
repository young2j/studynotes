#![allow(unused)]
// 一次处理一个值的模式: next, try_next
// 迭代一个Stream: 使用map，filter，fold方法、遇到错误提前返回的版本 try_map，try_filter，try_fold
// 并发处理多个值的方式： for_each_concurrent 或 try_for_each_concurrent

use futures::Stream;
use std::{io, pin::Pin};

async fn sum_with_next(mut stream: Pin<&mut dyn Stream<Item = i32>>) -> i32 {
    use futures::stream::StreamExt; // 引入 next
    let mut sum = 0;
    while let Some(item) = stream.next().await {
        sum += item;
    }
    sum
}

async fn sum_with_try_next(
    mut stream: Pin<&mut dyn Stream<Item = Result<i32, io::Error>>>,
) -> Result<i32, io::Error> {
    use futures::stream::TryStreamExt; // 引入 try_next
    let mut sum = 0;
    while let Some(item) = stream.try_next().await? {
        sum += item;
    }
    Ok(sum)
}

async fn jump_around(
    mut stream: Pin<&mut dyn Stream<Item = Result<u8, io::Error>>>,
) -> Result<(), io::Error> {
    use futures::stream::TryStreamExt; // 引入 `try_for_each_concurrent`
    const MAX_CONCURRENT_JUMPERS: usize = 100;

    stream
        .try_for_each_concurrent(MAX_CONCURRENT_JUMPERS, |num| async move { Ok(()) })
        .await?;

    Ok(())
}
