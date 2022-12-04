use crate::solution;

const ALPHABET: &str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

fn part1(input: &str) -> String {
    input
        .lines()
        .map(|line| {
            let chars: Vec<char> = line.chars().collect();
            let chunks: Vec<&[char]> = chars.chunks(line.len() / 2).collect();
            let contained = chunks[0]
                .into_iter()
                .filter(|c| chunks[1].contains(c))
                .collect::<Vec<&char>>()[0];
            (ALPHABET.find(*contained).unwrap() + 1) as u32
        })
        .sum::<u32>()
        .to_string()
}

fn part2(input: &str) -> String {
    let lines: Vec<&str> = input.lines().collect();
    let groups: Vec<&[&str]> = lines.chunks(3).collect();
    let contained: Vec<u32> = groups
        .into_iter()
        .map(|v| {
            v[0].chars()
                .filter(|c| v[1].contains(*c) && v[2].contains(*c))
                .map(|c| (ALPHABET.find(c).unwrap() + 1) as u32)
                .last()
                .unwrap()
        })
        .collect();
    contained.into_iter().sum::<u32>().to_string()
    // println!("{:?}", groups);
}

pub fn solution() -> solution::Solution {
    solution::Solution {
        day: 3,
        input: "src/data/day3.txt".to_string(),
        part1,
        part2,
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_test() {
        let inp = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";
        assert!(part1(inp) == "157")
    }

    #[test]
    fn part2_test() {
        let inp = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";
        println!("{:?}", part2(inp));
    }
}
