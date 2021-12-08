
(defvar li (map 'list #'parse-integer (uiop:read-file-lines "day1_input.txt")) "list")

(defun part1 (numbers)
  (loop for (a b) on numbers
    when b count (< a b)))

(defun part2 (numbers)
  (part1 (loop for (a b c) on numbers
             when c collect (+ a b c))))
