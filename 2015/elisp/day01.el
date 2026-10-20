;;; day01.el --- Not Quite Lisp -*- lexical-binding: t; -*-

(require 'ert)
(require 'aoc (expand-file-name "aoc" (file-name-directory
                                (or load-file-name buffer-file-name))))

(defun day01-step (c)
  (if (eq c ?\() 1 -1))

(defun day01-part1 (input)
  (seq-reduce (lambda (floor c) (+ floor (day01-step c))) input 0))

(defun day01-part2 (input)
  (let ((floor 0)
        (pos 0))
    (while (and (>= floor 0) (< pos (length input)))
      (setq floor (+ floor (day01-step (aref input pos))))
      (setq pos (1+ pos)))
    pos))

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

(unless load-file-name
  (aoc-run 1 #'day01-part1 #'day01-part2))

(provide 'day01)
;;; day01.el ends here
