enum Instruction {
    Cpy(String, String),
    Inc(String),
    Dec(String),
    Jnz(String, String),
}

fn main() {
    let input = include_str!("input.txt");
    let instructions = parse_instructions(input);

    let mut registers: Vec<usize> = vec![0, 0, 0, 0];
    apply_instructions(&instructions, &mut registers);
    println!("Part 1: {}", registers[0]);

    let mut registers: Vec<usize> = vec![0, 0, 1, 0];
    apply_instructions(&instructions, &mut registers);
    println!("Part 2: {}", registers[0]);
}

fn apply_instructions(instructions: &Vec<Instruction>, registers: &mut Vec<usize>) {
    let mut index: usize = 0;
    while index < instructions.len() {
        let instruction = &instructions[index];
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
            _ => panic!("Unknown instruction"),
        }
    }
    instructions
}

fn cpy(x: &str, y: &str, registers: &mut Vec<usize>) {
    let parsed_x = x.parse::<i32>();
    let y = y.chars().nth(0).unwrap() as i32 - 97;
    match parsed_x {
        Ok(parsed_x) => {
            registers[y as usize] = parsed_x as usize;
        }
        Err(_) => {
            let x = x.chars().nth(0).unwrap() as i32 - 97;
            registers[y as usize] = registers[x as usize];
        }
    }
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
    let parsed_x = x.parse::<i32>();
    let y = y.parse::<i32>();
    match parsed_x {
        Ok(parsed_x) => {
            if parsed_x != 0 {
                *index = (*index as i32 + y.unwrap()) as usize;
            } else {
                *index += 1;
            }
        }
        Err(_) => {
            let x = x.chars().nth(0).unwrap() as i32 - 97;
            if registers[x as usize] != 0 {
                *index = (*index as i32 + y.unwrap()) as usize;
            } else {
                *index += 1;
            }
        }
    }
}
