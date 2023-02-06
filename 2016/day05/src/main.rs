use std::vec;

use md5;

const INPUT: &str = "uqwqemis";

fn main() {
    let mut found_part1 = 0;
    let mut found_part2 = 0;
    let mut password_part1 = String::new();
    let mut password_part2 = vec!['-'; 8];
    let mut current_index = 0;
    while found_part1 < 8 || found_part2 < 8 {
        let current = format!("{}{}", INPUT, current_index.to_string());
        let hash = format!("{:x}", md5::compute(current.clone()));
        if hash.starts_with("00000") {
            if found_part1 < 8 {
                password_part1.push(hash.chars().nth(5).unwrap());
                found_part1 += 1;
            }

            let position = hash.chars().nth(5).unwrap().to_digit(10);
            if position.is_some() && position.unwrap() < 8 {
                let position = position.unwrap() as usize;
                if password_part2[position] == '-' {
                    password_part2[position] = hash.chars().nth(6).unwrap();
                    found_part2 += 1;
                }
            }
        }
        current_index += 1
    }

    println!("Part1: {}", password_part1);
    println!("Part2: {}", password_part2.iter().collect::<String>());
}
