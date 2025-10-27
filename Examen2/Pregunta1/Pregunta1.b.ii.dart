// función recursiva que trabaja con una porción de la lista definida por índices
void mergeSort(List<int> lista, List<int> aux, int ini, int fin) {
  // caso base, la porción tiene 0 o 1 elemento
  if (ini >= fin) {
    return;
  }

  // encontrar el punto medio
  int med = ini + (fin - ini) ~/ 2;

  // llamadas recursivas para ordenar cada mitad
  mergeSort(lista, aux, ini, med);        // llama al lado izquierdo
  mergeSort(lista, aux, med + 1, fin);   // llama al lado derecho

  // combinar las dos mitades ordenadas en la lista
  merge(lista, aux, ini, med, fin);
}

// función merge que usa el array auxiliar para el ordenamiento
void merge(List<int> lista, List<int> aux, int ini, int med, int fin) {
  // copiar la porción de la lista al array aux
  for (int k = ini; k <= fin; k++) {
    aux[k] = lista[k];
  }

  int i = ini;      // puntero para la mitad izquierda en aux
  int j = med + 1;  // puntero para la mitad derecha en aux
  
  // combinar desde aux y reemplazar en lista
  for (int k = ini; k <= fin; k++) {
    if (i > med) {
      // la mitad izquierda se agotó, tomar de la derecha
      lista[k] = aux[j++];
    } else if (j > fin) {
      // la mitad derecha se agotó, tomar de la izquierda
      lista[k] = aux[i++];
    } else if (aux[i] <= aux[j]) {
      // el elemento de la izquierda es menor o igual
      lista[k] = aux[i++];
    } else {
      // el elemento de la derecha es menor
      lista[k] = aux[j++];
    }
  }
}

void main() {
  List<int> listaDesordenada = [8, 3, 1, 1, 9, 2, 5, 4, 7];
  print('Lista original: $listaDesordenada'); 
  List<int> aux = List<int>.from(listaDesordenada); 
  
  // llamar a la función recursiva con los índices iniciales (ini=0, fin=length-1)
  mergeSort(listaDesordenada, aux, 0, listaDesordenada.length - 1);
  
  print('Lista ordenada: $listaDesordenada'); 
}
