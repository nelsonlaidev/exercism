pub fn reply(message: &str) -> &str {
    let message = message.trim();

    match message {
        "" => "Fine. Be that way!",
        _ if is_question(message) && is_yelling(message) => "Calm down, I know what I'm doing!",
        _ if is_question(message) => "Sure.",
        _ if is_yelling(message) => "Whoa, chill out!",
        _ => "Whatever.",
    }
}

fn is_question(message: &str) -> bool {
    message.ends_with("?")
}

fn is_yelling(message: &str) -> bool {
    if let Some(_) = message.find(|x: char| x.is_alphabetic()) {
        return message.to_uppercase() == message;
    }

    false
}
