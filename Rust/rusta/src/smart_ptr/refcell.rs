// RefCell<T> 主要用于要求内部可变性的场景
// RefCell<T>.borrow()：返回一个Ref<T>智能指针，代表不可变借用。
//    当调用borrow()时，RefCell<T>内部活跃的不可变借用数+1；
//    当Ref<T>离开作用域时，RefCell<T>内部活跃的不可变借用数-1。
//    RefCell<T>活跃的不可变借用数可以有多个。
// RefCell<T>.borrow_mut()：返回一个RefMut<T>智能指针，代表可变借用。
//    当调用borrow_mut()时，RefCell<T>内部活跃的可变借用数+1；
//    当RefMut<T>离开作用域时，RefCell<T>内部活跃的可变借用数-1。
//    (在相同的作用域中)RefCell<T>活跃的可变借用数同时只能存在一个！否则会在运行时panic!

pub trait Messenger {
    fn send(&self, msg: &str);
}

pub struct LimitTracker<'a, T: Messenger> {
    messenger: &'a T,
    value: usize,
    max: usize,
}

impl<'a, T> LimitTracker<'a, T>
where
    T: Messenger,
{
    pub fn new(messenger: &T, max: usize) -> LimitTracker<T> {
        LimitTracker {
            messenger,
            value: 0,
            max,
        }
    }

    pub fn set_value(&mut self, value: usize) {
        self.value = value;

        let percentage_of_max = self.value as f64 / self.max as f64;

        if percentage_of_max >= 1.0 {
            self.messenger.send("Error: You are over your quota!");
        } else if percentage_of_max >= 0.9 {
            self.messenger
                .send("Urgent warning: You've used up over 90% of your quota!");
        } else if percentage_of_max >= 0.75 {
            self.messenger
                .send("Warning: You've used up over 75% of your quota!");
        }
    }
}

#[cfg(test)]
mod tests {
    use std::cell::RefCell;

    use super::*;

    struct MockMessenger {
        sent_messages: RefCell<Vec<String>>,
    }

    impl MockMessenger {
        fn new() -> MockMessenger {
            MockMessenger {
                sent_messages: RefCell::new(vec![]),
            }
        }
    }

    impl Messenger for MockMessenger {
        fn send(&self, message: &str) {
            self.sent_messages.borrow_mut().push(String::from(message));
        }
    }

    #[test]
    fn it_sends_an_over_75_percent_warning_message() {
        let mock_messenger = MockMessenger::new();
        let mut limit_tracker = LimitTracker::new(&mock_messenger, 100);

        limit_tracker.set_value(80);

        assert_eq!(mock_messenger.sent_messages.borrow().len(), 1);
    }
}
