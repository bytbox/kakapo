; Testing rig library

(let ((-section "UNDEFINED-SECTION"))
  (define S'
    (lambda (s)
      (define -section s)))

  (define T'
    (lambda (b)
      (if b nil
        (panic -section))))

  (define F'
    (lambda (b)
      (T' (not b)))))
