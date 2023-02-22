use itertools::Itertools;
use std::collections::{HashMap, HashSet, VecDeque};

#[derive(Debug, Hash, Eq, PartialEq, Clone)]
struct Node {
    id: String,
    size: usize,
    used: usize,
    available: usize,
    usage: usize,
}

impl Node {
    fn is_viable_pair(&self, other: &Node) -> bool {
        self.used > 0 && self.used <= other.available
    }

    fn get_x(&self) -> usize {
        let mut tokens = self.id.split('-');
        let x = &tokens.next().unwrap()[1..].parse::<usize>().unwrap();
        *x
    }

    fn get_y(&self) -> usize {
        let mut tokens = self.id.split('-');
        let y = &tokens.next_back().unwrap()[1..].parse::<usize>().unwrap();
        *y
    }

    fn is_goal(&self) -> bool {
        self.id == "x37-y0"
    }

    fn is_empty(&self) -> bool {
        self.used == 0
    }

    fn is_full(&self) -> bool {
        self.usage > 90
    }
}

impl From<&str> for Node {
    fn from(s: &str) -> Self {
        let mut tokens = s.split_whitespace();
        let id = &tokens.next().unwrap()[15..];
        let mut size = tokens.next().unwrap().chars();
        size.next_back();
        let size = size.as_str().parse::<usize>().unwrap();
        let mut used = tokens.next().unwrap().chars();
        used.next_back();
        let used = used.as_str().parse::<usize>().unwrap();
        let mut available = tokens.next().unwrap().chars();
        available.next_back();
        let available = available.as_str().parse::<usize>().unwrap();
        let mut usage = tokens.next().unwrap().chars();
        usage.next_back();
        let usage = usage.as_str().parse::<usize>().unwrap();
        Node {
            id: id.to_string(),
            size,
            used,
            available,
            usage,
        }
    }
}

fn main() {
    let input = include_str!("input.txt");

    println!("Part 1: {}", viable_pair_count(input));
    println!("Part 2: {}", fewest_steps_to_move_data(input));
}

fn viable_pair_count(input: &str) -> usize {
    let nodes: HashSet<Node> = parse_nodes(input);
    let permutations = nodes.iter().permutations(2);
    permutations.filter(|p| p[0].is_viable_pair(p[1])).count()
}

fn fewest_steps_to_move_data(input: &str) -> usize {
    let nodes: HashSet<Node> = parse_nodes(input);
    let max_x = nodes.clone().iter().map(|n| n.get_x()).max().unwrap();
    let max_y = nodes.clone().iter().map(|n| n.get_y()).max().unwrap();
    let temp_node = Node {
        id: "0-0".to_string(),
        size: 0,
        used: 0,
        available: 0,
        usage: 0,
    };
    let mut grid = vec![vec![temp_node; max_x + 1]; max_y + 1];
    for node in nodes.clone() {
        let x = node.get_x();
        let y = node.get_y();
        grid[y][x] = node;
    }

    let empty_node = nodes.iter().find(|n| n.is_empty()).unwrap();
    let goal_node = nodes.iter().find(|n| n.is_goal()).unwrap();

    let steps = astar(&grid, &empty_node, &goal_node);
    (steps - 1) + (max_x - 1) * 5 + 1
}

fn astar(grid: &Vec<Vec<Node>>, start: &Node, goal: &Node) -> usize {
    let mut open_set = VecDeque::new();
    let mut closed_set = grid
        .iter()
        .flatten()
        .filter(|n| n.is_full())
        .cloned()
        .collect::<HashSet<Node>>();
    let mut g_score = HashMap::new();
    let mut f_score = HashMap::new();
    open_set.push_back(start.clone());
    g_score.insert(start.clone(), 0);
    f_score.insert(start.clone(), heuristic_cost_estimate(start, goal));

    while !open_set.is_empty() {
        let current = open_set.pop_front().unwrap();
        if current.is_goal() {
            return *g_score.get(&current).unwrap() as usize;
        }
        closed_set.insert(current.clone());
        for neighbor in get_neighbors(&current, &grid) {
            if closed_set.contains(&neighbor) {
                continue;
            }
            let tentative_g_score = g_score.get(&current).unwrap() + 1;
            if !open_set.contains(&neighbor) {
                open_set.push_back(neighbor.clone());
            } else if tentative_g_score >= *g_score.get(&neighbor).unwrap() {
                continue;
            }
            g_score.insert(neighbor.clone(), tentative_g_score);
            f_score.insert(
                neighbor.clone(),
                tentative_g_score + heuristic_cost_estimate(&neighbor, goal),
            );
        }
    }
    0
}

fn heuristic_cost_estimate(start: &Node, goal: &Node) -> usize {
    let x1 = start.get_x();
    let y1 = start.get_y();
    let x2 = goal.get_x();
    let y2 = goal.get_y();
    x1.abs_diff(x2) + y1.abs_diff(y2)
}

fn get_neighbors(node: &Node, grid: &Vec<Vec<Node>>) -> Vec<Node> {
    let mut neighbors = Vec::new();
    let x = node.get_x();
    let y = node.get_y();
    if x > 0 {
        neighbors.push(grid[y][x - 1].clone());
    }
    if x < grid[0].len() - 1 {
        neighbors.push(grid[y][x + 1].clone());
    }
    if y > 0 {
        neighbors.push(grid[y - 1][x].clone());
    }
    if y < grid.len() - 1 {
        neighbors.push(grid[y + 1][x].clone());
    }
    neighbors
}

fn parse_nodes(input: &str) -> HashSet<Node> {
    input
        .lines()
        .skip(2)
        .map(|line| Node::from(line))
        .collect::<HashSet<Node>>()
}
