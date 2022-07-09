(define true
    (lambda (f s) (^ f)))

(define false
    (lambda (f s) (^ s)))

(define if
    (lambda (c e et) (^ (c e et))))

(define and
    (lambda (x y) (if x y false)))

(define or
    (lambda (x y) (if x true (if y true false))))

(println (if true 1 2))

(define fun
	(lambda (n) (println n)))

(fun 1 2)

(println (1 2 3 ("hello" 1)))

(println "hello")