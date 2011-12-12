; Read-Eval-Print Loop for Kakapo

(let ()
  (import "fmt")
  (let ((readSexpr
        (lambda () (begin
                    (fmt.Print "> ")
                    (read)))))
    (define REPL
      (lambda ()
        (recover (quote (eof))
          (lambda ()
            (for 1
              (print (eval (readSexpr)))))
          (lambda () (fmt.Println "Bye!")))))))

(REPL)
