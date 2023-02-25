use itertools::Itertools;
use std::collections::{HashMap, HashSet, VecDeque};

fn main() {
    let input = include_str!("input.txt");
    let maze = input
        .lines()
        .map(|line| line.chars().collect())
        .collect::<Vec<Vec<char>>>();

    let distances = maze
        .iter()
        .enumerate()
        .flat_map(|(y, row)| {
            row.iter().enumerate().filter_map(move |(x, &c)| match c {
                '0'..='9' => Some((c, (x, y))),
                _ => None,
            })
        })
        .map(|(c, start)| bfs(&maze, start, c))
        .collect::<HashMap<char, HashMap<char, usize>>>();

    println!("{:?}", part1(&distances));
    println!("{:?}", part2(&distances));
}

fn bfs(maze: &Vec<Vec<char>>, start: (usize, usize), c: char) -> (char, HashMap<char, usize>) {
    let mut queue = VecDeque::new();
    queue.push_back((start, 0));
    let mut visited = HashSet::new();
    visited.insert(start);
    let mut distances = HashMap::new();
    while let Some((pos, distance)) = queue.pop_front() {
        for &dir in &[(0, 1), (0, -1), (1, 0), (-1, 0)] {
            let next = (pos.0 as i32 + dir.0, pos.1 as i32 + dir.1);
            if next.0 < 0
                || next.1 < 0
                || next.0 >= maze[0].len() as i32
                || next.1 >= maze.len() as i32
            {
                continue;
            }
            let next = (next.0 as usize, next.1 as usize);
            if visited.contains(&next) {
                continue;
            }
            visited.insert(next);
            match maze[next.1][next.0] {
                '#' => continue,
                '0'..='9' => {
                    distances.insert(maze[next.1][next.0], distance + 1);
                }
                _ => {}
            }
            queue.push_back((next, distance + 1));
        }
    }
    (c, distances)
}

fn part1(distances: &HashMap<char, HashMap<char, usize>>) -> usize {
    distances
        .keys()
        .permutations(distances.len())
        .filter(|path| path[0] == &'0')
        .map(|path| {
            path.iter()
                .tuple_windows()
                .map(|(a, b)| distances[a][b])
                .sum::<usize>()
        })
        .min()
        .unwrap()
}

fn part2(distances: &HashMap<char, HashMap<char, usize>>) -> usize {
    distances
        .keys()
        .permutations(distances.len())
        .filter(|path| path[0] == &'0')
        .map(|path| {
            path.iter()
                .tuple_windows()
                .map(|(a, b)| distances[a][b])
                .sum::<usize>()
                + distances[path[path.len() - 1]][&'0']
        })
        .min()
        .unwrap()
}
