#include <stdio.h>
#include "_cgo_export.h"

void SayHello(){

    printf("%d", add(1,2));
    puts("say hello");
}

int* gsort(int *arr, int len) {
    for (int i = 0; i < len; i++) {
        printf("%d", arr[i]);
    }
    arr[0] = 1;
    return arr;
}