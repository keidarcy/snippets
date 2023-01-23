pub fn show() {
    let my_number = 9;
    let my_reference = &my_number;
    // println("{}", my_number == my_reference); can't compare `{integer}` with `&{integer}`
    println!("{}", my_number == *my_reference); // true
    println!("{}", 9 == **&&9);

    // & <> *
    // && <> **
}
