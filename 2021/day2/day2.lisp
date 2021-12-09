
(defparameter lii (uiop:read-file-lines "day2_input.txt"))

(defun get-instruction (str)
  (first (sb-unicode:words str)))

(defun get-amount (str)
  (third (sb-unicode:words str)))

(defun rec (inst horizontal depth)

  (cond ((= (length inst) 0) (* horizontal depth))
        ((string= (get-instruction (car inst)) "forward") (rec (cdr inst) (+ horizontal (parse-integer (get-amount (car inst)))) depth))
        ((string= (get-instruction (car inst)) "down") (rec (cdr inst) horizontal (- depth (parse-integer (get-amount (car inst))))))
        ((string= (get-instruction (car inst)) "down") (rec (cdr inst) horizontal (+ depth (parse-integer (get-amount (car inst))))))

        )
)

(defun rec2 (i h d a)

  (cond  ((= (length i) 0) (* h d))
         ((string= (get-instruction (car i)) "forward") (rec2 (cdr i) (+ h (parse-integer (get-amount (car i)))) (+ d (* a (parse-integer (get-amount (car i))))) a))
         ((string= (get-instruction (car i)) "down") (rec2 (cdr i) h d (+ a (parse-integer (get-amount (car i))))))
         ((string= (get-instruction (car i)) "up") (rec2 (cdr i) h d (- a (parse-integer (get-amount (car i))))))

         )
  )

(defun part1 (instructions)
    (rec insructions 0 0)
 )

(defun part2 (i)
    (rec2 i 0 0 0)
 )
