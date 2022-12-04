use crate::solution::Solution;

#[derive(Debug)]
enum Move {
    Rock,
    Paper,
    Scissor,
}

enum Outcome {
    Win,
    Loss,
    Draw,
}

fn who_wins(player1: &Move, player2: &Move) -> Outcome {
    match (player1, player2) {
        (Move::Rock, Move::Rock) => Outcome::Draw,
        (Move::Rock, Move::Paper) => Outcome::Loss,
        (Move::Rock, Move::Scissor) => Outcome::Win,
        (Move::Paper, Move::Rock) => Outcome::Win,
        (Move::Paper, Move::Paper) => Outcome::Draw,
        (Move::Paper, Move::Scissor) => Outcome::Loss,
        (Move::Scissor, Move::Rock) => Outcome::Loss,
        (Move::Scissor, Move::Paper) => Outcome::Win,
        (Move::Scissor, Move::Scissor) => Outcome::Draw,
    }
}

fn move_from_outcome(mov: &Move, outcome: &Outcome) -> Move {
    match outcome {
        Outcome::Win => match mov {
            Move::Rock => Move::Paper,
            Move::Paper => Move::Scissor,
            Move::Scissor => Move::Rock,
        },
        Outcome::Loss => match mov {
            Move::Rock => Move::Scissor,
            Move::Paper => Move::Rock,
            Move::Scissor => Move::Paper,
        },
        Outcome::Draw => match mov {
            Move::Rock => Move::Rock,
            Move::Paper => Move::Paper,
            Move::Scissor => Move::Scissor,
        },
    }
}

fn score(mov: &Move, outcome: &Outcome) -> i32 {
    let move_score = match mov {
        Move::Rock => 1,
        Move::Paper => 2,
        Move::Scissor => 3,
    };
    let outcome_score = match outcome {
        Outcome::Win => 6,
        Outcome::Loss => 0,
        Outcome::Draw => 3,
    };
    move_score + outcome_score
}

fn moves_mapper(mov: &str) -> Move {
    match mov {
        "A" | "X" => Move::Rock,
        "B" | "Y" => Move::Paper,
        "C" | "Z" => Move::Scissor,
        _ => panic!("Not a move"),
    }
}

fn real_moves_mapper(mov: &str) -> Move {
    match mov {
        "A" => Move::Rock,
        "B" => Move::Paper,
        "C" => Move::Scissor,
        _ => panic!("Not a move"),
    }
}

fn outcomes_mapper(outcome: &str) -> Outcome {
    match outcome {
        "X" => Outcome::Loss,
        "Y" => Outcome::Draw,
        "Z" => Outcome::Win,
        _ => panic!("Not a move"),
    }
}

fn part1(input: &str) -> String {
    let parseed: i32 = input
        .lines()
        .map(|l| {
            let val = l;
            let moves: Vec<Move> = val.split(" ").map(moves_mapper).collect();
            let outcome = who_wins(&moves[1], &moves[0]);
            let score = score(&moves[1], &outcome);
            score
        })
        .into_iter()
        .sum::<i32>();
    parseed.to_string()
}

fn part2(input: &str) -> String {
    let parseed: i32 = input
        .lines()
        .map(|l| {
            let val = l;
            let moves: Vec<&str> = val.split(" ").collect();
            let outcome = outcomes_mapper(moves[1]);
            let opponent_move = real_moves_mapper(moves[0]);
            let my_move = move_from_outcome(&opponent_move, &outcome);
            let score = score(&my_move, &outcome);

            score
        })
        .into_iter()
        .sum::<i32>();

    parseed.to_string()
}

pub fn solution() -> Solution {
    Solution::new(2, "src/data/day2.txt", part1, part2)
}
