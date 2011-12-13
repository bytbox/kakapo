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
(T' (>))
(T' (> 1))
(T' (> 2 1))
(F' (> 1 2))
(T' (> 2e5 3e4))
(F' (> 1 2 3))
(T' (> 3 2 1))
(F' (> 3 1 2))

(S' "Less-than")
(T' (<))
(T' (< 1))
(F' (< 2 1))
(T' (< 1 2))
(F' (< 2e5 3e4))
(T' (< 1 2 3))
(F' (< 3 2 1))
(F' (< 3 1 2))

(S' "Greater-than-or-equal")
(T' (>=))

(S' "Less-than-or-equal")
(T' (<=))

(S' "Addition")
(T' (= (+ 1 1) 2))
(T' (= (* 5 4) 20))

(S' "Subtraction")
(T' (= -1 (- 1)))
(T' (= -1 (- 1 2)))

