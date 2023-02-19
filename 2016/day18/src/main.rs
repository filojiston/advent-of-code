use std::vec;

const INPUT: &str = ".^^^^^.^^^..^^^^^...^.^..^^^.^^....^.^...^^^...^^^^..^...^...^^.^.^.......^..^^...^.^.^^..^^^^^...^.";

fn main() {
    println!("Part 1: {}", count_safe_tiles(40));
    println!("Part 2: {}", count_safe_tiles(400000));
}

fn map_input(input: &str) -> Vec<bool> {
    input.chars().map(|x| x == '.').collect()
}

fn count_safe_tiles(rows: usize) -> usize {
    let input = map_input(INPUT);
    let mut result = vec![input];
    for i in 0..rows - 1 {
        let current_row = &result[i];
        let mut next_row = vec![];

        for j in 0..current_row.len() {
            next_row.push(get_tile(current_row, j));
        }
        result.push(next_row);
    }

    result.iter().map(|x| safe_tiles_in_row(x)).sum()
}

fn safe_tiles_in_row(row: &Vec<bool>) -> usize {
    row.iter().filter(|x| **x).count()
}

fn get_tile(current_row: &Vec<bool>, index: usize) -> bool {
    let left = if index == 0 {
        true
    } else {
        current_row[index - 1]
    };
    let center = current_row[index];
    let right = if index == current_row.len() - 1 {
        true
    } else {
        current_row[index + 1]
    };
    !is_trap(left, center, right)
}

fn is_trap(left: bool, center: bool, right: bool) -> bool {
    if left && center && !right {
        return true;
    }
    if !left && center && right {
        return true;
    }
    if left && !center && !right {
        return true;
    }
    if !left && !center && right {
        return true;
    }
    false
}
