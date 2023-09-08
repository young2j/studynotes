pub fn liba_add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod liba_tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = liba_add(2, 2);
        assert_eq!(result, 4);
    }
}
