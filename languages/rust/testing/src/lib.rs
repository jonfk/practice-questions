//! The `adder` crate provides functions that add numbers to other numbers.
//!
//! # Examples
//!
//! ```
//! assert_eq!(4, testing::add_two(2));
//! ```
//!
//! anotherone
//!
//! ```
//! println!("aoeu");
//! ```

/// This function adds two to its argument.
///
/// # Examples
///
/// ```
/// use testing::add_two;
///
/// assert_eq!(4, add_two(2));
/// ```
///
/// # Panics
///
/// ```should_panic
/// use testing::add_two;
///
/// assert_eq!(4, add_two(1));
/// ```
pub fn add_two(a: i32) -> i32 {
    a + 2
}


#[cfg(test)]
mod tests {

    use super::add_two;

    #[test]
    fn it_works() {
        assert_eq!(4, add_two(2));
    }

    #[test]
    #[should_panic]
    fn expect_panic_works() {
        assert!(false);
    }

    #[test]
    #[should_panic(expected = "assertion failed")]
    fn expect_panic_eq_works() {
        assert_eq!("hello","world");
    }
}
