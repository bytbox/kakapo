; Read-Eval-Print Loop for Kakapo

(let () 
  (import "fmt")
  (define REPL
    (lambda ()
      (recover (eof)
        (for 1 
          (print (eval (read))))
        nil))))

(REPL)
