#|
2) Construya la funcion recursiva primer-numero. Esta funcion recibe una lista y devuelve el primer elemento de esa lista que es un número.
   Debe retornar el número si lo encuentra, sino retornar null. Ejemplos:

> (primer-numero '(a g y 5 7 w v 6))
5

> (primer-numero '(a g y w v))
()

|#

(define (primer-numero lista)
  (cond
    ((null? lista)
     '()
     )
    ((number? (car lista))
     (car lista)
     )
    (else
     (primer-numero (cdr lista))
     )

   )
  )