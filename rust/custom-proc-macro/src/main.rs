extern crate custom_derive;

use custom_derive::{log_att, make_hello, Hello};

make_hello!(world);
make_hello!(张三);

#[log_att(struct, "world")]
struct Hello {
    pub name: String,
}
#[log_att(func, "test")]
fn invoked() {}

#[derive(Hello)]
struct World;

fn main() {
    hello_world();
    hello_张三();
    World::hello_macro_derive();
    invoked();
}
