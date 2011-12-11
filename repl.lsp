; Read-Eval-Print Loop for Kakapo

(recover (eof)
  (for 1 (print (eval (read))))
  nil)
