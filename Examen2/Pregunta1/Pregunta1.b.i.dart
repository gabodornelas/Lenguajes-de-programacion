import 'dart:io';

void main() {
    stdout.write('Ingresa un entero: ');
    // la entrada es String?
    String? entradaEntero = stdin.readLineSync(); 
    // intentamos convertir la cadena a entero
    int? num = int.tryParse(entradaEntero ?? '');
  
    if (num != null && num > 0) {// los numeros negativos entrarian en bucle infinito, se omiten
      int n = num; // hay que usar variable auxiliar porque dart no esta seguro de si num es null
      int contador = 0;
        while(n != 1){
            if(n % 2 == 0){  // caso par
                n = n ~/ 2;  // division entera
            }else{  // caso impar
                n = 3 * n + 1;
            }
            contador++;
        }
        print('\n$contador');
    } else {
        print('\nError: La entrada no es un número entero válido.');
    }
}
