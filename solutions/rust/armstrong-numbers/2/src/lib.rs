pub fn is_armstrong_number(num: u32) -> bool {
    let s = num.to_string();
    let mut sum = 0;

    for c in s.chars() {
        sum += c.to_digit(10).unwrap().pow(s.chars().count() as u32);
    }

    sum == num
}
