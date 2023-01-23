pub fn show() {
    // let name = "☺️"; // This is a &str
    // let other_name = String::from("Hello"); // String is a heap-allocated string
    // let n = String::from("Hello");
    // let other_string = "あああ".to_string(); // This is a String

    // &str and String
    // println!(
    //     "A String is always {:?} bytes. It is Sized.",
    //     std::mem::size_of::<String>()
    // ); // std::mem::size_of::<Type>() gives you the size in bytes of a type
    // println!(
    //     "And an i8 is always {:?} bytes. It is Sized.",
    //     std::mem::size_of::<i8>()
    // );
    // println!(
    //     "And an f64 is always {:?} bytes. It is Sized.",
    //     std::mem::size_of::<f64>()
    // );
    // println!(
    //     "But a &str? It can be anything. 'あ' is {:?} bytes. It is not Sized.",
    //     std::mem::size_of_val("あ")
    // ); // std::mem::size_of_val() gives you the size in bytes of a variable
    // println!(
    //     "And '🐶' is {:?} bytes. It is not Sized.",
    //     std::mem::size_of_val("🐶")
    // );

    // format!() marco
    let a = "a";
    let b = "b";
    let c = format!("{} and {}", a, b);
    println!("{}", c);
}
