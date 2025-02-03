#lang racket

(provide (all-defined-out))

(define errors '())
(define interpretation-fault? #f)
(define repl? #f)

(define (add-err! err)
  (set! interpretation-fault? #t)
  (set! errors (cons errors err)))

(define (print-errors)
  (displayln "Interpretation Fault:")
  (for ([(n e) (in-indexed errors)])
    (error (format "\t~a, ~s" n e)))
  (unless repl?
    (exit 0)))

(define (reset-errors)
  (set! interpretation-fault? #f)
  (set! errors '()))