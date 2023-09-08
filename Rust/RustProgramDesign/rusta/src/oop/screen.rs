use super::button::Button;
use super::checkbox::CheckBox;
use super::draw::Draw;

struct Screen {
    components: Vec<Box<dyn Draw>>,
}

impl Screen {
    fn run(&self) {
        for comp in self.components.iter() {
            comp.draw();
        }
    }
}

pub fn run() {
    let screen = Screen {
        components: vec![
            Box::new(Button::new(1, 2, String::from("button"))),
            Box::new(CheckBox::new(
                1,
                2,
                vec![String::from("yes"), String::from("no")],
            )),
        ],
    };
    screen.run();
}
