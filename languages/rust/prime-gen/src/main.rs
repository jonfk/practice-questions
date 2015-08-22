
use std::env;
use std::process;

fn main() {
    let args = env::args();
    let count = args.len();
    if  count < 2 {
        println!("Please enter a limit");
        process::exit(1);
    }
    let limit = args.take(2).last().unwrap().parse::<u64>();
    match limit {
        Ok(limit) => {
            println!("{}", limit);
            let primes = sieve(limit);
            println!("Number of primes: {}", primes.len());
            //println!("{:?}", primes);
        },
        Err(_) => {
            println!("Please enter a valid number");
        },
    }

}

fn sieve(bound: u64) -> Vec<u64> {
    let mut is_prime = vec![true; bound as usize + 1];
    is_prime[0] = false; is_prime[1] = false;
    // start from 2
    for i in 2..bound {
        if is_prime[i as usize] {
            let mut j = i + i;
            while j <= bound {
                is_prime[j as usize] = false;
                j += i;
            }
        }
    }
    is_prime.iter().enumerate().filter_map(|(prime, &is_prime)| if is_prime { Some(prime as u64) } else { None }).collect::<Vec<_>>()
}
