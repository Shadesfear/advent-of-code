;;; .el --- I Was Told There Would Be No Math -*- lexical-binding: t; -*-

(require 'ert)
(require 'aoc (expand-file-name "aoc" (file-name-directory
                                       (or load-file-name buffer-file-name))))

(defun is-vowel-p (c)
  (member c '(?a ?e ?i ?u ?o)))

(defun has-bad-strings-p (input)
  (let ((bad-strings (list "ab" "cd" "pq" "xy")))
    (seq-some (lambda (bad) (string-search bad input))
              bad-strings)))

(defun atleast-3-vowels-p (input)
  (>= (length (remq nil (seq-map #'is-vowel-p input))) 3))

(defun double-letter-p (input)
  (let* ((f (car input))
         (s (car (cdr input))))
    (cond ((or (null f) (null s)) nil)
          ((equal f s) t )
          (t (double-letter-p (cdr input))))))

(defun not-naughty-p (input)
  (and
   (not (has-bad-strings-p input))
   (atleast-3-vowels-p input)
   (double-letter-p (string-split input "" t))))

(defun day05-part1 (input)
  (length (remq nil (seq-map #'not-naughty-p (string-split input "\n" t)))))



(defun not-naughty-part2-p (input)
  (and (string-match-p "\\(..\\).*\\1" input) (string-match-p "\\(.\\).\\1" input)))

(defun day05-part2 (input)
  (length (remq nil (seq-map #'not-naughty-part2-p (string-split input "\n" t)))))

(ert-deftest day05-part1 ()
  )

(ert-deftest day05-part2 ()
  )

(unless load-file-name
  (aoc-run 5 #'day05-part1 #'day05-part2))

(provide 'day05)
;;; day05.el ends here
