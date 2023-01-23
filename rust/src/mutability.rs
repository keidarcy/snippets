pub fn show() {
    let mut my_number = 9;
    println!("{}", my_number);
    my_number = 10;
    // my_number = "Now I am a string";  mismatched types expected integer, found `&str`
    // let my_number = "Now I a string"; shadowing
    println!("{}", my_number);
}
