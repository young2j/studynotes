use super::draw::Draw;

#[derive(Debug)]
pub struct CheckBox {
    width: u32,
    height: u32,
    options: Vec<String>,
}

impl CheckBox {
   pub fn new(width: u32, height: u32, options: Vec<String>) -> CheckBox {
        CheckBox {
            width,
            height,
            options,
        }
    }
}

impl Draw for CheckBox {
    fn draw(&self) {
        println!("draw checkbox {:#?}", &self);
    }
}
