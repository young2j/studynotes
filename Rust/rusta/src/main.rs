// mod ctl_flow;
// mod game;
// mod generics;
// mod traits;
// mod lifetime;
// mod closure;
// mod iterator;
// mod doccomment;
// pub use self::doccomment::demo_func;

// mod smart_ptr;
mod concurrent;
// mod oop;
// mod patt_match;

// mod advance;

mod async_await;

#[cfg(test)]
pub mod autotest;

#[allow(unused)]
fn main() {
    // game::guess_num();
    // ctl_flow::r#for::run();
    // ctl_flow::r#if::run();
    // ctl_flow::r#while::run();
    // ctl_flow::r#loop::run();
    // generics::run();
    // traits::main::run();
    // lifetime::run();
    // closure::run();
    // iterator::run();

    // smart_ptr::r#box::run();
    // smart_ptr::rc::run();
    // smart_ptr::weak::run();

    // concurrent::thread::run();
    // concurrent::channel::run();
    // concurrent::mutex::run();

    // oop::screen::run();

    // advance::r#unsafe::run();
    // advance::advance_trait::run();
    // advance::advance_func::run();
    // advance::macros::run();

    async_await::futures::run();


}
