pub trait Draw{
    fn draw(&self);
}

pub struct Screen{
    pub compoents: Vec<Box<dyn Draw>>,
}

impl Screen{
    pub fn run(&self) {
        for compoent in self.compoents.iter() {
            compoent.draw();
        }
    }
}

pub struct Button {
    pub width: u32,
    pub height: u32,
    pub label: String,
}

impl Draw for Button{
    fn draw(&self) {
        println!("draw button {}, {}, {}", self.width, self.height, self.label);
    }
}