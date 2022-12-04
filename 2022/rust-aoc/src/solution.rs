use std::path::Path;

pub struct Solution {
    pub day: u8,
    pub input: String,
    pub part1: fn(&str) -> String,
    pub part2: fn(&str) -> String,
}

impl Solution {
    pub fn new(
        day: u8,
        input: &str,
        part1: fn(&str) -> String,
        part2: fn(&str) -> String,
    ) -> Solution {
        Solution {
            day,
            input: input.to_string(),
            part1,
            part2,
        }
    }

    fn run(&self, part: u8) {
        if Path::new(&self.input).exists() {
            let input = std::fs::read_to_string(&self.input).unwrap();
            let result = match part {
                1 => (self.part1)(&input),
                2 => (self.part2)(&input),
                _ => panic!("Invalid part"),
            };
            println!("Day {} Part {} Result: {}", self.day, part, result);
        } else {
            println!("Input file not found");
        }
    }

    pub fn run_all(&self) {
        println!("-------------------------");
        self.run(1);
        self.run(2);
        println!("-------------------------");
    }
}
