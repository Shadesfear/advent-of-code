use crate::solution;

fn part1(input: &str) -> String {
    input
        .split("\n\n")
        .map(|calories| {
            calories
                .lines()
                .filter_map(|s| s.parse::<u32>().ok())
                .sum::<u32>()
        })
        .max()
        .unwrap()
        .to_string()
}

fn part2(input: &str) -> String {
    let mut calories: Vec<u32> = input
        .split("\n\n")
        .map(|calories| {
            calories
                .lines()
                .filter_map(|s| s.parse::<u32>().ok())
                .sum::<u32>()
        })
        .collect();
    calories.sort();
    calories.into_iter().rev().take(3).sum::<u32>().to_string()
}

pub fn solution() -> solution::Solution {
    solution::Solution {
        day: 1,
        input: "src/data/day1.txt".to_string(),
        part1,
        part2,
    }
}
