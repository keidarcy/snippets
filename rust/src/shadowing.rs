pub fn show() {
    let my_number = 9; // This is an i32
    println!("{}", my_number);
    {
        let my_number = 9.2;
        println!("{}", my_number);
    }
    println!("{}", my_number);
}
