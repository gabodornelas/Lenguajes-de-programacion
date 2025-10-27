/// Devuelve un Iterable (que es un iterador) con los elementos de la lista
/// en orden de menor a mayor.
Iterable<int> obtenerOrdenado(List<int> lista) sync* {
  // 1. Crea una copia de la lista y la ordena
  // El método .sort() modifica la lista original, por eso usamos .toList()
  // para trabajar con una copia.
  var listaOrdenada = lista.toList();
  listaOrdenada.sort();

  // 2. Itera sobre la lista ordenada y "produce" cada elemento.
  for (final elemento in listaOrdenada) {
    yield elemento; // La palabra clave 'yield' devuelve el siguiente valor.
  }
}

// Ejemplo de Uso:
void main() {
  final numeros = [1, 3, 3, 2, 1, 5, 0];
  
  // Llamamos a la función para obtener el generador (Iterable)
  final iteradorOrdenado = obtenerOrdenado(numeros);

  print('Lista original: $numeros'); // [1, 3, 3, 2, 1, 5, 0]
  print('Elementos ordenados:');

  // Iteramos sobre el generador para obtener los valores
  for (final n in iteradorOrdenado) {
    print(n);
  }
  
  // Salida:
  // 0
  // 1
  // 1
  // 2
  // 3
  // 3
  // 5
}
