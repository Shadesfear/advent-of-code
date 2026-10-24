;;; .el --- I Was Told There Would Be No Math -*- lexical-binding: t; -*-

(require 'ert)
(require 'aoc (expand-file-name "aoc" (file-name-directory
                                       (or load-file-name buffer-file-name))))

(defun day02-parse (input)
  (seq-map (lambda (line)
             (seq-map #'string-to-number (split-string line "x")))
           (split-string input "\n" t)))

(defun day02-paper (lwh)
  (seq-let (l w h) lwh
    (let ((lw (* l w))
          (wh (* w h))
          (hl (* h l)))
      (+ (* 2 lw) (* 2 wh) (* 2 hl) (min lw wh hl)))))

(defun day02-ribbon (lwh)
  (+ (* 2 (- (apply #'+ lwh) (apply #'max lwh)))
     (apply #'* lwh)))

(defun day02-part1 (input)
  (apply #'+ (seq-map #'day02-paper (day02-parse input))))

(defun day02-part2 (input)
  (apply #'+ (seq-map #'day02-ribbon (day02-parse input))))

(ert-deftest day02-part1 ()
  (should (= 58 (day02-part1 "2x3x4")))
  (should (= 43 (day02-part1 "1x1x10")))
  (should (= 101 (day02-part1 "2x3x4\n1x1x10"))))

(ert-deftest day02-part2 ()
  (should (= 34 (day02-part2 "2x3x4")))
  (should (= 14 (day02-part2 "1x1x10")))
  (should (= 48 (day02-part2 "2x3x4\n1x1x10"))))

(unless load-file-name
  (aoc-run 2 #'day02-part1 #'day02-part2))

(provide 'day02)
;;; day02.el ends here
