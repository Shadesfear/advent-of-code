use crate::solution;

fn part1(input: &str) -> String {
    let calories: Vec<Vec<i32>> = input.lines().fold(Vec::new(), |mut acc, line| {
        let n = line;
        if n == "" || acc.is_empty() {
            acc.push(Vec::new());
        } else {
            acc.last_mut().unwrap().push(n.parse().unwrap())
        }
        acc
    });
    calories
        .into_iter()
        .map(|l| l.into_iter().sum::<i32>())
        .max()
        .unwrap()
        .to_string()
}

fn part2(input: &str) -> String {
    let calories: Vec<Vec<i32>> = input.lines().fold(Vec::new(), |mut acc, line| {
        let n = line;
        if n == "" || acc.is_empty() {
            acc.push(Vec::new());
        } else {
            acc.last_mut().unwrap().push(n.parse().unwrap())
        }
        acc
    });
    let mut max_calories = calories
        .into_iter()
        .map(|l| l.into_iter().sum::<i32>())
        .collect::<Vec<i32>>();

    max_calories.sort();

    max_calories[max_calories.len() - 3..]
        .into_iter()
        .sum::<i32>()
        .to_string()
}

pub fn solution() -> solution::Solution {
    solution::Solution {
        day: 1,
        input: "src/data/day1.txt".to_string(),
        part1,
        part2,
    }
}
