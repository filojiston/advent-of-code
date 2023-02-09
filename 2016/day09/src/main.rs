fn main() {
    let input = include_str!("input.txt").replace(" ", "");
    let decompressed_len = part1(input.clone()).len();
    let all_decompressed_len = part2(input);

    println!("Part 1: {}", decompressed_len);
    println!("Part 2: {}", all_decompressed_len);
}

fn part1(input: String) -> String {
    let chars = input.chars().collect::<Vec<char>>();

    let mut decompressed = String::new();
    let mut i = 0;
    while i < chars.len() {
        if chars[i] == '(' {
            let (decompressed_data, idx) = decompress(&chars, i);
            decompressed.push_str(decompressed_data.as_str());
            i = idx;
        } else {
            decompressed.push(chars[i]);
        }
        i += 1;
    }

    decompressed
}

// algo: https://github.com/rhardih/aoc/blob/master/2016/9p2.c
fn part2(input: String) -> usize {
    let chars = input.chars().collect::<Vec<char>>();

    let mut weights = vec![1; input.len()];
    let mut length = 0;

    let mut i = 0;
    while i < chars.len() {
        if chars[i] == '(' {
            let (count, repeat, idx) = parse_marker(&chars, i);
            i = idx;
            for j in i..i + count {
                weights[j] *= repeat;
            }
        } else {
            length += weights[i];
            i += 1;
        }
    }

    length
}

fn decompress(chars: &Vec<char>, idx: usize) -> (String, usize) {
    let (count, repeat, idx) = parse_marker(chars, idx);
    let decompressed_data = chars[idx..idx + count]
        .repeat(repeat)
        .iter()
        .collect::<String>();
    (decompressed_data, idx + count - 1)
}

fn parse_marker(chars: &Vec<char>, idx: usize) -> (usize, usize, usize) {
    let mut marker = String::new();
    let mut curr = chars[idx];
    let mut idx = idx;

    while curr != ')' {
        idx += 1;
        curr = chars[idx];
        if curr != ')' {
            marker.push(curr);
        }
    }

    let mut marker = marker.split('x');
    let count = marker.next().unwrap().parse::<usize>().unwrap();
    let repeat = marker.next().unwrap().parse::<usize>().unwrap();
    (count, repeat, idx + 1)
}
