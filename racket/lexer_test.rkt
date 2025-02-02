#lang racket

(require rackunit "tokens.rkt" "lexer.rkt")

(check
 (foldl (lambda (acc x) (and acc (token=? (car x) (cdr x))))
        (map list
             (scan #:code "123 hello (]\n\"hello\"")
             (list ;; type   content x  y
              (token 'num    "123"   0  0)
              (token 'name   "hello" 4  0)
              (token 'lparen "("     10 0)
              (token 'rparen ")"     11 0)
              (token 'string "hello" 0  1))))
 "dscrpt")