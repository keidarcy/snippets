pub fn show() {
    let father_name = "John";
    let son_name = "Jack";
    let family_name = "Doe";
    println!(
        "This is {0} {2}, son of {1} {2}",
        son_name, father_name, family_name
    );
    println!(
        "This is {son} {family}, son of {father} {family}",
        son = son_name,
        father = father_name,
        family = family_name
    );

    // system {variable: padding alignment minimum.maximum}
    let letter = "a";
    println!("{:あ^5.}", letter);

    // example
    // print beautiful table
    let title = "Title";
    println!("{:-^40}", title);
    let bar = "|";
    println!("{} {:<20} {} {:>10} {}", bar, "Name", bar, "Age", bar);
    println!("{} {:<20} {} {:>10} {}", bar, "John", bar, 20, bar);
    println!("{} {:<20} {} {:>10} {}", bar, "Jack", bar, 10, bar);
    println!("{} {:<20} {} {:>10} {}", bar, "Jill", bar, 15, bar);
    println!("{:-^40}", "");
}
