const PUZZLE: &str = include_str!("./input.txt");

fn main() {
    let ids: Vec<&str> = PUZZLE.lines().collect::<Vec<_>>();
    prob1(ids.clone());
    prob2(ids);
}

fn prob1(ids: Vec<&str>) {
    let mut twos = 0;
    let mut threes = 0;
    for id in ids {
        if id.chars().any(|ch| id.matches(ch).count() == 2) {
            twos += 1;
        }
        if id.chars().any(|ch| id.matches(ch).count() == 3) {
            threes += 1;
        }
    }
    let checksum: i32 = twos * threes;
    println!("The checksum is: {}", checksum);
}

fn prob2(ids: Vec<&str>) {
    let mut t: std::string::String = "No match found".to_string();
    for (index, id) in ids.iter().enumerate() {
        for comp in ids.iter().skip(index + 1) {
            // zip combines multiple iterators together. Zie |a,b|
            if id.chars().zip(comp.chars()).filter(|(a, b)| a != b).count() == 1 {
                t = id
                    .chars()
                    .zip(comp.chars())
                    .filter(|(a, b)| a == b)
                    .map(|(a, _)| a.to_string())
                    .collect();
            }
        }
    }
    println!("The remaining letters are:\n{}", t);
}
