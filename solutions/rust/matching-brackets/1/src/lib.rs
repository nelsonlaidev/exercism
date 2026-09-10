pub fn brackets_are_balanced(string: &str) -> bool {
    let mut slots = Vec::new();

    for s in string.chars() {
        if !matches!(s, '[' | ']' | '{' | '}' | '(' | ')') {
            continue;
        }

        if matches!(s, '[' | '{' | '(') {
            slots.push(s);
            continue;
        }

        let last = slots.last();
        let latest: char;

        match last {
            Some(value) => latest = *value,
            None => return false,
        }

        if s == ']' && latest == '[' || s == '}' && latest == '{' || s == ')' && latest == '(' {
            slots.pop();
        } else {
            return false;
        }
    }

    slots.len() == 0
}
