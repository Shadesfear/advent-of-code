
;;; .el --- I Was Told There Would Be No Math -*- lexical-binding: t; -*-

(require 'ert)
(require 'aoc (expand-file-name "aoc" (file-name-directory
                                       (or load-file-name buffer-file-name))))

(defun day04-parse (input)
  (seq-map (lambda (line)
             (seq-map #'string-to-number (split-string line "x")))
           (split-string input "\n" t)))


(defun day04 (key prefix)
  (let ((n 0))
    (while (not (string-prefix-p prefix (md5 (format "%s%d" key n))))
      (setq n (1+ n)))
    n))

(defun day04-part1 (input)
  (day04 input "00000")
  )

(defun day04-part2 (input)
  (day04 input "000000")
  )

(ert-deftest day04-part1 ()
  (should (= 609043 (day04-part1 "abcdef")))
  (should (= 1048970 (day04-part1 "pqrstuv"))))

(ert-deftest day04-part2 ())

(unless load-file-name
  (aoc-run 4 #'day04-part1 #'day04-part2))

(provide 'day04)
;;; day04.el ends here
