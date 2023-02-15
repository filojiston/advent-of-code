use std::{
    cmp::Reverse,
    collections::{BinaryHeap, HashSet, VecDeque},
};

#[derive(Debug, Copy, Clone, PartialEq, Eq, PartialOrd, Ord, Hash)]
struct Point {
    x: usize,
    y: usize,
}

const MAZE_ROWS: usize = 100;
const MAZE_COLS: usize = 100;

const FAVOURITE_NUMBER: usize = 1352;

fn main() {
    let mut maze = [[0u8; MAZE_COLS]; MAZE_ROWS];
    set_walls(&mut maze);

    let start = Point { x: 1, y: 1 };
    let dest = Point { x: 39, y: 31 };

    println!("Part 1: {}", a_star(&maze, &start, &dest));
    println!(
        "Part 2: {}",
        distinct_location_count_in_max_50_steps(&maze, &start)
    );
}

// https://en.wikipedia.org/wiki/A*_search_algorithm
fn a_star(maze: &[[u8; MAZE_COLS]; MAZE_ROWS], start: &Point, dest: &Point) -> usize {
    let mut open_set: BinaryHeap<Reverse<Point>> = BinaryHeap::new();
    open_set.push(Reverse(*start));

    let mut g_score = [[std::usize::MAX; MAZE_COLS]; MAZE_ROWS];
    g_score[start.x][start.y] = 0;

    let mut f_score = [[std::usize::MAX; MAZE_COLS]; MAZE_ROWS];
    f_score[start.x][start.y] = heuristic(start, dest);

    while let Some(Reverse(current)) = open_set.pop() {
        if current == *dest {
            return g_score[current.x][current.y];
        }

        for neighbour in get_neighbours(maze, &current) {
            let tentative_g_score = g_score[current.x][current.y] + 1;

            if tentative_g_score < g_score[neighbour.x][neighbour.y] {
                g_score[neighbour.x][neighbour.y] = tentative_g_score;
                f_score[neighbour.x][neighbour.y] =
                    g_score[neighbour.x][neighbour.y] + heuristic(&neighbour, dest);
                open_set.push(Reverse(neighbour));
            }
        }
    }

    panic!("No path found");
}

fn heuristic(p1: &Point, p2: &Point) -> usize {
    p1.x.abs_diff(p2.x) + p1.y.abs_diff(p2.y)
}

fn distinct_location_count_in_max_50_steps(
    maze: &[[u8; MAZE_COLS]; MAZE_ROWS],
    start: &Point,
) -> usize {
    let mut visited: HashSet<Point> = HashSet::new();
    let mut queue: VecDeque<Point> = VecDeque::new();
    queue.push_back(*start);

    for _ in 0..50 {
        let mut new_queue: VecDeque<Point> = VecDeque::new();

        while let Some(point) = queue.pop_front() {
            for neighbour in get_neighbours(maze, &point) {
                if !visited.contains(&neighbour) {
                    visited.insert(neighbour);
                    new_queue.push_back(neighbour);
                }
            }
        }

        queue = new_queue;
    }

    visited.len()
}

fn get_neighbours(maze: &[[u8; MAZE_COLS]; MAZE_ROWS], point: &Point) -> Vec<Point> {
    let mut neighbours = Vec::new();

    for (dx, dy) in &[(0, 1), (0, -1), (1, 0), (-1, 0)] {
        let x = point.x as isize + dx;
        let y = point.y as isize + dy;

        if x >= 0 && x < MAZE_ROWS as isize && y >= 0 && y < MAZE_COLS as isize {
            let x = x as usize;
            let y = y as usize;

            if maze[x][y] == 0 {
                neighbours.push(Point { x, y });
            }
        }
    }
    neighbours
}

fn set_walls(maze: &mut [[u8; MAZE_COLS]; MAZE_ROWS]) {
    for x in 0..MAZE_ROWS {
        for y in 0..MAZE_COLS {
            if is_wall(y, x) {
                maze[x][y] = 1;
            } else {
                maze[x][y] = 0;
            }
        }
    }
}

fn is_wall(x: usize, y: usize) -> bool {
    let sum = x * x + 3 * x + 2 * x * y + y + y * y + FAVOURITE_NUMBER;
    sum.count_ones() % 2 == 1
}
