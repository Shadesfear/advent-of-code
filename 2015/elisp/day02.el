;;; day02.el --- I Was Told There Would Be No Math -*- lexical-binding: t; -*-

(require 'aoc)

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

(provide 'day02)
;;; day02.el ends here
