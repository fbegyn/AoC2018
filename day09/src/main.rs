use std::collections::VecDeque;

const PUZZLE: &str = include_str!("./input.txt");

fn main() {
    let input: String = PUZZLE.lines().collect::<String>();
    let inp: Vec<&str> = input.split_whitespace().collect::<Vec<&str>>();
    let players: usize = inp[0].parse().expect("Unable to parse plaeys");
    let marbles: usize = inp[6].parse().expect("Unable to parse high value");
    println!("Players: {} -- Highest value: {}", players, marbles);

    let mut circle: VecDeque<usize> = VecDeque::with_capacity(marbles);
    circle.push_back(0);
    let mut scores = vec![0; players];

    for marble in 1..=marbles * 100 {
        if marble % 23 == 0 {
            scores[marble % players] += marble;
            for _ in 0..7 {
                let back = circle.pop_back().unwrap();
                circle.push_front(back);
            }
            scores[marble % players] += circle.pop_front().unwrap();
        } else {
            for _ in 0..2 {
                let front = circle.pop_front().unwrap();
                circle.push_back(front);
            }
            circle.push_front(marble)
        }
        if marble == marbles {
            println!("Max score part 1: {}", scores.iter().max().unwrap());
        }
    }
    println!("Max score part 2: {}", scores.iter().max().unwrap())
}
