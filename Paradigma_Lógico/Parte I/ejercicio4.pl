append([],Ys,Ys).
append([X|Xs], Ys, [X|Zs]) :- append(Xs,Ys,Zs).

reverso([],[]).
reverso([X|Xs],Res):-
    reverso(Xs,XsReverso),
    append(XsReverso,[X],Res).

soniguales([M|N],[P|Q]) :- M=P, soniguales(N,Q).
soniguales(M,N) :- M=[], N=[].

invertir(L, R):-
    reverso(R, Res),
    soniguales(L, Res).
