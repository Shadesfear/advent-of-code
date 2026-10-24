;;; .el --- I Was Told There Would Be No Math -*- lexical-binding: t; -*-

(require 'ert)
(require 'aoc (expand-file-name "aoc" (file-name-directory
                                       (or load-file-name buffer-file-name))))

(defun day09-parse (input)
  (split-string input "\n" t))

(defun perms (l)
  (if (null l)
      (list nil)
    (cl-loop for x in l
             append (mapcar (lambda (p) (cons x p) )
                            (perms (remq x l)))))
  )


(defun day09-part1 (input)
  

  )

(defun day09-part2 (input)
  
  )


(ert-deftest day09-raw-length ()
  (should (= 2 (day09-raw-length "\"\"")))
  
  )

(ert-deftest day09-memory()
  (should (= 0 (day09-memory-length "\"\"")))
  )


(unless load-file-name
  (aoc-run 9 #'day09-part1 #'day09-part2))

(provide 'day09)
;;; day09.el ends here
