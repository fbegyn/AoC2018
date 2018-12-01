use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let file = File::open("input.txt").expect("Failed to open file");

    let mut _numbers: Vec<i32> = vec![];
    for line in BufReader::new(file).lines() {
        _numbers.push(
            line.expect("Failed to read the line")
                .parse()
                .expect("Failed to parse into i32"),
        );
    }

    let mut freq = 0;
    let mut freq_list: HashMap<i32, i32> = HashMap::new();
    let mut ok: bool = true;
    let mut first: bool = true;
    while ok {
        for v in &_numbers {
            freq += v;
            let c = freq_list.entry(freq).or_insert(1);
            if c >= &mut 2 {
                ok = false;
                break;
            }
            *c += 1;
        }
        if first {
            println!("Solution for problem 1 is {}", freq);
            first = false;
        }
    }
    println!("Solution for problem 2 is {}", freq);
}
