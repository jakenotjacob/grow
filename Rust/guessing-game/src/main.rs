//Prelude
use std::io;


fn main() {
    println!("Guess the number!");

    //Create mutatable variable bound to new String instance
    // 'new' is an associated function
    let mut guess = String::new();

    io::stdin().read_line(&mut guess)
        .expect("Failed to read line");

    println!("You guessed: {}", guess);
}
