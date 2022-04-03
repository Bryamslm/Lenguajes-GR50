#include <stdio.h>
#include <malloc.h>
#include <iostream>

using namespace std;

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
     * PARA INSERTAR UN ELEMENTO EN EL ARREGLO PRIMERO RECORRE DE FORMAS ASCENDENTE EN EL PRIMER WHILE
     * HASTA LLEGAR A LA POSICION DONDE QUIERO INSERTAR CON UN CONTADOR, EL VALOR QUE SE REEMPLAZA ES GUARDADO
     * COMO RESPALDO EN LA VARIANLE VALOR Y LUEGO SE RECORRE LA LISTA DE FORMA DESCENDENTE PARA IR MOVEINDO
     * LOS ELEMENTOS EN LA POSICION DONDE ESTAN +1 PARA ASI CONCLUIR.
     */
    int* end_ptr = source_ptr+size;
    int* aux_ptr = source_ptr;
    int cont = 0;
    source_ptr = static_cast<int *>(realloc(source_ptr, size + 1));
    while(aux_ptr < end_ptr) {
        if(cont == pos){
            int valor=*(source_ptr+pos);//el número que fue reemplazdo y debe ser colocado en pos+1
            *(source_ptr+pos) = ele;
            while(end_ptr > source_ptr+pos) {
                if(end_ptr == source_ptr+pos+1){
                   *end_ptr=valor;
                    break;
                }
                *end_ptr = *(end_ptr-1);
                end_ptr--;
            }
            break;

        }
        cont++;
        aux_ptr++;
    }
    //*(source_ptr+pos) = ele;
    // implementar ciclo para insertar en una posición específica que contemple mover todos los elementos posteriores... DEBE INCLUIR EL PARÁMETRO pos
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

    new_size = insert_array(arreglo2,100,SIZE_ARRAY, 1);//INDICE 1
    print_array(arreglo2,new_size);


    return 0;
}