miembro(X,[X|_]).
miembro(X,[_|Ys]) :- miembro(X,Ys).
verifica(_,[]).
verifica(L,[H|C]):-
         miembro(H,L),
         verifica(L,C).
subconj(S, S1):-
    verifica(S, S1).%%%S es el conjubto y S1 el conjunto a verificar

