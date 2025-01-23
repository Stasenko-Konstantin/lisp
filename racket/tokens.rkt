#lang racket

(require "util.rkt")

(define token-types '(num name string lparen rparen))

(struct token (type content x y) #:mutable)

(define (token->string t)
  (concat-strings (let* [(type-sym (token-type t))
                         (type (symbol->string type-sym))]
                    (if (member type-sym token-types)
                        (concat-strings "type = " type ", ")
                        (error (concat-strings type ": is not a valid token type"))))
                  "content = " (token-content t)            ", "
                  "x = "       (number->string (token-x t)) ", "
                  "y = "       (number->string (token-y t)) ";\n"))

(define (token=? t1 t2)
  (if (eq? (token-type t1) (token-type t2))
      (string=? (token-content t1) (token-content t2))
      #f))