use itertools::Itertools;
use std::collections::BTreeSet;
use std::thread;

#[derive(PartialEq, Eq, Hash, Clone, Debug, PartialOrd, Ord)]
enum Device {
    Generator(String),
    Microchip(String),
}

#[derive(PartialEq, Eq, Hash, Clone, Debug, PartialOrd, Ord)]
struct State {
    elevator: usize,
    floors: Vec<BTreeSet<Device>>,
}

// thanks dude: https://github.com/BartMassey/advent-of-code-2016/blob/master/day11/soln.rs
impl State {
    fn initial_part1() -> Self {
        State {
            elevator: 0,
            floors: vec![
                vec![
                    Device::Microchip("promethium".to_string()),
                    Device::Generator("promethium".to_string()),
                ]
                .into_iter()
                .collect(),
                vec![
                    Device::Generator("cobalt".to_string()),
                    Device::Generator("curium".to_string()),
                    Device::Generator("ruthenium".to_string()),
                    Device::Generator("plutonium".to_string()),
                ]
                .into_iter()
                .collect(),
                vec![
                    Device::Microchip("cobalt".to_string()),
                    Device::Microchip("curium".to_string()),
                    Device::Microchip("ruthenium".to_string()),
                    Device::Microchip("plutonium".to_string()),
                ]
                .into_iter()
                .collect(),
                BTreeSet::new(),
            ],
        }
    }

    fn initial_part2() -> Self {
        let state = State::initial_part1();
        let mut floors = state.floors.clone();
        floors[0].insert(Device::Generator("elerium".to_string()));
        floors[0].insert(Device::Microchip("elerium".to_string()));
        floors[0].insert(Device::Generator("dilithium".to_string()));
        floors[0].insert(Device::Microchip("dilithium".to_string()));
        State {
            elevator: state.elevator,
            floors,
        }
    }

    fn try_move(&self, direction: isize) -> Option<(usize, usize)> {
        let source = self.elevator as isize;
        let destination = source + direction;
        if destination < 0 || destination >= self.floors.len() as isize {
            return None;
        }
        Some((source as usize, destination as usize))
    }

    fn try_traverse(&self, direction: isize, devices_to_move: &BTreeSet<Device>) -> Option<State> {
        if devices_to_move.is_empty() {
            return None;
        }

        let (source, destination) = match self.try_move(direction) {
            Some((source, destination)) => (source, destination),
            None => return None,
        };

        let new_source = self.floors[source]
            .difference(devices_to_move)
            .cloned()
            .collect();
        if !is_safe(&new_source) {
            return None;
        }

        let new_destination = self.floors[destination]
            .union(devices_to_move)
            .cloned()
            .collect();
        if !is_safe(&new_destination) {
            return None;
        }

        let mut new_floors = self.floors.clone();
        new_floors[source] = new_source;
        new_floors[destination] = new_destination;
        Some(State {
            elevator: destination,
            floors: new_floors,
        })
    }

    fn traversals(&self) -> BTreeSet<State> {
        let mut traversals = BTreeSet::new();
        for direction in [-1, 1].iter() {
            for devices_to_move in self.floors[self.elevator].iter().combinations(1) {
                if let Some(traversal) =
                    self.try_traverse(*direction, &devices_to_move.into_iter().cloned().collect())
                {
                    traversals.insert(traversal);
                }
            }
            for devices_to_move in self.floors[self.elevator].iter().combinations(2) {
                if let Some(traversal) =
                    self.try_traverse(*direction, &devices_to_move.into_iter().cloned().collect())
                {
                    traversals.insert(traversal);
                }
            }
        }
        traversals
    }
}

fn is_safe(devices: &BTreeSet<Device>) -> bool {
    let mut generators = BTreeSet::new();
    let mut microchips = BTreeSet::new();
    for device in devices {
        match device {
            Device::Generator(name) => {
                generators.insert(name);
            }
            Device::Microchip(name) => {
                microchips.insert(name);
            }
        }
    }

    generators.is_empty() || microchips.is_subset(&generators)
}

fn main() {
    let part1 = thread::spawn(|| {
        let start_state_part1 = State::initial_part1();
        bfs_states(start_state_part1, 10)
    });
    let part2 = thread::spawn(|| {
        let start_state_part2 = State::initial_part2();
        bfs_states(start_state_part2, 14)
    });

    println!("Part 1: {}", part1.join().unwrap());
    println!("Part 2: {}", part2.join().unwrap());
}

fn bfs_states(start_state: State, item_count: usize) -> usize {
    let mut states = vec![start_state];
    let mut visited = BTreeSet::new();
    let mut steps = 0;

    loop {
        let mut new_states = BTreeSet::new();
        for state in states {
            if state.floors[3].len() == item_count {
                return steps;
            }
            for traversal in state.traversals() {
                if !visited.contains(&traversal) {
                    new_states.insert(traversal);
                }
            }
            visited.insert(state);
        }
        states = new_states.into_iter().collect();
        steps += 1;
    }
}
