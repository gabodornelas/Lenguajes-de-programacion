// explicacion de nmenorNumero
Iterable<int> menorNumero(List<int> lista, menor) sync* {
  if(!lista.isEmpty){
    for (final elemento in lista){
      if(elemento < menor){
        menor = elemento;
      }
    }
    yield menor;
    lista.remove(menor);
    for(final x in menorNumero(lista, 9223372036854775807)){
        yield x;
    }
  }
}

void main() {
  List<int> numeros = [1, 3, 3, 2, 1, 5, 0];
  
  for (final elemento in menorNumero(numeros,9223372036854775807)) {
    print(elemento);
  }
}
