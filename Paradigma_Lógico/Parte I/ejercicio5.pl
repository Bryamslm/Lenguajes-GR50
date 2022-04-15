divide([],_,_,_).

divide([X|Xs],Umb, [X|Mrs], Mys) :-
    X<Umb,
    divide(Xs,Umb, Mrs, Mys).

divide([X|Xs],Umb, Mrs, [X|Mys]) :-
    X>=Umb,
    divide(Xs,Umb, Mrs, Mys).

partir(Lista, Umbral, Menores, Mayores):-
    divide(Lista, Umbral, Menores, Mayores).

