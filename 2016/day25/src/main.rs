#[derive(Debug, Clone)]
enum Instruction {
    Cpy(String, String),
    Inc(String),
    Dec(String),
    Jnz(String, String),
    Out(String),
}

// third time of same question. no optimization, runs about 7 seconds
fn main() {
    let input = include_str!("input.txt");

    let mut instructions = parse_instructions(input);
    let mut a_value = 0;
    loop {
        let mut registers: Vec<usize> = vec![a_value, 0, 0, 0];
        let result = apply_instructions(&mut instructions, &mut registers);
        if result == vec![0, 1, 0, 1, 0, 1, 0, 1, 0, 1] {
            println!("Part1: {}", a_value);
            break;
        } else {
            a_value += 1;
        }
    }
}

fn parse_instructions(input: &str) -> Vec<Instruction> {
    let mut instructions: Vec<Instruction> = Vec::new();
    for line in input.lines() {
        let parts = line.split_whitespace().collect::<Vec<&str>>();
        match parts[0] {
            "cpy" => instructions.push(Instruction::Cpy(
                String::from(parts[1]),
                String::from(parts[2]),
            )),
            "inc" => instructions.push(Instruction::Inc(String::from(parts[1]))),
            "dec" => instructions.push(Instruction::Dec(String::from(parts[1]))),
            "jnz" => instructions.push(Instruction::Jnz(
                String::from(parts[1]),
                String::from(parts[2]),
            )),
            "out" => instructions.push(Instruction::Out(String::from(parts[1]))),
            _ => panic!("Unknown instruction"),
        }
    }
    instructions
}

fn apply_instructions(instructions: &mut Vec<Instruction>, registers: &mut Vec<usize>) -> Vec<i32> {
    let mut index: usize = 0;
    let mut outs = Vec::new();
    while index < instructions.len() {
        let insts = instructions.clone();
        let instruction = insts.get(index).unwrap();
        match instruction {
            &Instruction::Cpy(ref x, ref y) => {
                cpy(x, y, registers);
                index += 1;
            }
            &Instruction::Inc(ref x) => {
                inc(x, registers);
                index += 1;
            }
            &Instruction::Dec(ref x) => {
                dec(x, registers);
                index += 1;
            }
            &Instruction::Jnz(ref x, ref y) => {
                jnz(x, y, &mut index, registers);
            }
            &Instruction::Out(ref x) => {
                outs.push(out(x, registers));
                index += 1;
            }
        }

        if outs.len() >= 10 {
            return outs;
        }
    }

    outs
}

fn cpy(x: &str, y: &str, registers: &mut Vec<usize>) {
    let x = x.parse::<i32>().unwrap_or_else(|_| {
        let x = x.chars().nth(0).unwrap() as i32 - 97;
        registers[x as usize] as i32
    });
    let y = y.chars().nth(0).unwrap() as i32 - 97;

    registers[y as usize] = x as usize;
}

fn inc(x: &str, registers: &mut Vec<usize>) {
    let x = x.chars().nth(0).unwrap() as i32 - 97;
    registers[x as usize] += 1;
}

fn dec(x: &str, registers: &mut Vec<usize>) {
    let x = x.chars().nth(0).unwrap() as i32 - 97;
    registers[x as usize] -= 1;
}

fn jnz(x: &str, y: &str, index: &mut usize, registers: &mut Vec<usize>) {
    let x = x.parse::<i32>().unwrap_or_else(|_| {
        let x = x.chars().nth(0).unwrap() as i32 - 97;
        registers[x as usize] as i32
    });
    let y = y.parse::<i32>().unwrap_or_else(|_| {
        let y = y.chars().nth(0).unwrap() as i32 - 97;
        registers[y as usize] as i32
    });

    if x != 0 {
        *index = (*index as i32 + y) as usize;
    } else {
        *index += 1;
    }
}

fn out(x: &str, registers: &mut Vec<usize>) -> i32 {
    let x = x.parse::<i32>().unwrap_or_else(|_| {
        let x = x.chars().nth(0).unwrap() as i32 - 97;
        registers[x as usize] as i32
    });
    x
}
