largo([],0).
largo([_|T],N) :-
          largo(T,N1),
          N is 1+N1.
largopar(L):-
          largo(L,N),
          0 is mod(N,2).
largoimpar(L):-
          largo(L,N),
          1 is mod(N,2).
