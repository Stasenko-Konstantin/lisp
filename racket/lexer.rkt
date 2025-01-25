#lang racket

(require "tokens.rkt")

(define x 0)
(define y 0)
(define symbols "\\|/?.><!#@`^~%&*-_+=;")

;; listof token, string, '(int, int), int
(struct lexer (tokens code span idx) #:mutable)

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
             (lexer-span-set! l '(-1, 1))]
            [(and (char=? c #\-) (char=? (string-ref code (+ i 1)) #\-))
             (while #t
                    (if (char=? (string-ref code (+ i 1)) #\newline)
                        (return)
                        (set! i (+ i 1))))]
            [() ()])))))