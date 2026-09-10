pub fn square(s: u32) -> u64 {
    let base: u64 = 2;

    base.pow(s - 1)
}

pub fn total() -> u64 {
    (1..=64).into_iter().map(|x| square(x)).sum()
}
