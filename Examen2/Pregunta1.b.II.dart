// función recursiva que trabaja con una porción de la lista definida por índices
void mergeSort(List<int> lista, List<int> aux, int low, int high) {
  // caso base, la porción tiene 0 o 1 elemento
  if (low >= high) {
    return;
  }

  // encontrar el punto medio
  int mid = low + (high - low) ~/ 2;

  // llamadas recursivas para ordenar cada mitad
  mergeSort(lista, aux, low, mid);        // llama al lado izquierdo
  mergeSort(lista, aux, mid + 1, high);   // llama al lado derecho

  // combinar las dos mitades ordenadas en la lista
  merge(lista, aux, low, mid, high);
}

// función merge que usa el array auxiliar para el ordenamiento
void merge(List<int> lista, List<int> aux, int low, int mid, int high) {
  // copiar la porción de la lista al array aux
  for (int k = low; k <= high; k++) {
    aux[k] = lista[k];
  }

  int i = low;      // puntero para la mitad izquierda en aux
  int j = mid + 1;  // puntero para la mitad derecha en aux
  
  // combinar desde aux y reemplazar en lista
  for (int k = low; k <= high; k++) {
    if (i > mid) {
      // la mitad izquierda se agotó, tomar de la derecha
      lista[k] = aux[j++];
    } else if (j > high) {
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
  
  // llamar a la función recursiva con los índices iniciales (low=0, high=length-1)
  mergeSort(listaDesordenada, aux, 0, listaDesordenada.length - 1);
  
  print('Lista ordenada: $listaDesordenada'); 
}
