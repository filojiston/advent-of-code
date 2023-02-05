fn main() {
    let input = include_str!("input.txt");

    let mut valid_count_part1 = 0;
    let mut valid_count_part2 = 0;
    for lines in input.lines().collect::<Vec<_>>().chunks(3) {
        let sides = lines
            .iter()
            .flat_map(|x| x.split_whitespace())
            .map(|x| x.parse::<usize>().unwrap())
            .collect::<Vec<_>>();
        let sides_part2 = vec![
            vec![sides[0], sides[3], sides[6]],
            vec![sides[1], sides[4], sides[7]],
            vec![sides[2], sides[5], sides[8]],
        ]
        .into_iter()
        .flat_map(|x| x)
        .collect::<Vec<_>>();

        valid_count_part1 += valid_triangle_count(sides);
        valid_count_part2 += valid_triangle_count(sides_part2);
    }

    println!("Part 1: {}", valid_count_part1);
    println!("Part 2: {}", valid_count_part2);
}

fn valid_triangle_count(sides: Vec<usize>) -> usize {
    sides.chunks(3).filter(|x| valid_triangle(x)).count()
}

fn valid_triangle(sides: &[usize]) -> bool {
    sides[0] + sides[1] > sides[2]
        && sides[1] + sides[2] > sides[0]
        && sides[2] + sides[0] > sides[1]
}
