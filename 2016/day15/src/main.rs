#[derive(Debug)]
struct Disc {
    start_position: usize,
    positions: usize,
}

fn main() {
    let mut discs = read_discs();

    let mut time = 0;
    while !will_capsule_pass_through(&discs, time) {
        time += 1;
    }
    println!("Part 1: {}", time);

    discs.push(Disc {
        start_position: 0,
        positions: 11,
    });

    let mut time = 0;
    while !will_capsule_pass_through(&discs, time) {
        time += 1;
    }
    println!("Part 2: {}", time);
}

fn read_discs() -> Vec<Disc> {
    let input = include_str!("input.txt");

    let mut discs: Vec<Disc> = Vec::new();
    for line in input.lines() {
        let mut parts = line.split_whitespace();
        let positions = parts.nth(3).unwrap().parse::<usize>().unwrap();
        let start_position = parts
            .nth(7)
            .unwrap()
            .replace(".", "")
            .parse::<usize>()
            .unwrap();
        discs.push(Disc {
            start_position,
            positions,
        });
    }

    discs
}

fn will_capsule_pass_through(discs: &Vec<Disc>, time: usize) -> bool {
    for (i, disc) in discs.iter().enumerate() {
        if (disc.start_position + time + i + 1) % disc.positions != 0 {
            return false;
        }
    }

    true
}
