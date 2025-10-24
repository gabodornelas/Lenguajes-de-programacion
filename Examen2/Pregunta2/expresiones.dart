import 'dart:io';

bool SumaResta(String lado){
    for(int i=lado.length-1; i>0 ;i--){
        if(lado[i] == '+' || lado[i] == '-'){
            return true;
        }
    }
    return false;
}

int calcular(int op1, String operator, int op2, bool post){
    switch (operator) {
        case '+':
            return op1 + op2;
        case '-':
            if(post){
                return op2 - op1;
            }else{
                return op1 - op2;
            }
        case '*':
            return op1 * op2;
        case '/':
            return op1 ~/ op2; 
        default:
            throw ArgumentError('Operador desconocido: $operator');
    }
}


(int, int, String) Evaluar(int ini, List<String> tokens, bool post){
    int resultado = 0, izquierda, derecha;
    String izq, der, expresion = '', operador;
    if(tokens.length - ini - 1 >= 2){
        if(tokens[ini] == '+' || tokens[ini] == '-' || tokens[ini] == '*' || tokens[ini] == '/'){ // es un operador
            operador = tokens[ini];
            if(tokens[ini+1] == '+' || tokens[ini+1] == '-' || tokens[ini+1] == '*' || tokens[ini+1] == '/'){ // es un segundo operador
                (ini, izquierda, izq) = Evaluar(ini+1,tokens,post);
            }else{
                if(int.tryParse(tokens[ini+1]) != null){ // es un numero
                    izquierda = int.parse(tokens[ini+1]);
                    izq = tokens[ini+1];
                }else{
                    throw Exception('Error de sintaxis: El token no es operador ni numero.');
                }
            }
            if(tokens[ini+2] == '+' || tokens[ini+2] == '-' || tokens[ini+2] == '*' || tokens[ini+2] == '/'){ // es otro operador
                (ini, derecha, der) = Evaluar(ini+2,tokens,post);
            }else if(int.tryParse(tokens[ini+2]) != null){ // es otro numero
                derecha = int.parse(tokens[ini+2]);
                der = tokens[ini+2];
            }else{
                throw Exception('Error de sintaxis: El token no es operador ni numero.');
            }
            resultado = calcular(izquierda, operador, derecha, post);
            if(operador  == '*' || operador == '/'){
                if(SumaResta(izq)){
                    if(post){
                        izq = ')' + izq + '(';
                    }else{
                        izq = '(' + izq + ')';
                    }
                }
                if(SumaResta(der)){
                    if(post){
                        der = ')' + der + '(';
                    }else{
                        der = '(' + der + ')';
                    }
                }
                
            }
            expresion = izq + ' ' + operador + ' ' + der;
            return (ini+1, resultado, expresion);
        }else{
            throw Exception('Error de sintaxis: El token no es un operador');
        }
    }else{
        throw Exception('Error de sintaxis: Faltan argumentos.');
    }
    return (0,0,'');
}

void main() {
    while (true) {
        stdout.write('--- Manejador de Expresiones Aritméticas ---\n\nLas expresiones deben venir con espacios entre numeros y entre operandos\n\nEVAL <orden> <expr>\nMOSTRAR <orden> <expr>\nSALIR\n\n');
        final input = stdin.readLineSync();
        if (input == null) continue;
        // convertir a lista el input
        final tokens = input.trim().split(RegExp(r'\s+')).where((s) => s.isNotEmpty).toList();

        if (tokens.isEmpty) continue;

        final accion = tokens[0].toUpperCase();

        if (accion == 'SALIR') {
        print('Saliendo...');
        break;
        }

        if (tokens.length < 3) {
        print('Error: Comando incompleto. Formato: <accion> <orden> <expr>');
        continue;
        }

        final orden = tokens[1].toUpperCase();
        List<String> expresiones = tokens.sublist(2);

        if (orden != 'PRE' && orden != 'POST') {
        print('Error: Orden inválido. Use PRE (pre-fijo) o POST (post-fijo).');
        continue;
        }

        try {
        if(accion == 'EVAL' || accion == 'MOSTRAR') {
            int result, ini;
            String infix;
            bool post = false;
            if(orden == 'POST'){
                expresiones = expresiones.reversed.toList(); // se revierte para tratarlo como un pre
                post = true;
            }
            (ini, result, infix) = Evaluar(0, expresiones, post);
            if(accion == 'EVAL'){
                print('$result\n');
            }else{
                if(orden == 'POST'){
                    infix = infix.split('').reversed.join(); // se revierte para imprimirlo bien
                }
                print('$infix\n');
            }
        } else {
            print('Error: Acción inválida. Use EVAL, MOSTRAR o SALIR.');
        }
        } catch (e) {
            print('Eror de procesamiento: $e');
        }
    }
}
