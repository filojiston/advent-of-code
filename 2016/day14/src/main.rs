use fancy_regex::Regex;
use lazy_static::lazy_static;
use std::collections::HashMap;

const INPUT: &str = "ngcjuoqr";

lazy_static! {
    static ref TRIPLE_IN_A_ROW: Regex = Regex::new(r"(.)\1{2}").unwrap();
}

enum Part {
    Part1,
    Part2,
}

fn main() {
    let part1 = find_64th_index(&Part::Part1);
    println!("Part1: {}", part1);

    let part2 = find_64th_index(&Part::Part2);
    println!("Part2: {}", part2);
}

fn find_64th_index(part: &Part) -> usize {
    let mut count = 0;
    let mut index = 0;
    let mut cache: HashMap<String, String> = HashMap::new();
    while count < 64 {
        let to_hash = format!("{}{}", INPUT, index);
        let hash = calculate_hash(&to_hash, part);
        if let Some(c) = find_triple(&hash) {
            if find_quintuple(&mut cache, c, index + 1, part) {
                count += 1;
            }
        }
        index += 1;
    }
    index - 1
}

fn find_triple(hash: &str) -> Option<char> {
    let caps = TRIPLE_IN_A_ROW.captures(hash).unwrap();
    if let Some(c) = caps {
        return Some(c.get(1).unwrap().as_str().chars().next().unwrap());
    }

    None
}

fn find_quintuple(cache: &mut HashMap<String, String>, c: char, index: usize, part: &Part) -> bool {
    for i in index..index + 1001 {
        let to_hash = format!("{}{}", INPUT, i);
        let hash = cache
            .entry(to_hash.clone())
            .or_insert_with(|| calculate_hash(&to_hash, part));
        if hash.contains(&format!("{}{}{}{}{}", c, c, c, c, c)) {
            return true;
        }
    }
    false
}

fn hash_2017(hash: &str) -> String {
    let mut hash = hash.to_string();
    for _ in 0..2017 {
        hash = format!("{:x}", md5::compute(hash));
    }
    hash
}

fn calculate_hash(input: &str, part: &Part) -> String {
    match part {
        Part::Part1 => format!("{:x}", md5::compute(input)),
        Part::Part2 => hash_2017(input),
    }
}
