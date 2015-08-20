fn main() {
    let a = permutations(String::from("123456789")).iter().map(|x| x.parse::<u64>().unwrap() ).collect::<Vec<_>>();
    let l = a.iter().last().unwrap();
    println!("{:?}",l);

}

fn permutations(s: String) -> Vec<String> {
    if s.len() == 0 {
        vec![]
    } else if s.len() == 1 {
        vec![s]
    } else {
        let mut iter = s.chars();
        let c = iter.next().unwrap();

        let rest = iter.collect::<String>();
        let rest_permutations = permutations(rest);
        let mut all_permutations : Vec<String> = Vec::new();
        for perm in rest_permutations {
            let len = perm.len();
            for i in 0..len+1 {
                let mut new_perm = perm.clone();
                new_perm.insert(i, c);
                all_permutations.push(new_perm);
            }
        }
        all_permutations
    }
}
