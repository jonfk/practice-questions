
fn main() {
    println!("Hello, world!");

    let mut x: i32 = 5;
    {
        let y: &mut i32 = &mut x;
        *y += 10
    }
    //add_one(&mut x);

    println!("x: {},", x);
}

fn add_one(i: &mut i32) {
    *i += 1;
}
