pub fn show() {
    let my_number = {
        let second_number = 9;
        second_number + 9;
    };

    //   = help: the trait `std::fmt::Display` is not implemented for `()`
    // = note: in format strings you may be able to use `{:?}` (or {:#?} for pretty-print) instead
    // println!("My number is: {}", my_number); 
    println!("My number is: {:?}", my_number);  // debug print
    println!("My number is: {:#?}", my_number);  // pretty-print
}
