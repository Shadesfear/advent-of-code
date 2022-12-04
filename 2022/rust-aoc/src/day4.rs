use crate::solution::Solution;

fn pair_str2u32(pair: (&str, &str)) -> (u32, u32) {
    let (f, s) = pair;
    (f.parse::<u32>().unwrap(), s.parse::<u32>().unwrap())
}

fn part1(input: &str) -> String {
    input
        .lines()
        .map(|line| {
            let (first, second) = line.split_once(",").unwrap();
            let (first_lower, first_upper) = pair_str2u32(first.split_once("-").unwrap());
            let (second_lower, second_upper) = pair_str2u32(second.split_once("-").unwrap());
            if first_lower >= second_lower && first_upper <= second_upper {
                1
            } else if second_lower >= first_lower && second_upper <= first_upper {
                1
            } else {
                0
            }
        })
        .sum::<u32>()
        .to_string()
}

fn part2(input: &str) -> String {
    input
        .lines()
        .map(|line| {
            let (first, second) = line.split_once(",").unwrap();

            let (a1, a2) = pair_str2u32(first.split_once("-").unwrap());
            let (b1, b2) = pair_str2u32(second.split_once("-").unwrap());

            match a2.max(b2) - a1.min(b1) <= (a2 - a1) + (b2 - b1) {
                true => 1,
                false => 0,
            }
        })
        .sum::<u32>()
        .to_string()
}

pub fn solution() -> Solution {
    Solution::new(4, "src/data/day4.txt", part1, part2)
}

mod tests {
    use super::*;

    const INP: &str = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";

    #[test]
    fn test_part1() {
        println!("{:?}", part1(INP));

        // assert_eq!(part1(""), "");
    }

    #[test]
    fn test_part2() {
        println!("{:?}", part2(INP));
    }
}
