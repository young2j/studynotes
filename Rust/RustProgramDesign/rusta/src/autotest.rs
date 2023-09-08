#[derive(Debug, PartialEq)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width >= other.width && self.height >= other.height
    }
    fn can_panic(&self, short_msg: bool) {
        if short_msg {
            panic!("panic a short message")
        } else {
            panic!("panic a very very very long message")
        }
    }
}

#[test]
fn larger_can_hold_smaller() {
    let larger = Rectangle {
        width: 100,
        height: 100,
    };
    let smaller = Rectangle {
        width: 100,
        height: 50,
    };
    assert!(
        larger.can_hold(&smaller),
        "the larger can not hold the smaller"
    );
}

#[test]
fn cmp_rectangle() {
    let rect1 = Rectangle {
        width: 100,
        height: 100,
    };
    let rect2 = Rectangle {
        width: 100,
        height: 100,
    };

    assert_eq!(rect1, rect2); // 需要派生PartialEq trait, 才能比较
}

#[test]
#[should_panic]
fn short_panic() {
    let rect = Rectangle {
        width: 100,
        height: 100,
    };
    rect.can_panic(true); // 只有panic时，测试才会通过
}

#[test]
#[should_panic(expected = "long message")]
fn long_panic() {
    let rect = Rectangle {
        width: 100,
        height: 100,
    };
    rect.can_panic(false); // 只有panic抛出expected包含的信息时，测试才会通过
}

#[test]
#[ignore]
fn ingore_test() {}
