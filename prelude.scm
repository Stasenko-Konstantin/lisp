--- prelude ---

(define true
  (lambda (f s) (^ f)))

(define false
  (lambda (f s) (^ s)))

(define if
        (lambda (c e1 e2) (^ (c e1 e2))))

(define and
        (lambda (x y) (if x y false)))

(define or
        (lambda (x y) (if x true (if y true false))))

(define f
  (lambda (n) (
                (define g (lambda (m) (^ m)))
                (^ g n))))

(define cons
        (lambda (x y)
          ((define dispatch
             (lambda (m) (^ (m x y))))
            (^ dispatch))))