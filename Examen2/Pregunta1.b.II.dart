List<int> mergeSort(List<int> lista) {
  // caso base, la lista tiene 0 o 1 elemento, ya está ordenada.
  if (lista.length <= 1) {
    return lista;
  }

  // encontrar el punto medio con division entera
  int mitad = lista.length ~/ 2; 

  // dividir la lista en dos mitades
  List<int> izq = lista.sublist(0, mitad);
  List<int> der = lista.sublist(mitad);

  // llamadas recursivas para ordenar cada mitad
  List<int> izqOrdenada = mergeSort(izq);
  List<int> derOrdenada = mergeSort(der);

  // combinar las dos mitades ordenadas
  return merge(izqOrdenada, derOrdenada);
}

List<int> merge(List<int> izq, List<int> der) {
  List<int> resultado = [];
  int i = 0; // indice para la lista izquierda
  int j = 0; // indice para la lista derecha

  // mientras haya elementos en ambas listas
  while (i < izq.length && j < der.length) {  // compara el menor elemento en cada sublista
    if (izq[i] < der[j]) {
      resultado.add(izq[i]);
      i++;
    } else {
      resultado.add(der[j]);
      j++;
    }
  }

  // añadir los elementos restantes de la lista izquierda
  while (i < izq.length) {
    resultado.add(izq[i]);
    i++;
  }

  // añadir los elementos restantes de la lista derecha
  while (j < der.length) {
    resultado.add(der[j]);
    j++;
  }

  return resultado;
}

void main() {
  List<int> listaDesordenada = [8, 3, 1, 6, 9, 2, 5, 4, 7];
  print('Lista original: $listaDesordenada');

  List<int> listaOrdenada = mergeSort(listaDesordenada);
  print('Lista ordenada: $listaOrdenada');
}
