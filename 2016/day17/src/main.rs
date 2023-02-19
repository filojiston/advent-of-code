use std::collections::HashSet;
use std::collections::VecDeque;

const PASSCODE: &str = "pvhmgsws";
const DIRECTIONS: &str = "UDLR";

#[derive(Clone)]
struct Point {
    x: usize,
    y: usize,
}

#[derive(Clone)]
struct State {
    location: Point,
    path: String,
}

impl State {
    fn new() -> State {
        State {
            location: Point { x: 0, y: 0 },
            path: String::new(),
        }
    }

    fn move_to(&mut self, direction: char) {
        match direction {
            'U' => self.location.y -= 1,
            'D' => self.location.y += 1,
            'L' => self.location.x -= 1,
            'R' => self.location.x += 1,
            _ => panic!("Invalid direction"),
        }
        self.path.push(direction);
    }

    fn is_final(&self) -> bool {
        self.location.x == 3 && self.location.y == 3
    }

    fn is_valid_move(&self, direction: char, hex_digit: char) -> bool {
        hex_digit.is_ascii_hexdigit()
            && hex_digit.to_digit(16).unwrap() > 10
            && match direction {
                'U' => self.location.y > 0,
                'D' => self.location.y < 3,
                'L' => self.location.x > 0,
                'R' => self.location.x < 3,
                _ => panic!("Invalid direction"),
            }
    }
}

fn main() {
    let all_paths = all_paths();
    let shortest_path = all_paths.iter().min_by_key(|p| p.len()).unwrap();
    let longest_path = all_paths.iter().max_by_key(|p| p.len()).unwrap();

    println!("Part 1: {}", shortest_path);
    println!("Part 2: {}", longest_path.len());
}

fn all_paths() -> HashSet<String> {
    let mut states = VecDeque::new();
    let mut paths = HashSet::new();
    states.push_back(State::new());

    while let Some(state) = states.pop_front() {
        if state.is_final() {
            paths.insert(state.path);
            continue;
        }

        for next_state in get_next_states(&state) {
            states.push_back(next_state);
        }
    }

    paths
}

fn get_next_states(state: &State) -> Vec<State> {
    let mut states = Vec::new();
    let hash = md5::compute(format!("{}{}", PASSCODE, state.path));
    let hash = &format!("{:x}", hash)[0..4];
    for (i, c) in hash.chars().enumerate() {
        let direction = DIRECTIONS.chars().nth(i).unwrap();
        if state.is_valid_move(direction, c) {
            let mut new_state = state.clone();
            new_state.move_to(direction);
            states.push(new_state);
        }
    }
    states
}
