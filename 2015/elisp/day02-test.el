;;; day02-test.el -*- lexical-binding: t; -*-

(require 'ert)
(require 'day02)

(ert-deftest day02-part1 ()
  (should (= 58 (day02-part1 "2x3x4")))
  (should (= 43 (day02-part1 "1x1x10")))
  (should (= 101 (day02-part1 "2x3x4\n1x1x10"))))

(ert-deftest day02-part2 ()
  (should (= 34 (day02-part2 "2x3x4")))
  (should (= 14 (day02-part2 "1x1x10")))
  (should (= 48 (day02-part2 "2x3x4\n1x1x10"))))
