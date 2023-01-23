pub fn show() {
    print!("\t Start with a tab \n and move to a new line");
    println!("");
    println!("");
    println!(
        "Inside quotes
    you can write over
many lines
and it will print just fine"
    );
    println!("");

    // hashtag
    let hashtag_string = r##"The hash tag "" #"##;
    println!("{}", hashtag_string);

    // bite
    println!("{:?}", b"a This will look like number");

    // Cast char as u32 to get the hexadecimal value
    println!("{:X}", 'い' as u32);
    println!("{:X}", 'a' as u32);

    // pointer, show memory address
    println!("{:p}", &9);

    let number = 'a' as u32;
    println!(
        "Binary: {:b}, hexadecimal: {:x}, octal: {:o}, decimal: {}",
        number, number, number, number
    );
    let number = 55;
    println!(
        "Binary: {:b}, hexadecimal: {:x}, octal: {:o}, decimal: {}",
        number, number, number, number
    );
}
