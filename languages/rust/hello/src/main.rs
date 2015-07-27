
fn main() {
    let a : Option<_>= (1..10).map(|x| x + 1).collect();
    assert_eq!(a, vec![2,3,4,5,6,7,8,9,10]);
}
