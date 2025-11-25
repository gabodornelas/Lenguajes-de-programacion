#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "manejador.h"

//---------------------------------------------------------------------------------------------------------------
//--------------FUNCIONES DE LISTAS------------------------------------------------------------------------------



struct Tipos* crearNodoS(struct Simple* tipo) {
    struct Tipos* nuevo = (struct Tipos*)malloc(sizeof( struct Tipos));
    if (!nuevo) {
        printf("Error al reservar memoria\n");
        exit(1);
    }
    nuevo->atomico = tipo;
    nuevo->sig = NULL;
    return nuevo;
}

void insertarS(struct Tipos** cabeza, struct Simple* tipo) {
    struct Tipos* nuevo = crearNodoS(tipo);
    if (*cabeza == NULL) {
        *cabeza = nuevo;
        return;
    }
    struct Tipos* actual = *cabeza;
    while (actual->sig != NULL) {
        actual = actual->sig;
    }
    actual->sig = nuevo;
}

void liberarNodoS(struct Tipos* cabeza, int op) {
    struct Tipos* temp;
    while (cabeza != NULL) {
        temp = cabeza;
        cabeza = cabeza->sig;
        if(op){
            free(temp->atomico);
        }
        free(temp);
    }
}

struct Construcciones* crearNodoC(struct Compuesto* tipo) {
    struct Construcciones* nuevo = (struct Construcciones*)malloc(sizeof(struct Construcciones));
    if (!nuevo) {
        printf("Error al reservar memoria\n");
        exit(1);
    }
    nuevo->construido = tipo;
    nuevo->sig = NULL;
    return nuevo;
}

void insertarC(struct Construcciones** cabeza, struct Compuesto* tipo) {
    struct Construcciones* nuevo = crearNodoC(tipo);
    if (*cabeza == NULL) {
        *cabeza = nuevo;
        return;
    }
    struct Construcciones* actual = *cabeza;
    while (actual->sig != NULL) {
        actual = actual->sig;
    }
    actual->sig = nuevo;
}

void liberarNodoC(struct Construcciones* cabeza) {
    struct Construcciones* temp;
    while (cabeza != NULL) {
        temp = cabeza;
        cabeza = cabeza->sig;
        liberarNodoS(temp->construido->campos,0);
        free(temp->construido); 
        free(temp);
    }
}



//---------------------------------------------------------------------------------------------------------------
//--------------INICIALIZAR ESTRUCTURAS--------------------------------------------------------------------------



struct Simple* crearTipoS(char *nom, int rep, int al) {
    struct Simple* nuevo = (struct Simple*)malloc(sizeof(struct Simple));
    if (!nuevo) {
        printf("Error al reservar memoria\n");
        exit(1);
    }
    strncpy(nuevo->nombre, nom, sizeof(nuevo->nombre)-1);
    nuevo->nombre[sizeof(nuevo->nombre)-1] = '\0'; // asegurar terminación
    nuevo->representacion = rep;
    nuevo->alineacion = al;
    return nuevo;
}

struct Compuesto* crearTipoC(char *nom, struct Tipos* cam, int al, int tam, int rep, int cant) {
    struct Compuesto* nuevo = (struct Compuesto*)malloc(sizeof(struct Compuesto));
    if (!nuevo) {
        printf("Error al reservar memoria\n");
        exit(1);
    }
    strncpy(nuevo->nombre, nom, sizeof(nuevo->nombre)-1);
    nuevo->nombre[sizeof(nuevo->nombre)-1] = '\0';
    nuevo->alineacion = al;
    nuevo->tamMax = tam;
    nuevo->repMax = rep;
    nuevo->cantBloques = cant;
    nuevo->campos = cam;
    return nuevo;
}



//---------------------------------------------------------------------------------------------------------------
//--------------BUSCAR EN ESTRUCTURAS----------------------------------------------------------------------------



struct Compuesto* existeCompuesto(char *nombre, struct Construcciones* complejos){
    struct Construcciones* auxiliar = complejos;
    while(auxiliar != NULL){
        if (strcmp(auxiliar->construido->nombre,nombre) == 0){
            return auxiliar->construido;
        }
        auxiliar = auxiliar->sig;
    }
    return NULL;
}

struct Simple* existeSimple(char *nombre, struct Tipos* atomicos){
    struct Tipos* auxiliar = atomicos;
    while(auxiliar != NULL){
        if (strcmp(auxiliar->atomico->nombre,nombre) == 0){
            return auxiliar->atomico;
        }
        auxiliar = auxiliar->sig;
    }
    return NULL;
} 



//---------------------------------------------------------------------------------------------------------------
//--------------PROCESAMIENTOS-----------------------------------------------------------------------------------



struct Tipos* procesar_atomico(struct Construcciones* complejos, struct Construcciones* uniones, struct Tipos* atomicos , char *linea){
    int representacion, alineacion;
    char nombre[10];
    if (sscanf(linea, "%9s %d %d", nombre, &representacion, &alineacion) == 3){
        if(alineacion >= representacion){
            if(existeSimple(nombre, atomicos) == NULL){
                if(existeCompuesto(nombre, complejos) == NULL && existeCompuesto(nombre, uniones) == NULL){
                    struct Simple* nuevo = crearTipoS(nombre, representacion, alineacion);
                    insertarS(&atomicos,nuevo);
                    printf("✅ Se proceso el tipo Atomico con exito\n");
                    return atomicos;
                }else{
                    printf("❌ Error: El nombre %s ya existe como tipo compuesto.\n", nombre);
                }
            }else{
                printf("❌ Error: El tipo simple %s ya existe.\n", nombre);
            }
        }else{
            printf("❌ Error: La alineacion no puede ser menor que la representacion de un tipo.\n");
        }
    }else{
        printf("❌ Error: Formato ATOMICO incorrecto. Uso: ATOMICO <nombre> <representacion> <alineacion>\n");
    }
    return atomicos;
}

struct Construcciones* procesar_compuesto(struct Construcciones* complejos, struct Construcciones* uniones, struct Tipos * atomicos, char *linea){
    int alineacion = 0, tam = 0, cant = 0, espEnBloque, rep = 0;
    char nombre[10];
    char simples[10];
    struct Tipos* TipoSimples = NULL;
    if (sscanf(linea, "%9s", nombre) == 1) {
        //buscar nombre en complejos y uniones
        if(existeCompuesto(nombre, complejos) == NULL && existeCompuesto(nombre, uniones) == NULL ){
            //buscar nombre en atomicos
            if(existeSimple(nombre, atomicos) == NULL){
                linea += strcspn(linea, " \t\n");   // saltar palabra
                linea += strspn(linea, " \t\n");    // saltar espacios
                //leemos los tipos simples
                while (*linea != '\0') {
                    if (sscanf(linea, "%9s", simples) == 1) {
                        //revisa los nombres simples
                        if(existeCompuesto(simples, complejos) == NULL && existeCompuesto(simples, uniones) == NULL && strcmp(simples,nombre) != 0){
                            struct Simple* nombreSimple = existeSimple(simples, atomicos);
                            if(nombreSimple != NULL){
                                insertarS(&TipoSimples,nombreSimple);
                                //asegura la alineacion del struct/union y la cantidad de bloques
                                if(nombreSimple->alineacion >= alineacion){ //hay una nueva alineacion
                                    alineacion = nombreSimple->alineacion;
                                    espEnBloque = alineacion;
                                    cant += 1;
                                }else{
                                    if(nombreSimple->alineacion >= espEnBloque || offset(alineacion,alineacion - espEnBloque,nombreSimple->alineacion)){
                                        //no se puede almacenar en el bloque actual, ya sea porque no hay espacio, o dentro del espacio no puede iniciar en multiplo de su alineacion correspondiente
                                        cant += 1;
                                        espEnBloque = alineacion - nombreSimple->alineacion;
                                    }else{  //se asigna en el bloque
                                        espEnBloque -= nombreSimple->alineacion;
                                    }
                                }
                                tam += nombreSimple->representacion;
                                if(nombreSimple->representacion > rep){ //asigno la representacion maxima, util para uniones
                                    rep = nombreSimple->representacion;
                                }
                                // avanzar ptr hasta después del nombre leído
                                linea += strcspn(linea, " \t\n");   // saltar palabra
                                linea += strspn(linea, " \t\n");    // saltar espacios
                            }else{
                                printf("❌ Error: Intentas crear un tipo compuesto con un tipo simple que no existe.\n");
                                liberarNodoS(TipoSimples, 0);
                                return complejos;
                            }
                        }else{
                            printf("❌ Error: Intentas crear un tipo compuesto con otro tipo compuesto.\n");
                            liberarNodoS(TipoSimples, 0);
                            return complejos;
                        }
                    } else {
                        break;
                    }
                }
                if(TipoSimples != NULL){    //creacion del tipo Compuesto
                    struct Compuesto* nuevo = crearTipoC(nombre, TipoSimples, alineacion, tam, rep, cant + 1);
                    insertarC(&complejos, nuevo);
                    printf("✅ Se proceso el tipo Compuesto con exito\n");
                }else{
                    printf("❌ Error: Intentas crear un tipo compuesto sin tipos simples.\n");
                }
            }else{
                printf("❌ Error: El nombre %s ya existe como tipo simple.\n", nombre);
            }
        }else{
            printf("❌ Error: El tipo compuesto %s ya existe.\n", nombre);
        }
    }else{
        printf("❌ Error: Formato STRUCT/UNION incorrecto. Uso: STRUCT/UNION <nombre> [<tipo>]\n");
    }
    return complejos;
}



//---------------------------------------------------------------------------------------------------------------
//--------------DESCRIPCIONES------------------------------------------------------------------------------------



int offset(int alineacion, int espacio, int i){
    for(int j = espacio; j <= alineacion; j++){
        if(j % i == 0 && alineacion - j >= i){
            return 1;
        }
    }
    return 0;
}

int optimo(struct Compuesto* nombreEstructura){
    struct Tipos* aux = nombreEstructura->campos;
    size_t tamanio = nombreEstructura->alineacion + 1;
    // Asignar e inicializar a cero usando calloc
    int *ordenes = (int*)calloc(tamanio, sizeof(int));
    int al = nombreEstructura->alineacion, cantBloques = 0, espEnBloque = al;
    while(aux != NULL){
        ordenes[aux->atomico->alineacion] += 1;
        aux = aux->sig;
    }
    for (int i = al; i >= 1; i--) {
        while(ordenes[i] > 0){
            if(espEnBloque - i >= 0 && offset(al, al - espEnBloque, i)){ // si cabe en el bloque
                espEnBloque -= i;
            }else{  //no cabe en el bloque
                cantBloques += 1;
                espEnBloque = al - i;
            }
            ordenes[i] -= 1;
        }
    }
    cantBloques += 1;
    free(ordenes);
    return cantBloques * al;
}

void describir(struct Construcciones* complejos, struct Construcciones* uniones, struct Tipos* atomicos, char *linea){
    char nombre[10];
    if (sscanf(linea, "%9s", nombre) != 1) {
        printf("❌ Error: Formato DESCRIBIR incorrecto. Uso: DESCRIBIR <nombre>\n");
        return;
    }
    struct Compuesto* nombreEstructura = existeCompuesto(nombre, complejos);
    int tamOptimo;
    if(nombreEstructura != NULL){

        printf("Sobre el tipo %s:\n", nombre);
        printf("Es un tipo compuesto, un STRUCT\n");
    
        printf("\tSin empaquetar:\n");
        printf("\t\tTamaño: %d bytes\n", nombreEstructura->alineacion * nombreEstructura->cantBloques);
        printf("\t\tAlineación: %d bytes\n", nombreEstructura->alineacion);
        printf("\t\tBytes Desperdiciados: %d bytes \n\n", (nombreEstructura->alineacion * nombreEstructura->cantBloques) - nombreEstructura->tamMax);

        printf("\tCon empaquetado:\n");
        printf("\t\tTamaño: %d bytes\n", nombreEstructura->tamMax);
        printf("\t\tAlineación: 1 bytes\n");
        printf("\t\tBytes Desperdiciados: 0 bytes\n");

        tamOptimo = optimo(nombreEstructura);
        printf("\tReordenamiento Optimo:\n");
        printf("\t\tTamaño: %d bytes\n", tamOptimo);
        printf("\t\tAlineación: %d bytes\n",nombreEstructura->alineacion);
        printf("\t\tBytes Desperdiciados: %d bytes\n", tamOptimo - nombreEstructura->tamMax);

    }else{
        struct Compuesto* nombreUnion = existeCompuesto(nombre, uniones);
        if(nombreUnion != NULL){
            printf("Sobre el tipo %s:\n", nombre);
            printf("Es un tipo compuesto, un UNION\n");
        
            printf("\tSin empaquetar:\n");
            printf("\t\tTamaño: %d bytes\n", nombreUnion->repMax);
            printf("\t\tAlineación: %d bytes\n", nombreUnion->alineacion);
            printf("\t\tBytes Desperdiciados: %d bytes \n\n", nombreUnion->alineacion - nombreUnion->repMax);

            printf("\tCon empaquetado:\n");
            printf("\t\tTamaño: %d bytes\n", nombreEstructura->repMax);
            printf("\t\tAlineación: 1 bytes\n");
            printf("\t\tBytes Desperdiciados: 0 bytes\n");

            printf("\tReordenamiento Optimo:\n");
            printf("\t\tTamaño: %d bytes\n", nombreUnion->repMax);
            printf("\t\tAlineación: %d bytes\n", nombreUnion->alineacion);
            printf("\t\tBytes Desperdiciados: %d bytes \n\n", nombreUnion->alineacion - nombreUnion->repMax);

        }else{
            struct Simple* nombreSimple = existeSimple(nombre, atomicos);
            if(nombreSimple != NULL){

                printf("Sobre el tipo %s:\n", nombre);
                printf("Es un tipo simple, por lo tanto esta informacion no es muy relevante\n");

                printf("\tSin empaquetar:\n");
                printf("\t\tTamaño: %d bytes\n", nombreSimple->representacion);
                printf("\t\tAlineación: %d bytes\n", nombreSimple->alineacion);
                printf("\t\tBytes Desperdiciados: 0 bytes \n\n");

                printf("\tCon empaquetado:\n");
                printf("\t\tTamaño: %d bytes\n", nombreSimple->representacion);
                printf("\t\tAlineación: 1 bytes\n");
                printf("\t\tBytes Desperdiciados: 0 bytes\n");

                printf("\tReordenamiento Optimo:\n");
                printf("\t\tTamaño: %d bytes\n", nombreSimple->representacion);
                printf("\t\tAlineación: %d bytes\n", nombreSimple->alineacion);
                printf("\t\tBytes Desperdiciados: 0 bytes\n");

            }else{
                printf("Error: El tipo %s no existe.\n", nombre);
            }
        }
    }
}



//---------------------------------------------------------------------------------------------------------------
//--------------MAIN---------------------------------------------------------------------------------------------


int main_manejador() {

    char buffer[256];
    char comando[10];

    struct Tipos* atomicos = NULL;
    struct Construcciones* estructuras = NULL;
    struct Construcciones* uniones = NULL;

    // Inicio del bucle
    while (1) {
        //  MENU  
        printf("===========================================================\n");
        printf("       Simulador de Manejador de Tipos de Datos\n");
        printf("===========================================================\n");
        printf("Comandos: ATOMICO <nombre> <representacion> <alineacion>\n");
        printf("          STRUCT <nombre> [<tipo>]\n");
        printf("          UNION <nombre> [<tipo>]\n");
        printf("          DESCRIBIR <nombre>\n");
        printf("          SALIR\n");
        printf("-----------------------------------------------------------\n");

        printf("\n> ");
        if (fgets(buffer, sizeof(buffer), stdin) == NULL) {
            break;
        }

        if (sscanf(buffer, "%s", comando) != 1) {
            continue;
        }
        
        // Determinar el resto de la línea después del comando
        size_t len = strlen(comando);
        char *resto_linea = buffer + len;
        while (*resto_linea == ' ' || *resto_linea == '\t') {
            resto_linea++;
        }

        if (strcmp(comando, "ATOMICO") == 0) {
            atomicos = procesar_atomico(estructuras, uniones, atomicos, resto_linea);
        } else if (strcmp(comando, "STRUCT") == 0) {
            estructuras  = procesar_compuesto(estructuras, uniones, atomicos, resto_linea);
        } else if (strcmp(comando, "UNION") == 0) {
            uniones = procesar_compuesto(uniones, estructuras, atomicos, resto_linea);  //uniones y estructuras va al reves para que se agregue la nueva union en uniones
        } else if (strcmp(comando, "DESCRIBIR") == 0) {
            describir(estructuras, uniones, atomicos, resto_linea);
        } else if (strcmp(comando, "SALIR") == 0) {
            liberarNodoC(estructuras);
            liberarNodoC(uniones);
            liberarNodoS(atomicos,1);
            break;
        } else {
            printf("Comando no reconocido.\n");
        }

    }
    printf("\nSimulador terminado.\n");
    return 0;
}
