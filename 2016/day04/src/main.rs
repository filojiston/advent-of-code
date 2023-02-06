use regex::Regex;
use std::collections::HashMap;

fn main() {
    let input = include_str!("input.txt");
    let pattern = Regex::new(r"(.+)-(\d+)\[(.+)\]").unwrap();

    let mut sum = 0;
    let mut decrypted_sector_id = 0;
    for line in input.lines() {
        let mut chars: HashMap<char, usize> = HashMap::new();
        let split = pattern.captures(line).unwrap();
        let encyrpted_name = split.get(1).unwrap().as_str();
        let sector_id = split.get(2).unwrap().as_str().parse::<usize>().unwrap();
        let checksum = split.get(3).unwrap().as_str();

        for c in encyrpted_name.chars() {
            if c != '-' {
                chars.insert(c, chars.get(&c).unwrap_or(&0) + 1);
            }
        }

        let sorted_chars = sort(&chars);
        let calculated_checksum = calculated_checksum(&sorted_chars);
        if calculated_checksum == checksum {
            sum += sector_id;
        }

        let decrypted_name = decrypt(encyrpted_name, sector_id);
        if decrypted_name == "northpole object storage" {
            decrypted_sector_id = sector_id;
        }
    }

    println!("Part1: {}", sum);
    println!("Part2: {}", decrypted_sector_id);
}

fn decrypt(encrypted_name: &str, sector_id: usize) -> String {
    let mut decrypted_name = String::new();

    for c in encrypted_name.chars() {
        if c == '-' {
            decrypted_name.push(' ');
        } else {
            let mut new_c = c as u8 + (sector_id % 26) as u8;
            if new_c > 122 {
                new_c -= 26;
            }
            decrypted_name.push(new_c as char);
        }
    }

    decrypted_name
}

fn sort(chars: &HashMap<char, usize>) -> Vec<(&char, &usize)> {
    let mut sorted_chars: Vec<(&char, &usize)> = chars.iter().collect();
    sorted_chars.sort_by(|a, b| {
        if a.1 == b.1 {
            a.0.cmp(b.0)
        } else {
            b.1.cmp(a.1)
        }
    });
    sorted_chars
}

fn calculated_checksum(sorted_chars: &Vec<(&char, &usize)>) -> String {
    let mut checksum = String::new();
    for i in 0..5 {
        checksum.push(*sorted_chars[i].0);
    }
    checksum
}
