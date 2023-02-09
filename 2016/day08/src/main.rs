use matrix::prelude::*;

const WIDTH: usize = 50;
const HEIGHT: usize = 6;
const INVALID_INPUT: usize = 61;

struct Instruction {
    command: String,
    row: usize,
    col: usize,
    shift_by: usize,
}

impl From<&str> for Instruction {
    fn from(s: &str) -> Instruction {
        let words = s.split(' ').collect::<Vec<&str>>();
        if words[0] == "rect" {
            let mut dimensions = words[1].split('x');
            Instruction {
                command: "rect".to_string(),
                row: dimensions.next().unwrap().parse::<usize>().unwrap(),
                col: dimensions.next().unwrap().parse::<usize>().unwrap(),
                shift_by: 0,
            }
        } else {
            let is_row = if words[1] == "row" { true } else { false };
            let num = words[2]
                .split('=')
                .last()
                .unwrap()
                .parse::<usize>()
                .unwrap();
            Instruction {
                command: "rotate".to_string(),
                row: if is_row { num } else { INVALID_INPUT },
                col: if is_row { INVALID_INPUT } else { num },
                shift_by: words[4].parse::<usize>().unwrap(),
            }
        }
    }
}

fn main() {
    let input = include_str!("input.txt");
    let mut screen: matrix::format::Compressed<bool> = Compressed::zero((HEIGHT, WIDTH));

    for line in input.lines() {
        let instruction = Instruction::from(line);
        match instruction.command.as_str() {
            "rect" => {
                for i in 0..instruction.col {
                    for j in 0..instruction.row {
                        screen.set((i, j), true);
                    }
                }
            }
            "rotate" => {
                if instruction.row == INVALID_INPUT {
                    rotate_col(&mut screen, instruction);
                } else {
                    rotate_row(&mut screen, instruction);
                }
            }
            _ => panic!("unknown command!"),
        }
    }

    println!("Part1: {}", screen.nonzeros());
    println!("Part2:");
    draw_screen(&screen);
}

fn rotate_col(screen: &mut matrix::format::Compressed<bool>, instruction: Instruction) {
    let mut new_col = (0..HEIGHT)
        .map(|row| screen.get((row, instruction.col)))
        .collect::<Vec<bool>>();
    new_col.rotate_right(instruction.shift_by);
    (0..HEIGHT).for_each(|row| screen.set((row, instruction.col), new_col[row]));
}

fn rotate_row(screen: &mut matrix::format::Compressed<bool>, instruction: Instruction) {
    let mut new_row = (0..WIDTH)
        .map(|col| screen.get((instruction.row, col)))
        .collect::<Vec<bool>>();
    new_row.rotate_right(instruction.shift_by);
    (0..WIDTH).for_each(|col| screen.set((instruction.row, col), new_row[col]));
}

fn draw_screen(screen: &matrix::format::Compressed<bool>) {
    for i in 0..screen.rows() {
        for j in 0..screen.columns() {
            if screen.get((i, j)) == true {
                print!("#")
            } else {
                print!(".")
            }
        }
        println!()
    }
}
