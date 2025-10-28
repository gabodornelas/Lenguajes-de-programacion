#include <stdio.h>

// la funcion auxiliar es recursiva de cola, hace las llamadas recursivas con un estilo similar al de la funcion iterativa
long long funAux(long long n, long long i, long long valores[22]) {
    if (i > n) return valores[n % 22];
    valores[i % 22] = valores[(i - 7) % 22] + valores[(i - 14) % 22] + valores[(i - 21) % 22];
    return funAux(n, i + 1, valores);
}

long long funCola(long long n) {
    if (n <= 21) return n;
    long long valores[22];
    for(long long i = 0; i <= 21; i++){ // se llena la primera instancia de los valores
        valores[i] = i;
    }
    return funAux(n, 22, valores);
}

int main(){
  printf("%lld\n",funCola(27));
  return 0;
}
