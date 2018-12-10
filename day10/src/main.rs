const PUZZLE: &str = include_str!("./input.txt");

fn main() {
    let points: Vec<&str> = PUZZLE.lines().collect::<Vec<_>>();

    let mut coord: Vec<Vec<i32>> = Vec::new();
    let mut vel: Vec<Vec<i32>> = Vec::new();

    for p in points {
        coord.push(
            p[10..24]
                .trim()
                .split(", ")
                .map(|s| s.trim().parse::<i32>().expect(""))
                .collect::<Vec<i32>>(),
        );
        vel.push(
            p[36..42]
                .trim()
                .split(", ")
                .map(|s| s.trim().parse::<i32>().expect(""))
                .collect::<Vec<i32>>(),
        );
    }

    for _ in 1..=10036 {
        coord = coord
            .iter()
            .zip(vel.iter())
            .map(|(c, v)| vec![c[0] + v[0], c[1] + v[1]])
            .collect();
    }

    let mut grid: [[char;200];200] = [[' ';200];200];

    for c in coord.iter() {
        let y = c[1] as usize;
        let x = c[0] as usize;
        grid[y][x] = '#'
    }

    for y in grid.iter() {
        for x in y.iter() {
            print!("{}", x);
        }
        println!("");
    }
}
