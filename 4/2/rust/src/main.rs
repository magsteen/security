fn main() {
    let start_value = "Foo<&>Bar".to_string();
    println!("Initial: {}", start_value);

    let end_value = swap(start_value);
    println!(" Result: {}", end_value);
}

fn swap(input: String) -> String {
    return input.replace("&", "&amp;").replace("<", "&lt;").replace(">", "&gt;");
}
