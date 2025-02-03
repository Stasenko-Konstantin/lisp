#lang racket

(require "tokens.rkt" "errors.rkt" "util.rkt")

(provide scan)

(define x 0)
(define y 0)
(define symbols "\\|/?.><!#@`^~%&*-_+=;")

;;           '(token) string '(int int) int
(struct lexer (tokens code     span     idx) #:mutable)

(define (lexer-err err)
  (add-err! (format "lexer error: ~s; x = ~a, y = ~a\n" err x y)))

(define (add l ttype content)
  (set-lexer-tokens! l
                     (cons
                      (token ttype content
                             (car (lexer-span l))
                             (list-ref (lexer-span l) 1))
                      (lexer-tokens l))))

(define (take-str l ttype left right p p-err err)
  (define content "")
  (begin
    (for/last ([c (string->list (substring
                                 (lexer-code l)
                                 (+ (lexer-idx l) right)))]
               [i (in-naturals)])
      #:break (p c)
      (cond
        [(p-err c) (begin
                     (lexer-err err)
                     (when repl?
                       (begin
                         (set-lexer-code!
                          (substring (lexer-code l) (+ i (lexer-idx l) right)))
                         (set-lexer-idx! (+ (lexer-idx l) 1)))))]
        [else (set! content (concat-strings content (string c)))]))
    (set-lexer-code! l (substring (lexer-code l)
                                (+ (string-length content) (lexer-idx l) left)))
    (set-lexer-idx! l -1)
    (add l ttype content)
    (set-lexer-span! l
     (list (- (+
               (string-length content)
               (car (lexer-span l))) 1)
           (list-ref (lexer-span l) 1)))))

(define (take l ttype p p-err)
  (take-str l ttype 0 0 p p-err '()))

(define (scan #:code [code '()] #:lexer [l '()])
  (if (null? l)
      (scan #:lexer (lexer '() code '(0 0) 0))
      (begin
        (when (null? l) (error "lexer is nil"))
        (for ([i (in-naturals)])
          #:break (>= (lexer-idx l) (string-length (lexer-code l)))
          (begin
            (define c (string-ref (lexer-code l) (lexer-idx l)))
            (cond
              [(or (char=? c #\() (char=? c #\[))
               (add l 'lparen "(")]
              [(or (char=? c #\)) (char=? c #\]))
               (add l 'rparen ")")]
              [(or (char=? c #\return) (char=? c #\tab) (char=? c #\space))
               (void)]
              [(char=? c #\newline)
               (set-lexer-span! l '(-1 1))]
              [(and
                (char=? c #\-)
                (char=? (string-ref (lexer-code l) (+ (lexer-idx l) 1)) #\-))
               (for ([i (in-naturals)])
                 #:break (char=? (string-ref (lexer-code l) (+ i 1)) #\newline)
                 (set! i (+ i 1)))]
              [(char=? c #\-)
               (take l
                     (lambda (c) (or
                                  (not (char-alphabetic? c))
                                  (string-contains? symbols (string c))
                                  (char-numeric? c)))
                     (lambda (c) #f))]
              [(char=? c #\")
               (take-str l 'string 2 1
                         (lambda (c) (char=? c #\"))
                         (lambda (c) (char=? c #\newline))
                         "dangling \"")]
              [(or (char-alphabetic? c) (string-contains? symbols (string c)))
               (take l 'name
                     (lambda (c) (or
                                  (not (char-alphabetic? c))
                                  (string-contains? symbols (string c))
                                  (char-numeric? c)))
                     (lambda (c) #f))]
              [(char-numeric? c)
               (take l 'num
                     (lambda (c) (not (char-numeric? c)))
                     (lambda (c) #f))])
            (set-lexer-span! l
             (list (+ (car (lexer-span l)) 1) (list-ref (lexer-span l) 1)))
            (set-lexer-idx! l (+ (lexer-idx l) 1))))
        (reverse (lexer-tokens l)))))