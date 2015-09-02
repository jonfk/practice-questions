extern crate rand;

use rand::Rng;

#[derive(Debug,Clone)]
enum Choice {
    Sheep,
    Car,
}

fn main() {
    let switch = true;
    let mut win = 0;
    let mut lose = 0;
    for i in 0..10000 {
        let choice = simulate(switch);
        match choice {
            Choice::Car => {
                win += 1;
            },
            Choice::Sheep => {
                lose += 1;
            },
        }
    }
    println!("Wins: {}\nLosses: {}", win, lose);

    //println!("Hello, world! {:?}", choice);
}

fn generate_doors(rng: &mut Rng) -> Vec<Choice> {
    vec![Choice::Sheep, Choice::Car, Choice::Sheep]
}

fn simulate(switch: bool) -> Choice {
    let mut rng = rand::thread_rng();
    let mut choice = (rng.gen::<i32>() % 3).abs();
    let mut doors = generate_doors(&mut rng);

    if choice == 2 {
        doors.remove(0);
        choice = 1;
    } else {
        doors.remove(2);
    }

    if switch {
        choice = (choice - 1).abs();
    }
    (*doors.get(choice as usize).unwrap()).clone()
}
