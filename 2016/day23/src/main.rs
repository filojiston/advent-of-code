// assembly optimizations... thx reddit
#[derive(Debug, Clone)]
enum Instruction {
    Cpy(String, String),
    Inc(String),
    Dec(String),
    Jnz(String, String),
    Tgl(String),
    Mul(String, String),
    Add(String, String),
    Nop,
}

impl Instruction {
    fn is_valid(&self) -> bool {
        match self {
            &Instruction::Cpy(ref x, ref y) => {
                let x = x.parse::<i32>();
                let y = y.parse::<i32>();
                match (x, y) {
                    (Ok(_), Ok(_)) => false,
                    (Ok(_), Err(_)) => true,
                    (Err(_), Ok(_)) => false,
                    (Err(_), Err(_)) => true,
                }
            }
            &Instruction::Inc(ref x) => {
                let x = x.parse::<i32>();
                match x {
                    Ok(_) => false,
                    Err(_) => true,
                }
            }
            &Instruction::Dec(ref x) => {
                let x = x.parse::<i32>();
                match x {
                    Ok(_) => false,
                    Err(_) => true,
                }
            }
            &Instruction::Jnz(ref _x, ref _y) => true,
            &Instruction::Tgl(ref _x) => true,
            &Instruction::Mul(ref _x, ref _y) => true,
            &Instruction::Add(ref _x, ref _y) => true,
            &Instruction::Nop => true,
        }
    }
}

fn main() {
    let input = include_str!("input.txt");

    let mut instructions = parse_instructions(input);
    let mut registers: Vec<usize> = vec![7, 0, 0, 0];
    apply_instructions(&mut instructions, &mut registers);
    println!("Part 1: {}", registers[0]);

    let mut instructions = parse_instructions(input);
    let mut registers: Vec<usize> = vec![12, 0, 0, 0];
    apply_instructions(&mut instructions, &mut registers);
    println!("Part 2: {}", registers[0]);
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
            "tgl" => instructions.push(Instruction::Tgl(String::from(parts[1]))),
            "mul" => instructions.push(Instruction::Mul(
                String::from(parts[1]),
                String::from(parts[2]),
            )),
            "add" => instructions.push(Instruction::Add(
                String::from(parts[1]),
                String::from(parts[2]),
            )),
            "nop" => instructions.push(Instruction::Nop),
            _ => panic!("Unknown instruction"),
        }
    }
    instructions
}

fn apply_instructions(instructions: &mut Vec<Instruction>, registers: &mut Vec<usize>) {
    let mut index: usize = 0;
    while index < instructions.len() {
        let insts = instructions.clone();
        let instruction = insts.get(index).unwrap();
        if !instruction.is_valid() {
            index += 1;
            continue;
        }
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
            &Instruction::Tgl(ref x) => {
                tgl(x, &mut index, registers, instructions);
                index += 1;
            }
            &Instruction::Mul(ref x, ref y) => {
                mul(x, y, registers);
                index += 1;
            }
            &Instruction::Add(ref x, ref y) => {
                add(x, y, registers);
                index += 1;
            }
            &Instruction::Nop => {
                index += 1;
            }
        }
    }
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

fn tgl(
    x: &str,
    index: &mut usize,
    registers: &mut Vec<usize>,
    instructions: &mut Vec<Instruction>,
) {
    let x = x.parse::<i32>().unwrap_or_else(|_| {
        let x = x.chars().nth(0).unwrap() as i32 - 97;
        registers[x as usize] as i32
    });
    let index_to_toggle = (*index as i32 + x) as usize;
    if index_to_toggle < instructions.len() {
        match instructions[index_to_toggle] {
            Instruction::Cpy(ref x, ref y) => {
                instructions[index_to_toggle] = Instruction::Jnz(x.to_string(), y.to_string());
            }
            Instruction::Inc(ref x) => {
                instructions[index_to_toggle] = Instruction::Dec(x.to_string());
            }
            Instruction::Dec(ref x) => {
                instructions[index_to_toggle] = Instruction::Inc(x.to_string());
            }
            Instruction::Jnz(ref x, ref y) => {
                instructions[index_to_toggle] = Instruction::Cpy(x.to_string(), y.to_string());
            }
            Instruction::Tgl(ref x) => {
                instructions[index_to_toggle] = Instruction::Inc(x.to_string());
            }
            Instruction::Mul(ref x, ref y) => {
                instructions[index_to_toggle] = Instruction::Jnz(x.to_string(), y.to_string());
            }
            Instruction::Add(ref x, ref y) => {
                instructions[index_to_toggle] = Instruction::Jnz(x.to_string(), y.to_string());
            }
            Instruction::Nop => {
                instructions[index_to_toggle] = Instruction::Inc(x.to_string());
            }
        }
    }
}

fn mul(x: &str, y: &str, registers: &mut Vec<usize>) {
    let x = x.chars().nth(0).unwrap() as i32 - 97;
    let y = y.chars().nth(0).unwrap() as i32 - 97;
    registers[y as usize] *= registers[x as usize];
}

fn add(x: &str, y: &str, registers: &mut Vec<usize>) {
    let x = x.chars().nth(0).unwrap() as i32 - 97;
    let y = y.chars().nth(0).unwrap() as i32 - 97;
    registers[y as usize] += registers[x as usize];
}
