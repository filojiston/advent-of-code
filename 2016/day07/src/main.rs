use fancy_regex::Regex;
use lazy_static::lazy_static;

lazy_static! {
    static ref IS_ABBA_PATTERN: Regex = Regex::new(r"(.)(?!\1)(.)\2\1").unwrap();
}

fn main() {
    let input = include_str!("input.txt");

    let mut tls_count = 0;
    let mut ssl_count = 0;
    for line in input.lines() {
        let (words_inside, words_outside) = split_line(line);
        tls_count += if is_tls(&words_inside, &words_outside) {
            1
        } else {
            0
        };
        ssl_count += if is_ssl(&words_inside, &words_outside) {
            1
        } else {
            0
        };
    }

    println!("Part1: {}", tls_count);
    println!("Part2: {}", ssl_count);
}

fn split_line(line: &str) -> (Vec<String>, Vec<String>) {
    let mut words_inside = Vec::new();
    let mut words_outside = Vec::new();

    let mut word = String::new();
    for ch in line.chars() {
        if ch == '[' {
            words_outside.push(word);
            word = String::new();
        } else if ch == ']' {
            words_inside.push(word);
            word = String::new();
        } else {
            word.push(ch);
        }
    }
    words_outside.push(word);
    (words_inside, words_outside)
}

fn is_tls(words_inside: &[String], words_outside: &[String]) -> bool {
    for word in words_inside {
        if is_abba(word) {
            return false;
        }
    }

    for word in words_outside {
        if is_abba(word) {
            return true;
        }
    }

    false
}

fn is_ssl(words_inside: &[String], words_outside: &[String]) -> bool {
    let abas = get_abas(&words_outside.join(""));
    for word in words_inside {
        for aba in &abas {
            if is_bab(word, aba) {
                return true;
            }
        }
    }

    false
}

fn is_abba(word: &str) -> bool {
    return IS_ABBA_PATTERN.is_match(word).unwrap();
}

fn get_abas(word: &str) -> Vec<String> {
    let mut abas = Vec::new();
    for i in 0..word.len() - 2 {
        let a = word.chars().nth(i).unwrap();
        let b = word.chars().nth(i + 1).unwrap();
        let c = word.chars().nth(i + 2).unwrap();
        if a == c && a != b {
            abas.push(format!("{}{}{}", a, b, c));
        }
    }
    abas
}

fn is_bab(word: &str, aba: &str) -> bool {
    let a = aba.chars().nth(0).unwrap();
    let b = aba.chars().nth(1).unwrap();
    let bab = format!("{}{}{}", b, a, b);
    word.contains(&bab)
}
