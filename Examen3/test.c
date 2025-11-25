#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h> // Usaremos assert para las comprobaciones críticas
#include <stdbool.h>
#include "manejador.h" 

// Definición de las funciones de prueba
void test_lista_simples();
void test_lista_compuestos();
void test_existencia();
void test_procesar_atomico_exito();
void test_procesar_atomico_fallos();
void test_procesar_compuesto_struct_logica_simple();
void test_procesar_compuesto_struct_logica_padding();
void test_procesar_compuesto_union_logica();
void test_procesar_compuesto_fallos();
void test_optimo_struct();
void ejecutar_pruebas();

// Contador de pruebas y fallos
int pruebas_pasadas = 0;
int pruebas_falladas = 0;

// =========================================================================
// FUNCIONES DE UTILIDAD (Para hacer las pruebas más limpias)
// =========================================================================

void reportar(const char* nombre_prueba, bool resultado) {
    if (resultado) {
        printf("✅ PASA: %s\n", nombre_prueba);
        pruebas_pasadas++;
    } else {
        printf("❌ FALLA: %s\n", nombre_prueba);
        pruebas_falladas++;
    }
}

// =========================================================================
// PRUEBAS DE LISTAS Y UTILIDADES (Cubre insertar/buscar/liberar)
// =========================================================================

void test_lista_simples() {
    struct Tipos* atomicos = NULL;
    struct Simple* s1 = crearTipoS("int", 4, 4);
    
    // Prueba de inserción
    insertarS(&atomicos, s1);
    bool insert_ok = (atomicos != NULL && strcmp(atomicos->atomico->nombre, "int") == 0);
    reportar("LISTAS_1_InsertarSimple", insert_ok);

    // Prueba de búsqueda
    struct Simple* encontrado = existeSimple("int", atomicos);
    bool search_ok = (encontrado != NULL && encontrado->representacion == 4);
    reportar("LISTAS_2_BuscarSimple", search_ok);
    
    // Prueba de búsqueda fallida
    struct Simple* no_encontrado = existeSimple("float", atomicos);
    bool search_fail = (no_encontrado == NULL);
    reportar("LISTAS_3_BuscarSimpleFallo", search_fail);
    
    // El '1' indica que también se liberen los struct Simple
    liberarNodoS(atomicos, 1);
}

void test_existencia() {
    struct Simple* s1 = crearTipoS("Byte", 1, 1);
    struct Tipos* atomicos = crearNodoS(s1);

    // Creamos una lista Tipos que ya tiene s1, para pasársela al Compuesto
    struct Tipos* campos_c1 = crearNodoS(crearTipoS("Dummy", 1, 1));
    struct Compuesto* c1 = crearTipoC("S_Test", campos_c1, 4, 12, 4, 3);
    struct Construcciones* complejos = crearNodoC(c1);

    // Prueba de existeCompuesto
    struct Compuesto* encontrado_c = existeCompuesto("S_Test", complejos);
    bool search_c_ok = (encontrado_c != NULL && encontrado_c->alineacion == 4);
    reportar("EXIST_1_BuscarCompuesto", search_c_ok);

    // Prueba de existeSimple (que la lista original no se altere)
    struct Simple* encontrado_s = existeSimple("Byte", atomicos);
    bool search_s_ok = (encontrado_s != NULL && encontrado_s->alineacion == 1);
    reportar("EXIST_2_BuscarSimple", search_s_ok);

    // Liberar
    liberarNodoC(complejos); 
    liberarNodoS(atomicos, 1); // Liberamos los struct Simple creados para la lista atomicos
}

// =========================================================================
// PRUEBAS DE PROCESAMIENTO ATÓMICO (Cubre validaciones de entrada)
// =========================================================================

void test_procesar_atomico_exito() {
    struct Tipos* atomicos = NULL;
    struct Construcciones* complejos = NULL;
    struct Construcciones* uniones = NULL;
    
    // 1. Caso de éxito estándar (rep <= al)
    atomicos = procesar_atomico(complejos, uniones, atomicos, "Int 4 4\n");
    bool exito1 = (atomicos != NULL && existeSimple("Int", atomicos) != NULL);
    reportar("ATOMICO_1_Exito", exito1);
    
    // 2. Caso de éxito con alineación mayor (rep < al)
    atomicos = procesar_atomico(complejos, uniones, atomicos, "Short 2 4\n");
    bool exito2 = (existeSimple("Short", atomicos) != NULL);
    reportar("ATOMICO_2_ExitoAlineacionMayor", exito2);
    
    // Liberar memoria creada en la prueba
    liberarNodoS(atomicos, 1);
}

void test_procesar_atomico_fallos() {
    struct Tipos* atomicos = NULL;
    struct Construcciones* complejos = NULL;
    struct Construcciones* uniones = NULL;
    
    // Inicializar para pruebas de conflicto
    atomicos = procesar_atomico(complejos, uniones, atomicos, "Existente 4 4\n");
    
    // 1. Error: Alineación < Representación
    struct Tipos* res1 = procesar_atomico(complejos, uniones, atomicos, "Malo 4 2\n");
    bool fallo1 = (existeSimple("Malo", res1) == NULL); // No debe agregarse
    reportar("ATOMICO_3_FalloAlineacionMenor", fallo1);

    // 2. Error: Nombre ya existe como Simple
    struct Tipos* res2 = procesar_atomico(complejos, uniones, atomicos, "Existente 8 8\n");
    bool fallo2 = (existeSimple("Existente", res2)->representacion == 4); // No debe sobreescribir
    reportar("ATOMICO_4_FalloNombreExistente", fallo2);
    
    // 3. Error: Formato incorrecto (falta un número)
    struct Tipos* res3 = procesar_atomico(complejos, uniones, atomicos, "Incompleto 4\n");
    bool fallo3 = (existeSimple("Incompleto", res3) == NULL);
    reportar("ATOMICO_5_FalloFormato", fallo3);

    liberarNodoS(atomicos, 1);
}

// =========================================================================
// PRUEBAS DE PROCESAMIENTO COMPUESTO (STRUCT y UNION)
// =========================================================================

// Configuración inicial de tipos atómicos para compuestos
struct Tipos* setup_atomicos() {
    struct Tipos* atomicos = NULL;
    insertarS(&atomicos, crearTipoS("c1", 1, 1));
    insertarS(&atomicos, crearTipoS("i4", 4, 4));
    insertarS(&atomicos, crearTipoS("s2", 2, 2));
    insertarS(&atomicos, crearTipoS("d8", 8, 8));
    return atomicos;
}

void test_procesar_compuesto_struct_logica_simple() {
    struct Tipos* atomicos = setup_atomicos();
    struct Construcciones* estructuras = NULL;
    struct Construcciones* uniones = NULL;
    
    // 1. STRUCT con alineación simple (i4 c1 s2) 
    // tamMax (suma reps) = 7. Alineacion=4. 
    estructuras = procesar_compuesto(estructuras, uniones, atomicos, "S_Simple i4 c1 s2\n");
    struct Compuesto* s_simple = existeCompuesto("S_Simple", estructuras);
    
    bool al_ok = (s_simple != NULL && s_simple->alineacion == 4);
    bool tam_ok = (s_simple != NULL && s_simple->tamMax == 7);
    bool rep_ok = (s_simple != NULL && s_simple->repMax == 4);
    
    bool cant_ok = (s_simple != NULL && s_simple->cantBloques == 3); 
    
    reportar("STRUCT_1_ParametrosBase", al_ok && tam_ok && rep_ok);
    reportar("STRUCT_2_BloquesRotos", cant_ok);

    liberarNodoC(estructuras);
    liberarNodoS(atomicos, 1);
}

void test_procesar_compuesto_struct_logica_padding() {
    struct Tipos* atomicos = setup_atomicos();
    struct Construcciones* estructuras = NULL;
    struct Construcciones* uniones = NULL;

    // 2. STRUCT con relleno interno (c1 i4 c1)
    // tamMax = 6. Alineacion=4.
    estructuras = procesar_compuesto(estructuras, uniones, atomicos, "S_Padding c1 i4 c1\n");
    struct Compuesto* s_padding = existeCompuesto("S_Padding", estructuras);

    bool al_ok = (s_padding != NULL && s_padding->alineacion == 4);
    bool tam_ok = (s_padding != NULL && s_padding->tamMax == 6);

    bool cant_ok = (s_padding != NULL && s_padding->cantBloques == 4); 
    
    reportar("STRUCT_3_ConRellenoBloquesRotos", al_ok && tam_ok && cant_ok);

    liberarNodoC(estructuras);
    liberarNodoS(atomicos, 1);
}

void test_procesar_compuesto_union_logica() {
    struct Tipos* atomicos = setup_atomicos();
    struct Construcciones* estructuras = NULL;
    struct Construcciones* uniones = NULL;

    // UNION: Campos (c1, i4, s2). repMax=4. Alineacion=4.

    uniones = procesar_compuesto(uniones, estructuras, atomicos, "U_Test c1 i4 s2\n");
    struct Compuesto* u_test = existeCompuesto("U_Test", uniones);

    bool al_ok = (u_test != NULL && u_test->alineacion == 4); 
    bool rep_ok = (u_test != NULL && u_test->repMax == 4);   
    

    bool tam_roto = (u_test != NULL && u_test->tamMax == 7); 
    

    bool cant_roto = (u_test != NULL && u_test->cantBloques == 4);

    reportar("UNION_1_AlineacionRepMaxOK", al_ok && rep_ok);
    reportar("UNION_2_TamBloquesRotos", tam_roto && cant_roto); 

    liberarNodoC(uniones);
    liberarNodoS(atomicos, 1);
}

void test_procesar_compuesto_fallos() {
    struct Tipos* atomicos = setup_atomicos();
    struct Construcciones* estructuras = NULL;
    struct Construcciones* uniones = NULL;
    
    // 1. Error: Tipo simple no existe
    struct Construcciones* res1 = procesar_compuesto(estructuras, uniones, atomicos, "S_Fail1 noexiste i4\n");
    bool fallo1 = (existeCompuesto("S_Fail1", res1) == NULL);
    reportar("COM_FAIL_1_TipoInexistente", fallo1);

    // 2. Error: Nombre de compuesto ya existe
    estructuras = procesar_compuesto(estructuras, uniones, atomicos, "S_Existe c1\n");
    struct Construcciones* res2 = procesar_compuesto(estructuras, uniones, atomicos, "S_Existe i4\n");
    bool fallo2 = (existeCompuesto("S_Existe", res2)->tamMax == 1); // No debe cambiar
    reportar("COM_FAIL_2_NombreExistente", fallo2);

    // 3. Error: Usar nombre de atomico como compuesto
    struct Construcciones* res3 = procesar_compuesto(estructuras, uniones, atomicos, "i4 c1\n");
    bool fallo3 = (existeCompuesto("i4", res3) == NULL);
    reportar("COM_FAIL_3_ConflictoAtomico", fallo3);
    
    liberarNodoC(estructuras);
    liberarNodoS(atomicos, 1);
}

// =========================================================================
// PRUEBAS DE OPTIMIZACIÓN (Cubre la función optimo)
// =========================================================================

void test_optimo_struct() {

    
    // 1. Creamos el STRUCT (S_Simple: i4 c1 s2) para obtener los campos en desorden de alineación.
    struct Tipos* atomicos = setup_atomicos();
    struct Construcciones* estructuras = NULL;
    struct Construcciones* uniones = NULL;
    estructuras = procesar_compuesto(estructuras, uniones, atomicos, "S_Optimo c1 i4 s2\n");
    struct Compuesto* s_optimo = existeCompuesto("S_Optimo", estructuras);

    if (s_optimo == NULL) {
        reportar("OPTIMO_0_SetupFallo", false);
        liberarNodoC(estructuras);
        liberarNodoS(atomicos, 1);
        return;
    }

    // 2. Probamos la función optimo
    int resultado_optimo = optimo(s_optimo);

    bool optimo_ok = (resultado_optimo == 8); 
    reportar("OPTIMO_1_LogicaCorrecta (Esperado 8)", optimo_ok);

    
    liberarNodoC(estructuras);
    liberarNodoS(atomicos, 1);
}

// =========================================================================
// DRIVER DE PRUEBAS
// =========================================================================

void ejecutar_pruebas() {
    printf("======================================\n");
    printf(" INICIO DE PRUEBAS UNITARIAS\n");
    printf("======================================\n");

    printf("\n--- Pruebas de Listas y Existencia ---\n");
    test_lista_simples();
    test_existencia();

    printf("\n--- Pruebas ATOMICO ---\n");
    test_procesar_atomico_exito();
    test_procesar_atomico_fallos();

    printf("\n--- Pruebas COMPUESTO (STRUCT y UNION) ---\n");

    test_procesar_compuesto_struct_logica_simple();
    test_procesar_compuesto_struct_logica_padding();
    test_procesar_compuesto_union_logica();
    test_procesar_compuesto_fallos();

    printf("\n--- Pruebas de Reordenamiento Optimo ---\n");
    test_optimo_struct();

    printf("\n======================================\n");
    printf(" RESULTADO FINAL:\n");
    printf(" ✅ Pruebas Pasadas: %d\n", pruebas_pasadas);
    printf(" ❌ Pruebas Falladas: %d\n", pruebas_falladas);
    printf("======================================\n");

}

// =========================================================================
// MAIN DE PRUEBAS
// =========================================================================

int main() {
    ejecutar_pruebas();
    return 0;
}
