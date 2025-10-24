import 'package:test/test.dart';
import 'aritmetica.dart'; // Asegúrese de que la ruta sea correcta

void main() {
  // Pruebas para SumaResta
  group('SumaResta', () {
    test('Debe retornar true si contiene +', () {
      expect(SumaResta('5 + 3'), isTrue);
    });

    test('Debe retornar true si contiene -', () {
      expect(SumaResta('10 - 2'), isTrue);
    });

    test('Debe retornar false si no contiene + ni -', () {
      expect(SumaResta('4 * 5'), isFalse);
      expect(SumaResta('8 / 2'), isFalse);
      expect(SumaResta('123'), isFalse);
      expect(SumaResta(''), isFalse);
    });
  });

  // ----------------------------------------------------------------

  // Pruebas para calcular
  group('calcular', () {
    test('Suma (post=false)', () {
      expect(calcular(5, '+', 3, false), 8);
    });

    test('Resta (post=false, normal)', () {
      expect(calcular(10, '-', 4, false), 6);
    });
    
    test('Resta (post=true, op2 - op1)', () {
      // Debería ser op2 - op1, es decir, 4 - 10 = -6
      expect(calcular(10, '-', 4, true), -6); 
    });

    test('Multiplicación', () {
      expect(calcular(6, '*', 7, false), 42);
    });

    test('División (entera)', () {
      // 10 ~/ 3 = 3
      expect(calcular(10, '/', 3, false), 3); 
    });

    test('Lanzar error para operador desconocido', () {
      // El operador desconocido debe lanzar ArgumentError
      expect(() => calcular(1, '!', 2, false), throwsArgumentError);
    });
  });
  
  // ----------------------------------------------------------------

  // Pruebas para Evaluar (asume notación prefija)
  group('Evaluar (Prefija)', () {
    // Expresión: + 5 3 -> 5 + 3 = 8
    test('Suma simple', () {
      final tokens = ['+', '5', '3'];
      final resultado = Evaluar(0, tokens, false);
      expect(resultado.$2, 8); // El resultado es el segundo elemento del record
      expect(resultado.$3, '5 + 3'); // El infix string
    });
    
    // Expresión: * + 5 3 2 -> (5 + 3) * 2 = 16
    test('Expresión anidada sin paréntesis requeridos', () {
      final tokens = ['*', '+', '5', '3', '2'];
      final resultado = Evaluar(0, tokens, false);
      expect(resultado.$2, 16);
      expect(resultado.$3, '(5 + 3) * 2'); // Paréntesis agregados por * sobre +
    });
    
    // Expresión: + 2 * 5 3 -> 2 + 5 * 3 = 17
    test('Expresión anidada con operadores de mayor precedencia', () {
      final tokens = ['+', '2', '*', '5', '3'];
      final resultado = Evaluar(0, tokens, false);
      expect(resultado.$2, 17);
      expect(resultado.$3, '2 + 5 * 3'); // Sin paréntesis
    });

    // Expresión: - * 10 2 + 5 3 -> 10 * 2 - 5 + 3 = 20 - 8 = 12
    test('Expresión compleja con anidamiento a ambos lados', () {
      final tokens = ['-', '*', '10', '2', '+', '5', '3'];
      final resultado = Evaluar(0, tokens, false);
      expect(resultado.$2, 12);
      expect(resultado.$3, '10 * 2 - 5 + 3');
    });

    // Casos de error
    test('Error: Falta el tercer argumento', () {
      final tokens = ['+', '5'];
      expect(() => Evaluar(0, tokens, false), throwsA(isA<Exception>()));
    });
    
    test('Error: Token no es número ni operador (izquierda)', () {
      final tokens = ['+', 'cinco', '3'];
      expect(() => Evaluar(0, tokens, false), throwsA(isA<Exception>()));
    });
    
    test('Error: Token no es número ni operador (derecha)', () {
      final tokens = ['+', '5', 'tres'];
      expect(() => Evaluar(0, tokens, false), throwsA(isA<Exception>()));
    });
  });

  // ----------------------------------------------------------------

  // Pruebas para Evaluar (asume notación postfija)
  group('Evaluar (Postfija - lógica invertida)', () {
    // Expresión POST: 5 3 + -> (inverso: + 3 5) -> 3 + 5 = 8
    test('Suma simple (post)', () {
      // En el código principal, el token list se invierte antes de llamar a Evaluar
      final tokens = ['+', '3', '5']; 
      final resultado = Evaluar(0, tokens, true);
      expect(resultado.$2, 8); 
      // La expresión se construye invertida: 3 + 5
      expect(resultado.$3, '3 + 5'); 
    });
    
    // Expresión POST: 2 5 3 * + -> (inverso: + * 3 5 2) -> 3 * 5 + 2 = 17
    test('Expresión anidada con operadores de mayor precedencia (post)', () {
      final tokens = ['+', '*', '3', '5', '2'];
      final resultado = Evaluar(0, tokens, true);
      expect(resultado.$2, 17);
      // La expresión se construye invertida: 2 + (5 * 3).
      // Note el orden de resta en calcular: op2 - op1 si es 'POST'.
      // La reversión de la cadena para 'MOSTRAR' se hace en main().
      // Aquí probamos el resultado y la expresión generada antes de la reversión final.
      expect(resultado.$3, '3 * 5 + 2'); 
    });
  });
}