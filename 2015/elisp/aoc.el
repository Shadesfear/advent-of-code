;;; aoc.el --- Shared helpers for AoC 2015 -*- lexical-binding: t; -*-

(defconst aoc-inputs-dir
  (expand-file-name "../inputs/"
                    (file-name-directory (or load-file-name buffer-file-name))))

(defun aoc-read-input (day)
  "Return the puzzle input for DAY as a string."
  (with-temp-buffer
    (insert-file-contents
     (expand-file-name (format "day%02d.txt" day) aoc-inputs-dir))
    (string-trim-right (buffer-string))))

(defun aoc-run (day part1 part2)
  (let ((input (aoc-read-input day)))
    (princ (format "Part 1: %s\nPart 2: %s\n"
                   (funcall part1 input)
                   (funcall part2 input)))))

(provide 'aoc)
;;; aoc.el ends here
