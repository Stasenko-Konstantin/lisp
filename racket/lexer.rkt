#lang racket

(require "tokens.rkt")

(provide scan)

(define x 0)
(define y 0)
(define symbols "\\|/?.><!#@`^~%&*-_+=;")

;;           '(token) string '(int int) int
(struct lexer (tokens code     span     idx) #:mutable)

(define (add l ttype content)
  (set-lexer-tokens! l (cons (lexer-tokens l) (token ttype content))))

(define (scan #:code [code '()] #:lexer [l '()])
  (if (null? code)
      (scan #:lexer (lexer '() code '(0 0) 0))
      (begin
        (when (null? l) (error "lexer is nil"))
        (for ([(i c) (in-indexed code)])
          (cond
            [(or (char=? c #\() (char=? c #\[))
             (add l 'lparen "(")]
            [(or (char=? c #\)) (char=? c #\]))
             (add l 'rparen ")")]
            [(or (char=? c #\return) (char=? c #\tab) (char=? c #\space))
             (void)]
            [(char=? c #\newline)
             (set-lexer-span! l '(-1, 1))]
            [(and (char=? c #\-) (char=? (string-ref code (+ i 1)) #\-))
             (for ([i (in-naturals)]) #:break (char=? (string-ref code (+ i 1)) #\newline)
               (set! i (+ i 1)))]
            ;[() ()]
            )))))