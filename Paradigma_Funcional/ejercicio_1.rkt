#|
1) Construya una funcion que se llame vector-inverso. Esta función recibe dos argumentos: un numero N y un vector con valores entre 0 y N.
   Debe producir un vector donde cada entrada ha sido restada al valor de N. Ejemplos:

>(vector-inverso 10 '())
()

>(vector-inverso 10 '(1 2 3))
(9 8 7)

>(vector-inverso 10 '(3 9 7 0))
(7 1 3 10)

>(vector-inverso 20 '(3 9 7 0))
(17 11 13 20)
|#

(define (vector-inverso N vector)
  (cond
    ((null? vector)
     '()
     )
    (else
     (cons (- N (car vector)) (vector-inverso N (cdr vector)))
     )

   )

 )