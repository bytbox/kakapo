; Basic arithmetic

(S' "Equality")
(T' (=))
(T' (= 1 1))
(F' (= 1 2))

(S' "Inequality")
(T' (/=))
(T' (/= 1 2 3))
(F' (/= 1 1 3))

(S' "Greater-than")
;(T' nil)

(S' "Addition")
(T' (= (+ 1 1) 2))
(T' (= (* 5 4) 20))

(S' "Subtraction")
(T' (= -1 (- 1)))
(T' (= -1 (- 1 2)))

