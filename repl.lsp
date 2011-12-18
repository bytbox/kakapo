; Read-Eval-Print Loop for Kakapo

; First print some information about us.
(let ()
  (import "fmt")
  (fmt.Printf "Welcome to %s %s\n"
    -interpreter
    -interpreter-version)

; Now start the actual REPL
  (let ((readSexpr
        (lambda () (begin
                    (fmt.Print "kakapo> ")
                    (read)))))
    (define REPL
      (lambda ()
        (recover '(eof)
          (lambda ()
            (for 1 (print (eval (readSexpr)))))
          (lambda (_)
            (fmt.Println "Bye!")))))))
(REPL)
