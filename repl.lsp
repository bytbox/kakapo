; Read-Eval-Print Loop for Kakapo

(let ()
  (import "fmt")
  (let ((readSexpr
        (lambda () (begin
                    (fmt.Print "> ")
                    (read)))))
    (define REPL
      (lambda ()
        (recover (eof)
          (for 1
            (print (eval (readSexpr))))
          nil)))))

(REPL)
