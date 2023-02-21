struct Range {
    start: usize,
    end: usize,
}

fn main() {
    let input = include_str!("input.txt");
    let ranges = parse_ranges(input);

    let mut addresses = vec![true; 4294967296];
    for range in ranges {
        for i in range.start..=range.end {
            addresses[i] = false;
        }
    }

    println!("Part 1: {}", addresses.iter().position(|&x| x).unwrap());
    println!("Part 2: {}", addresses.iter().filter(|&&x| x).count());
}

fn parse_ranges(input: &str) -> Vec<Range> {
    let mut ranges = Vec::new();
    for line in input.lines() {
        let mut parts = line.split('-');
        let start = parts.next().unwrap().parse().unwrap();
        let end = parts.next().unwrap().parse().unwrap();
        ranges.push(Range { start, end });
    }
    ranges
}
