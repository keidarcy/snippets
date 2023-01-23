pub fn show() {
    // let my_number = 10;
    // println!("{}", my_number);


    // let my_number = {
    //     let second_number = 9;
    //     second_number
    // };

    // println!("{}", my_number);

    let my_number = {
        let second_number = 9;
        second_number + 9 // No semicolon, so the code block return 8 + 9
                          // It works just like a function
    };

    println!("My number is: {}", my_number);
}
