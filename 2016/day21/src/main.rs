#[derive(Debug)]
enum Command {
    SwapPosition(usize, usize),
    SwapLetter(char, char),
    RotateLeft(usize),
    RotateRight(usize),
    RotateBased(char),
    Reverse(usize, usize),
    Move(usize, usize),
}

fn main() {
    let input = include_str!("input.txt");
    let commands = parse_commands(input);
    println!("Part 1: {}", scramble("abcdefgh", &commands));
    println!("Part 2: {}", unscramble("fbgdceah", &commands));
}

fn parse_commands(input: &str) -> Vec<Command> {
    input
        .lines()
        .map(|line| {
            let tokens = line.split_whitespace().collect::<Vec<_>>();
            let command = tokens[0];
            match command {
                "swap" => match tokens[1] {
                    "position" => Command::SwapPosition(
                        tokens[2].parse::<usize>().unwrap(),
                        tokens[5].parse::<usize>().unwrap(),
                    ),
                    "letter" => Command::SwapLetter(
                        tokens[2].chars().next().unwrap(),
                        tokens[5].chars().next().unwrap(),
                    ),
                    _ => panic!("Unknown command: {}", command),
                },
                "rotate" => match tokens[1] {
                    "left" => Command::RotateLeft(tokens[2].parse::<usize>().unwrap()),
                    "right" => Command::RotateRight(tokens[2].parse::<usize>().unwrap()),
                    "based" => Command::RotateBased(tokens[6].chars().next().unwrap()),
                    _ => panic!("Unknown command: {}", command),
                },
                "reverse" => Command::Reverse(
                    tokens[2].parse::<usize>().unwrap(),
                    tokens[4].parse::<usize>().unwrap(),
                ),
                "move" => Command::Move(
                    tokens[2].parse::<usize>().unwrap(),
                    tokens[5].parse::<usize>().unwrap(),
                ),
                _ => panic!("Unknown command: {}", command),
            }
        })
        .collect()
}

fn scramble(password: &str, commands: &Vec<Command>) -> String {
    let mut puzzle_input = password.chars().collect::<Vec<_>>();
    for command in commands {
        match *command {
            Command::SwapPosition(x, y) => {
                puzzle_input.swap(x, y);
            }
            Command::SwapLetter(x, y) => {
                let x = puzzle_input.iter().position(|&c| c == x).unwrap();
                let y = puzzle_input.iter().position(|&c| c == y).unwrap();
                puzzle_input.swap(x, y);
            }
            Command::RotateLeft(x) => {
                puzzle_input.rotate_left(x);
            }
            Command::RotateRight(x) => {
                puzzle_input.rotate_right(x);
            }
            Command::RotateBased(x) => {
                let x = puzzle_input.iter().position(|&c| c == x).unwrap();
                let count = x + 1 + if x >= 4 { 1 } else { 0 };
                let n = puzzle_input.len();
                for _ in 0..count / n {
                    puzzle_input.rotate_right(n);
                }
                puzzle_input.rotate_right(count % n);
            }
            Command::Reverse(x, y) => {
                puzzle_input[x..=y].reverse();
            }
            Command::Move(x, y) => {
                let c = puzzle_input.remove(x);
                puzzle_input.insert(y, c);
            }
        }
    }
    puzzle_input.iter().collect::<String>()
}

fn unscramble(password: &str, commands: &Vec<Command>) -> String {
    let mut puzzle_input = password.chars().collect::<Vec<_>>();
    for command in commands.into_iter().rev() {
        match *command {
            Command::SwapPosition(x, y) => {
                puzzle_input.swap(x, y);
            }
            Command::SwapLetter(x, y) => {
                let x = puzzle_input.iter().position(|&c| c == x).unwrap();
                let y = puzzle_input.iter().position(|&c| c == y).unwrap();
                puzzle_input.swap(x, y);
            }
            Command::RotateLeft(x) => {
                puzzle_input.rotate_right(x);
            }
            Command::RotateRight(x) => {
                puzzle_input.rotate_left(x);
            }
            // i admit that i've cheated here (thx reddit!)
            // i saw people used brute force with permutations to solve this
            // but this works too
            Command::RotateBased(x) => {
                let x = puzzle_input.iter().position(|&c| c == x).unwrap();
                match x {
                    0 | 1 => puzzle_input.rotate_left(1),
                    2 => puzzle_input.rotate_left(6),
                    3 => puzzle_input.rotate_left(2),
                    4 => puzzle_input.rotate_left(7),
                    5 => puzzle_input.rotate_left(3),
                    6 => continue,
                    7 => puzzle_input.rotate_left(4),
                    _ => panic!("Unknown position: {}", x),
                }
            }
            Command::Reverse(x, y) => {
                puzzle_input[x..=y].reverse();
            }
            Command::Move(x, y) => {
                let c = puzzle_input.remove(y);
                puzzle_input.insert(x, c);
            }
        }
    }
    puzzle_input.iter().collect::<String>()
}
