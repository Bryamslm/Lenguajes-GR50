#include <stdio.h>
#include <malloc.h>

const int SIZE_ARRAY = 10;
void print_array(int *arreglo,int size){
    int* end_ptr = arreglo+size;
    while(arreglo < end_ptr) {
        printf("%i ",*arreglo);
        arreglo++;
    }
    printf("\n");
}
void copy_array(int *source_ptr, int *dest_ptr, int size){
    int* end_ptr = source_ptr+size;
    while(source_ptr < end_ptr) {
        *dest_ptr = *source_ptr;
        source_ptr++;
        dest_ptr++;
    }
}
int insert_array(int* source_ptr, int ele, int size, int pos){
    /*
     * Se mueven los valores una posicion a la derecha hasta llegar a la posicion pos
     * leugo solo se inserta el nuevo elemento.
     */
    source_ptr = (int*)(realloc(source_ptr, size + 1));
    int* end_ptr = source_ptr+size;
    while(end_ptr > source_ptr+pos) {//mueve los valores una posicion a la derecha hasta pos.(de atrás a adelante)
        if(end_ptr == source_ptr+pos+1){
            *end_ptr = *(end_ptr-1);
            break;
        }
        *end_ptr = *(end_ptr-1);
        end_ptr--;
    }
    *(source_ptr+pos) = ele;//inserta el nuevo elemento
    return size+1;
}
int main() {
    int arreglo[SIZE_ARRAY];
    int *arreglo2 = (int*) malloc(SIZE_ARRAY*sizeof(int)) ;
    arreglo[0]=9; arreglo[1]=8; arreglo[2]=7; arreglo[3]=6; arreglo[4]=5;
    arreglo[5]=4; arreglo[6]=3; arreglo[7]=2; arreglo[8]=1; arreglo[9]=0;
    print_array(arreglo,SIZE_ARRAY);
    copy_array(arreglo,arreglo2,SIZE_ARRAY);
    int new_size = 0;

    new_size = insert_array(arreglo2,100,SIZE_ARRAY, 5);//INDICE 1
    print_array(arreglo2,new_size);

    new_size = insert_array(arreglo2,35,new_size, 0);//INDICE 1
    print_array(arreglo2,new_size);

    new_size = insert_array(arreglo2,99,new_size, 12);//INDICE 1
    print_array(arreglo2,new_size);

    return 0;
}
