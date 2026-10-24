;;; .el --- I Was Told There Would Be No Math -*- lexical-binding: t; -*-

(require 'ert)
(require 'aoc (expand-file-name "aoc" (file-name-directory
                                       (or load-file-name buffer-file-name))))

(defun day07-token (s)
  (if (string-match-p "\\`[0-9]+\\'" s) (string-to-number s) s))

(defun day07-value (s map)
  (or (gethash s map) s))

(defun day07-ready-p (operand map)
  (or (numberp operand) (gethash operand map)))

(defun day07-parse-line (line)
  (pcase-let ((`(,lhs ,wire) (split-string line " -> ")))
    (cons wire
          (pcase (mapcar #'day07-token (split-string lhs))
            (`(,a) (list 'set a))
            (`("NOT" ,a) (list 'not a))
            (`(,a "AND" ,b) (list 'and a b))
            (`(,a "OR" ,b) (list 'or a b))
            (`(,a "LSHIFT" ,n) (list 'lshift a n))
            (`(,a "RSHIFT" ,n) (list 'rshift a n))
            (_ (error "Bad line: %s" line))))))

(defun day07-do-instruction (instr map)
  (pcase instr
    (`(,wire set ,a) (puthash wire (day07-value a map) map))
    (`(,wire not ,a) (puthash wire (logand #xFFFF (lognot (day07-value a map))) map))
    (`(,wire and ,a ,b) (puthash wire (logand (day07-value a map) (day07-value b map) ) map))
    (`(,wire or ,a ,b) (puthash wire (logior (day07-value a map) (day07-value b map) ) map))
    (`(,wire lshift ,a ,n) (puthash wire (logand #xFFFF (ash (day07-value a map) n)) map))
    (`(,wire rshift ,a ,n) (puthash wire (logand #xFFFF (ash (day07-value a map) (- n))) map))
    ))


(defun day07-can-run-instr-p (instr map)
  (pcase instr
    (`(,_ ,(or 'set 'not 'lshift 'rshift) ,a . ,_) (day07-ready-p a map))
    (`(,_ ,(or 'and 'or) ,a ,b) (and (day07-ready-p a map) (day07-ready-p b map)))
    (_ (error "Bad instructions: %S" instr))))

(defun day07-parse (input)
  (seq-map (lambda (line)
             (day07-parse-line line)
             )
           (split-string input "\n" t)))

(defun day07-run-instrs (instrs map)
  (if instrs
      (let* ((to-run (seq-filter
                      (lambda (instr) (day07-can-run-instr-p instr map)) instrs))
             (next-items (seq-remove
                          (lambda (instr) (day07-can-run-instr-p instr map)) instrs)))

        (unless to-run
          (error "nothing to run"))

        (dolist (instr to-run)
          (day07-do-instruction instr map))

        (day07-run-instrs next-items map)

        )
    map

    ))

(defun day07-part1 (input)
  (let ((instrs (day07-parse input))
        (map (make-hash-table :test #'equal)))
    (gethash "a" (day07-run-instrs instrs map))
    
    )

  )

(defun day07-part2 (input)

  (let ((instrs (day07-parse input))
        (map (make-hash-table :test #'equal)))
    (let ((part1-res (gethash "a" (day07-run-instrs instrs map))))
      (setq map (make-hash-table :test #'equal))
      (setq instrs (seq-remove (lambda (instr) (equal (car instr) "b")) instrs))
      (puthash "b" part1-res map)
      (gethash "a" (day07-run-instrs instrs map))
      )
    )
  )

(ert-deftest day07-part1 ()
  
  
  )

(ert-deftest day07-part2 ()
  
  
  )

(unless load-file-name
  (aoc-run 7 #'day07-part1 #'day07-part2))

(provide 'day07)
;;; day07.el ends here
