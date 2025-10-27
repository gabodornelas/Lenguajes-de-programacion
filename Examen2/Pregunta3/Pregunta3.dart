//----------------------------------------------------------------------------------------------------------------------------------
//                                      NOTA
// hice 2 funciones porque ninguna de las 2 me terminaban de gustar y estuve intentando lograr un algoritmo mas limpio y optimo, 
// pero esto fue lo mejor que logre :)
//----------------------------------------------------------------------------------------------------------------------------------

// nmenorNumero busca el numero menor, lo remueve de la lista y lo devuelve con yield, llama recursivamente con un elemento menos
// recibe la lista y un numero menor generico que en este caso usamos el entero mas grande posible
Iterable<int> menorNumero(List<int> lista, menor) sync* {
  if(!lista.isEmpty){  // si la lista es vacia no hace nada
    int j=0,i;
    for (i = 0; i < lista.length;i++){
      if(lista[i] < menor){
        menor = lista[i];
        j = i; //indice del menor valor
      }
    }
    yield lista.removeAt(j);    
    for(final x in menorNumero(lista, 9223372036854775807)){
        yield x;
    }
  }
}

// mergeSortIterable es la implementacion del algoritmo mergesort pero con yield
Iterable<int> mergeSortIterable(List<int> lista) sync* {
  if (lista.length == 1) {
    yield lista[0];      // caso base, la lista tiene un elemento
  } else {
    int mid = lista.length ~/ 2;
    // convierte a lista todos los elementos que devuelve la funcion con la sublista izquierda
    List<int> izq = mergeSortIterable(lista.sublist(0, mid)).toList();
     // convierte a lista todos los elementos que devuelve la funcion con la sublista derecha
    List<int> der = mergeSortIterable(lista.sublist(mid)).toList();

    int i = 0, j = 0;
    // me da el orden en que se devuelven los valores
    while (i < izq.length && j < der.length) {
      if (izq[i] <= der[j]) {
        yield izq[i++];
      } else {
        yield der[j++];
      }
    }
    
    while (i < izq.length) { // se acabo der, aun quedan valores en izq
      yield izq[i++];
    }

    while (j < der.length) { // se acabo izq, aun quedan valores en der
      yield der[j++];
    }
  }
}

void main() {
  List<int> numeros = [1, 3, 3, 2, 1, 5, 0];
  
  for (final elemento in menorNumero(numeros,9223372036854775807)) {
    print(elemento);
  }
  for (final elemento in mergeSortIterable(numeros)) {
    print(elemento);
  }
}
