#ifndef MANEJADOR_H
#define MANEJADOR_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

// -----------------------------------------------------------------------------
// ESTRUCTURAS
// -----------------------------------------------------------------------------

struct Simple{
    char nombre[10];
    int representacion;
    int alineacion;
};

struct Tipos{    // Lista de Tipos simples
    struct Simple *atomico;
    struct Tipos *sig;
};

struct Compuesto{
    char nombre[10];
    int alineacion;         // la alineacion mas grande de algun campo 
    int tamMax;             // la suma de las representaciones de todos los campos
    int repMax;             // la representacion mas grande de algun campo
    int cantBloques;        // la cantidad de bloques
    struct Tipos *campos;
};

struct Construcciones{ // Lista de Tipos compuestos
    struct Compuesto *construido;
    struct Construcciones *sig;
};

// -----------------------------------------------------------------------------
// DECLARACIÓN DE FUNCIONES
// -----------------------------------------------------------------------------

// Funciones de Listas
struct Tipos* crearNodoS(struct Simple* tipo);
void insertarS(struct Tipos** cabeza, struct Simple* tipo);
void liberarNodoS(struct Tipos* cabeza, int op);

struct Construcciones* crearNodoC(struct Compuesto* tipo);
void insertarC(struct Construcciones** cabeza, struct Compuesto* tipo);
void liberarNodoC(struct Construcciones* cabeza);

// Inicializar Estructuras
struct Simple* crearTipoS(char *nom, int rep, int al);
struct Compuesto* crearTipoC(char *nom, struct Tipos* cam, int al, int tam, int rep, int cant);

// Buscar en Estructuras
struct Compuesto* existeCompuesto(char *nombre, struct Construcciones* complejos);
struct Simple* existeSimple(char *nombre, struct Tipos* atomicos);

// Procesamientos (Lógica Principal)
struct Tipos* procesar_atomico(struct Construcciones* complejos, struct Construcciones* uniones, struct Tipos* atomicos , char *linea);
struct Construcciones* procesar_compuesto(struct Construcciones* complejos, struct Construcciones* uniones, struct Tipos * atomicos, char *linea);

// Descripciones y Cálculos
int offset(int alineacion, int espacio, int i);
int optimo(struct Compuesto* nombreEstructura);
void describir(struct Construcciones* complejos, struct Construcciones* uniones, struct Tipos* atomicos, char *linea);


//MAIN
int main_manejador();

#endif // MANEJADOR_H
