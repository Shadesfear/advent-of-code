;;; .el --- I Was Told There Would Be No Math -*- lexical-binding: t; -*-

(require 'ert)
(require 'aoc (expand-file-name "aoc" (file-name-directory
                                       (or load-file-name buffer-file-name))))

(defun day08-parse (input)
  (split-string input "\n" t))

(defun day08-raw-length(s)
  (length s))

(defun day08-skip-length (c)
  (pcase c
    ((or  ?\" ?\\) 1)
    (?x 3)
    (_ 0)))


(defun day08-memory-length(s)
  (let ((cur 1)
        (end (- (length s) 1))
        (count 0))

    (while (< cur end)
      (cond ((equal (aref s cur) ?\\)
             (setq cur (+ cur (day08-skip-length (aref s (1+ cur)))))))
      (setq cur (1+ cur))
      (setq count (1+ count))
      )
    count
    )
  
  ) 

(defun day08-part1 (input)
  (apply #'+ (seq-map
              (lambda (line)
                (- (day08-raw-length line) (day08-memory-length line)))
              (day08-parse input)))

  )

(defun day08-expand (c)
  (pcase c
    ((or  ?\" ?\\) 1)
    (_ 0)))


(defun day08-expanded-length(s)
  (let ((cur 0)
        (end (length s))
        (count 2))
    (while (< cur end)
      (setq count (+ count (day08-expand (aref s cur))))
      (setq cur (1+ cur))
      (setq count (1+ count))
      )
    count)) 


(defun day08-part2 (input)
  (apply #'+ (seq-map
              (lambda (line)
                (- (day08-expanded-length line) (day08-raw-length line)))
              (day08-parse input)))
  )


(ert-deftest day08-raw-length ()
  (should (= 2 (day08-raw-length "\"\"")))
  (should (= 5 (day08-raw-length "\"abc\"")))
  (should (= 10 (day08-raw-length "\"aaa\\\"aaa\"")))
  (should (= 6 (day08-raw-length "\"\\x27\"")))
  
  )

(ert-deftest day08-memory()
  (should (= 0 (day08-memory-length "\"\"")))
  (should (= 3 (day08-memory-length "\"abc\"")))
  (should (= 7 (day08-memory-length "\"aaa\\\"aaa\"")))
  (should (= 1 (day08-memory-length "\"\\x27\"")))
  )


(unless load-file-name
  (aoc-run 8 #'day08-part1 #'day08-part2))

(provide 'day08)
;;; day08.el ends here
