#include <stdio.h>

// la funcion utiliza un arreglo para llevar registro de los valores de i-21, i-14, i-7 en cada iteracion entre 22 y n
// las posiciones se van rotando para no tener un arreglo tan grande, solo nos interesan hasta i-21
long long funIter(long long n) {
    if(n <= 21) return n;

    long long valores[22]; // usamos un arreglo de 22 posiciones, (del 0 al 21) para acumular los valores

    for(long long i = 0; i <= 21; i++){ // se llena la primera instancia de los valores
      valores[i] = i;
    }

    for(int i = 22; i <= n; i++){
        valores[i % 22] = valores[(i-7) % 22] + valores[(i-14) % 22] + valores[(i-21) % 22];
    }
    return valores[n % 22];
}

int main() {
    printf("%lld\n", funIter(27));
    return 0;
}
