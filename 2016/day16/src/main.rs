const FIRST_DISK_LENGTH: usize = 272;
const SECOND_DISK_LENGTH: usize = 35651584;
const INPUT: &str = "10111100110001111";

fn main() {
    println!(
        "Part 1: {}",
        checksum(fill_disk(INPUT, FIRST_DISK_LENGTH).as_str())
    );
    println!(
        "Part 2: {}",
        checksum(fill_disk(INPUT, SECOND_DISK_LENGTH).as_str())
    );
}

fn fill_disk(initial: &str, length: usize) -> String {
    let mut disk = initial.to_string();
    while disk.len() < length {
        let b = disk
            .chars()
            .rev()
            .map(|c| if c == '1' { '0' } else { '1' })
            .collect::<String>();
        disk.push('0');
        disk.push_str(&b);
    }
    disk.truncate(length);
    disk
}

fn checksum(data: &str) -> String {
    let mut result = data
        .chars()
        .collect::<Vec<char>>()
        .chunks(2)
        .map(|chunk| if chunk[0] == chunk[1] { '1' } else { '0' })
        .collect::<String>();
    if result.len() % 2 == 0 {
        result = checksum(result.as_str());
    }
    result
}
