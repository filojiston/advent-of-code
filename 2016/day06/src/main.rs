use std::collections::HashMap;

fn main() {
    let input = include_str!("input.txt");
    let chars = input
        .lines()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut password_part1 = String::new();
    let mut password_part2 = String::new();
    for row in 0..chars[0].len() {
        let mut map: HashMap<char, usize> = HashMap::new();
        for col in 0..chars.len() {
            let c = chars[col][row];
            *map.entry(c).or_insert(0) += 1;
        }

        let max_c = map.iter().max_by_key(|&(_, v)| v).unwrap().0;
        let min_c = map.iter().min_by_key(|&(_, v)| v).unwrap().0;
        password_part1.push(*max_c);
        password_part2.push(*min_c);
    }

    println!("Part1: {}", password_part1);
    println!("Part2: {}", password_part2);
}
