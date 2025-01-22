#lang racket

(require racket/base racket/path "util.rkt")

(define out ">>> ")
(define in  "<<< ")

(define usage (concat-strings (list "usage:\tracket main.rkt [*.scm] [arg] where arg:\n"
                               "\t--help \t -- prints the help\n"
                               "\t--past \t -- prints intermediate ast")))

(define help  (concat-strings (list out ":l, :load *.scm \t -- evaluate file\n"
                               out ":h, :help \t\t -- prints the help\n"
                               out ":p, :past \t\t -- prints intermediate ast\n"
                               out ":q, :quit \t\t -- program exit")))

(define print-ast #f)
;; (define env nil)

(define (eval code #:need-eval need-eval) ;; todo
  (displayln need-eval)
  (displayln code))

(define (eval-file filepath)
  (let* ([ext  (bytes->string/utf-8 (path-get-extension (string->path filepath)))]
         [in   (open-input-file filepath)]
         [file (port->string in)])
    (begin
      (unless (string=? ".scm" ext) (error "its not *.scm"))
      (close-input-port in)
      (eval (concat-strings (list "(" file ")")) #:need-eval #t))))

(define (repl)
  (displayln "todo repl"))

(command-line
 #:program "lisp"
 #:once-each
 [("-p" "--past") "Prints intermediate ast"
  (set! print-ast #t)]
 #:args args
 (begin
;;   (set! env (new-env make-builtins)) ;; todo make-builtins
   (cond
    [(null? args)
     (repl)]
    [(= (length args) 1)
     (eval-file (car args))]
    [else
     (displayln usage)])))