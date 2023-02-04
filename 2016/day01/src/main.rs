#[derive(Clone, PartialEq, Debug)]
struct Point {
    x: i32,
    y: i32,
}

enum Direction {
    North,
    East,
    South,
    West,
}

fn main() {
    let input = include_str!("input.txt");
    let mut loc = Point { x: 0, y: 0 };
    let mut current_direction = Direction::North;
    let mut seen_locations = vec![];
    let mut first_repeated_location: Option<Point> = None;

    for instruction in input.split(", ") {
        let (turn, steps) = instruction.split_at(1);
        let steps = steps.parse::<u8>().unwrap();

        current_direction = turn_to_direction(current_direction, turn);

        for _ in 0..steps {
            match current_direction {
                Direction::North => loc.y += 1,
                Direction::East => loc.x += 1,
                Direction::South => loc.y -= 1,
                Direction::West => loc.x -= 1,
            }

            if first_repeated_location.is_none() && seen_locations.contains(&loc) {
                first_repeated_location = Some(loc.clone());
            }

            seen_locations.push(loc.clone());
        }
    }

    let distance = loc.x.abs() + loc.y.abs();
    let first_repeated_distance = first_repeated_location
        .map(|p| p.x.abs() + p.y.abs())
        .unwrap_or_else(|| panic!("No repeated location found!"));

    println!("Part 1: {}", distance);
    println!("Part 2: {}", first_repeated_distance);
}

fn turn_to_direction(dir: Direction, turn: &str) -> Direction {
    match turn {
        "R" => match dir {
            Direction::North => Direction::East,
            Direction::East => Direction::South,
            Direction::South => Direction::West,
            Direction::West => Direction::North,
        },
        "L" => match dir {
            Direction::North => Direction::West,
            Direction::West => Direction::South,
            Direction::South => Direction::East,
            Direction::East => Direction::North,
        },
        _ => panic!("Invalid turn"),
    }
}
