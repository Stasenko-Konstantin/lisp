#lang racket

(require rackunit "lexer.rkt" "tokens.rkt" "util.rkt")

(check equal? #t
 (for/and ([pair (map list
             (scan #:code "123 hello (]\n\"hello\"")
             (list ;; type   content x  y
              (token 'num    "123"   0  0) ;;    1
              (token 'name   "hello" 4  0) ;;    2
              (token 'lparen "("     10 0) ;;    3
              (token 'rparen ")"     11 0) ;;    4
              (token 'string "hello" 0  1)))] ;; 5
           [i (in-naturals)])
   (begin
     (define t1 (car pair))
     (define t2 (list-ref pair 1))
     (define p (token==? t1 t2))
     (unless p (displayln
                (concat-strings "\n" (format "~a" (+ i 1)) ":\n"
                                (token->string t1) "  ≠\n" (token->string t2))))
     p))
 "lexer test")