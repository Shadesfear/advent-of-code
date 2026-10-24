;;; .el --- I Was Told There Would Be No Math -*- lexical-binding: t; -*-

(require 'ert)
(require 'aoc (expand-file-name "aoc" (file-name-directory
                                       (or load-file-name buffer-file-name))))

(defun day06-parse (input)
  (seq-map (lambda (line)
             (let* ((splits (split-string line " "))
                    (r (reverse splits))
                    (action (nth 3 r))
                    (from (string-split (nth 2 r) ","))
                    (to (string-split (nth 0 r) ",")))
               (list :action action
                     :from (list :x (string-to-number (nth 0 from)) :y (string-to-number (nth 1 from)))
                     :to (list :x (string-to-number (nth 0 to)) :y (string-to-number (nth 1 to))))))
           (split-string input "\n" t)))

(defun idx (x y w)
  (+ x  (* y w)))

(defun run-instr (instr grid)
  (let (
        (fx (plist-get (plist-get instr :from) :x))
        (fy (plist-get (plist-get instr :from) :y))
        (tx (plist-get (plist-get instr :to) :x))
        (ty (plist-get (plist-get instr :to) :y))
        (action (plist-get instr :action))
        )
    (cl-loop for y from fy to ty do
             (cl-loop for x from fx to tx do
                      (cond
                       ((string= action "on") (aset grid (idx x y 1000) 1))
                       ((string= action "off") (aset grid (idx x y 1000) 0))
                       ((string= action "toggle") (aset grid (idx x y 1000) (if (= (aref grid (idx x y 1000)) 0) 1 0)))
                       )
                      ))))

(defun run-instr-part2 (instr grid)
  (let (
        (fx (plist-get (plist-get instr :from) :x))
        (fy (plist-get (plist-get instr :from) :y))
        (tx (plist-get (plist-get instr :to) :x))
        (ty (plist-get (plist-get instr :to) :y))
        (action (plist-get instr :action))
        )
    (cl-loop for y from fy to ty do
             (cl-loop for x from fx to tx do
                      (let ((i (idx x y 1000)))

                        (cond
                         ((string= action "on") (cl-incf (aref grid i)))
                         ((string= action "off") (setf (aref grid i) (max 0 (1- (aref grid (idx x y 1000))))))
                         ((string= action "toggle") (cl-incf (aref grid i) 2))
                         )
                        )
                      ))))
;; (seq-map (lambda (x) (run-instr x grid)) instrs)

(defun day06-part1 (input)
  (let* ((grid (make-vector 1000000 0))
         (instrs (day06-parse input)))

    (seq-map (lambda (x) (run-instr x grid)) instrs)
    (seq-reduce #'+ grid 0 )
    ))

(defun day06-part2 (input)
  (let* ((grid (make-vector 1000000 0))
         (instrs (day06-parse input)))

    (seq-map (lambda (x) (run-instr-part2 x grid)) instrs)
    (seq-reduce #'+ grid 0 )
    )
  )

(ert-deftest day06-part1 ()
  )

(ert-deftest day06-part2 ()
  )

(unless load-file-name
  (aoc-run 6 #'day06-part1 #'day06-part2))

(provide 'day06)
;;; day06.el ends here
