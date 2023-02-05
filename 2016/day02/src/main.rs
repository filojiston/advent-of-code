#[derive(Clone, PartialEq, Debug)]
struct Position {
    x: usize,
    y: usize,
}

struct Solution {
    keypad: Vec<Vec<char>>,
    position: Position,
    code: String,
}

impl Solution {
    fn new(keypad: Vec<Vec<char>>, pos: Position) -> Solution {
        Solution {
            keypad,
            position: pos,
            code: String::new(),
        }
    }

    fn move_up(&mut self) {
        if self.position.y > 0 && self.keypad[self.position.y - 1][self.position.x] != '-' {
            self.position.y -= 1;
        }
    }

    fn move_down(&mut self) {
        if self.position.y < self.keypad.len() - 1
            && self.keypad[self.position.y + 1][self.position.x] != '-'
        {
            self.position.y += 1;
        }
    }

    fn move_left(&mut self) {
        if self.position.x > 0 && self.keypad[self.position.y][self.position.x - 1] != '-' {
            self.position.x -= 1;
        }
    }

    fn move_right(&mut self) {
        if self.position.x < self.keypad[0].len() - 1
            && self.keypad[self.position.y][self.position.x + 1] != '-'
        {
            self.position.x += 1;
        }
    }

    fn add_code(&mut self) {
        self.code
            .push(self.keypad[self.position.y][self.position.x]);
    }
}

fn main() {
    let input = include_str!("input.txt");
    let mut solution_part1 = Solution::new(
        vec![
            vec!['1', '2', '3'],
            vec!['4', '5', '6'],
            vec!['7', '8', '9'],
        ],
        Position { x: 1, y: 1 },
    );
    let mut solution_part2 = Solution::new(
        vec![
            vec!['-', '-', '1', '-', '-'],
            vec!['-', '2', '3', '4', '-'],
            vec!['5', '6', '7', '8', '9'],
            vec!['-', 'A', 'B', 'C', '-'],
            vec!['-', '-', 'D', '-', '-'],
        ],
        Position { x: 2, y: 0 },
    );

    for line in input.lines() {
        for c in line.chars() {
            match c {
                'U' => {
                    solution_part1.move_up();
                    solution_part2.move_up();
                }
                'D' => {
                    solution_part1.move_down();
                    solution_part2.move_down();
                }
                'L' => {
                    solution_part1.move_left();
                    solution_part2.move_left();
                }
                'R' => {
                    solution_part1.move_right();
                    solution_part2.move_right();
                }
                _ => panic!("Invalid input"),
            }
        }
        solution_part1.add_code();
        solution_part2.add_code();
    }

    println!("Part 1: {}", solution_part1.code);
    println!("Part 2: {}", solution_part2.code);
}
