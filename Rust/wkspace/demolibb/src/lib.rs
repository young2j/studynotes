pub fn libb_add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod libb_tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = libb_add(2, 2);
        assert_eq!(result, 4);
    }
}
