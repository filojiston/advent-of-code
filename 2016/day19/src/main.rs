const INPUT: u32 = 3005290;

fn main() {
    println!("Part 1: {}", josephus(INPUT));
    println!("Part 2: {}", josephus_directly_across(INPUT));
}

// related: https://en.wikipedia.org/wiki/Josephus_problem
fn josephus(n: u32) -> u32 {
    2 * (n - (1 << (32 - n.leading_zeros() - 1))) + 1
}

fn josephus_directly_across(n: u32) -> u32 {
    let mut power_of_three = 1;
    while power_of_three * 3 <= n {
        power_of_three *= 3;
    }
    n - power_of_three
}
