;;; day01-test.el -*- lexical-binding: t; -*-

(require 'ert)
(require 'day01)

(ert-deftest day01-part1 ()
  (should (= 0 (day01-part1 "(())")))
  (should (= 0 (day01-part1 "()()")))
  (should (= 3 (day01-part1 "(((")))
  (should (= 3 (day01-part1 "))(((((")))
  (should (= -1 (day01-part1 "())")))
  (should (= -3 (day01-part1 ")())())"))))

(ert-deftest day01-part2 ()
  (should (= 1 (day01-part2 ")")))
  (should (= 5 (day01-part2 "()())"))))
