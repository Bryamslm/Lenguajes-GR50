#|
3) Construya la funcion filtrar. Esta funcion recibe una lista y retorna otra lista con los mismos elementos de la primera, pero eliminando los elementos repetidos. Ejemplos:

> (filtrar '())
()

> (filtrar '(1 3 5 7 8 8 8 9 2 3))
(1 5 7 8 9 2 3)

> (filtrar '(a b c d))
(a b c d)
|#

(define (member? element lista)

  (cond
    ((null? lista)
     #f)
    ((equal? element (car lista))
     #t
     )
    (else
     (member? element (cdr lista))
     )
    )

  )

(define (filtrar lista)

  (cond
    ((null? lista)
     '()
     )
    ((member? (car lista) (cdr lista))
     (filtrar (cdr lista))
     )
    (else
     (cons (car lista) (filtrar (cdr lista)))
     )
    )

  )