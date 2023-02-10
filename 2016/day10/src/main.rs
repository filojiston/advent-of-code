mod container;
mod instruction;

use container::Container;
use instruction::Instruction;
use std::{
    cell::RefCell,
    collections::{HashMap, VecDeque},
};

fn main() {
    let input = include_str!("input.txt");
    let (containers, mut instructions) = parse_instructions(input);

    let mut target_container_id: Option<String> = Option::None;
    while !instructions.is_empty() {
        let instruction = instructions.pop_front().unwrap();
        if instruction.value.is_some() {
            containers
                .get(&instruction.source)
                .unwrap()
                .borrow_mut()
                .give(instruction.value.unwrap());
        } else {
            let mut source = containers.get(&instruction.source).unwrap().borrow_mut();
            let mut target_low = containers
                .get(&instruction.target_low)
                .unwrap()
                .borrow_mut();
            let mut target_high = containers
                .get(&instruction.target_high)
                .unwrap()
                .borrow_mut();

            if source.is_ready() {
                target_low.give(source.low());
                target_high.give(source.high());
                source.reset();
            } else {
                instructions.push_back(instruction);
            }
        }

        if target_container_id.is_none() {
            for (id, container) in containers.iter() {
                let container = container.borrow();
                if container.is_ready() && container.low() == 17 && container.high() == 61 {
                    target_container_id = Some(id.clone());
                    break;
                }
            }
        }
    }

    let product = containers
        .iter()
        .filter(|(_, container)| {
            (*container).borrow().id == "output 0"
                || (*container).borrow().id == "output 1"
                || (*container).borrow().id == "output 2"
        })
        .map(|(_, container)| container.borrow().low())
        .product::<u32>();

    println!("Part 1: {}", target_container_id.unwrap());
    println!("Part 2: {}", product);
}

fn parse_instructions(input: &str) -> (HashMap<String, RefCell<Container>>, VecDeque<Instruction>) {
    let mut containers: HashMap<String, RefCell<Container>> = HashMap::new();
    let mut instructions = VecDeque::new();

    for line in input.lines() {
        let instruction = Instruction::from(line);
        containers
            .entry(instruction.source.clone())
            .or_insert(RefCell::new(Container::new(instruction.source.clone())));

        if instruction.value.is_none() {
            containers
                .entry(instruction.target_low.clone())
                .or_insert(RefCell::new(Container::new(instruction.target_low.clone())));
            containers
                .entry(instruction.target_high.clone())
                .or_insert(RefCell::new(Container::new(
                    instruction.target_high.clone(),
                )));
        }
        instructions.push_back(instruction);
    }

    (containers, instructions)
}
