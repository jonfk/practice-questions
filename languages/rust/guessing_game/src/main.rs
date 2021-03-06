extern crate rand;

use std::io;
use std::cmp::Ordering;
use rand::Rng;


fn main() {
    println!("guessing game");

    println!("Please guess a number");

    let secret_number = rand::thread_rng().gen_range(1, 101);

    println!("random {}", secret_number);

    loop {

        let mut guess = String::new();

        io::stdin().read_line(&mut guess).ok().expect("Failed to read line");

        println!("guess {}", guess);

        let guess: u32 = match guess.trim().parse::<u32>() {
            Ok(num) => num,
            Err(_) => continue,
        };

        match guess.cmp(&secret_number) {
            Ordering::Less    => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal   => {
                println!("You win!");
                break;},
        }
    }
}
