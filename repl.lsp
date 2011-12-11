; Read-Eval-Print Loop for Kakapo

(recover (quote (eof))
  (for 1 (print (eval (read))))
  (lambda (key) nil))
