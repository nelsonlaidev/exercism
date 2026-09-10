pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    let mut combined = Vec::new();

    for factor in factors {
        if *factor == 0 {
            continue;
        }

        for n in 0..limit {
            if n % factor == 0 {
                combined.push(n);
            }
        }
    }

    combined.sort();
    combined.dedup();

    combined.iter().sum()
}
