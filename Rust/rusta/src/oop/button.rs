use super::draw::Draw;

#[derive(Debug)]
pub struct Button {
    width: u32,
    height: u32,
    label: String,
}

impl Button {
   pub fn new(width: u32, height: u32, label: String) -> Button {
        Button {
            width,
            height,
            label,
        }
    }
}

impl Draw for Button {
    fn draw(&self) {
        println!("draw button:{:#?}", &self);
    }
}
