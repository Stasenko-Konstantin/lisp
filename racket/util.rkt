#lang racket

(provide concat-strings)

(define (concat-strings . lst)
  (string-join lst ""))