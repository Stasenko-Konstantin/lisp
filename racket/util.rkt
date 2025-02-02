#lang racket

(provide (all-defined-out))

(define (concat-strings . lst)
  (string-join lst ""))