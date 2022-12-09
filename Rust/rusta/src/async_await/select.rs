use futures::future;
use futures::pin_mut;
use futures::select;
use futures::FutureExt;

async fn task_one() { /* ... */
}
async fn task_two() { /* ... */
}

// .fuse()方法可以让 Future 实现 FusedFuture 特征,保证select 不能再对已完成的Future进行轮询使用
// pin_mut! 宏会为 Future 实现 Unpin特征，保证select不会拿走Future的所有权，该 Future 若没有完成，它的所有权还可以继续被其它代码使用
// 这两个特征恰恰是使用 select 所必须的
async fn race_tasks() {
    let t1 = task_one().fuse();
    let t2 = task_two().fuse();

    pin_mut!(t1, t2);

    select! {
        () = t1 => println!("任务1率先完成"),
        () = t2 => println!("任务2率先完成"),
    }
}

async fn race_tasks_complete() {
    let mut a_fut = future::ready(4);
    let mut b_fut = future::ready(6);
    let mut total = 0;

    loop {
        select! {
            a = a_fut => total += a,
            b = b_fut => total += b,
            complete => break,
            default => panic!(), // 该分支永远不会运行，因为`Future`会先运行，然后是`complete`
        };
    }
    assert_eq!(total, 10);
}

#[allow(unused)]
pub fn run() {
    race_tasks();
    race_tasks_complete();
}
