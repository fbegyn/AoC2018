const PUZZLE: &str = include_str!("./input.txt");

fn main() {
    let ids: Vec<String> = PUZZLE
        .lines()
        .filter_map(|s| s.parse::<String>().ok())
        .collect::<Vec<_>>();
    prob1(ids);
    //prob2(ids);
}

fn prob1(ids: Vec<String>){
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
