; Testing rig library

(define T'
  (lambda (b)
    (if b nil
      (panic "FAILED"))))

(define F'
  (lambda (b)
    (T' (not b))))
