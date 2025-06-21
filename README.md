# Code Challenges in Go 🚀

## 📖 Descripción
Este proyecto reúne múltiples desafíos técnicos implementados en **Go**, encapsulados en endpoints de una API REST utilizando el framework **Gin Gonic**. Cada desafío aborda problemas técnicos comunes y los resuelve con prácticas eficientes y seguras. Además, el proyecto está diseñado para ser fácilmente extensible y probado.

---

## 💡 ¿Qué incluye este proyecto?

1. **Operaciones de Pila (Stack)**: Implementa una pila con concurrencia segura para agregar y eliminar elementos.
2. **Conversión de Números Romanos a Enteros**: Convierte números romanos a su representación decimal.
3. **API de Ejemplo**: Un endpoint básico que devuelve un mensaje sencillo.
4. **Suma de Múltiplos**: Calcula la suma de múltiplos de 3 o 5 por debajo de un límite dado.
5. **Conteo de Pares**: Cuenta cuántos pares en un arreglo suman un valor objetivo.
6. **Verificación de Anagramas**: Verifica si dos palabras son anagramas.
7. **Ruta más Corta (Matriz)**: Devuelve un placeholder para procesamiento de rutas en matrices.
8. **Identificación de Únicos y Duplicados**: Separa números únicos y duplicados en una lista.
9. **Simulación de Concurrencia**: Ejecuta tareas concurrentes usando goroutines y `sync.WaitGroup`.
10. **Tiempo de Respuesta de URLs**: Mide el tiempo de respuesta de múltiples URLs usando `libcurl`.
11. **Conteo de Caracteres**: Devuelve la frecuencia de cada carácter en una cadena.

---

## 🧪 Nuevos Retos Agregados

12. **Fork Reader** (`POST /fork-reader`):  
    Divide un flujo de texto alternando byte por byte en dos salidas separadas (`w1` y `w2`).

13. **Double Server** (`POST /double-server`):  
    Recibe un arreglo de números, los envía a un servidor interno vía canales y responde con cada número multiplicado por dos.

14. **Pair Sum (Two-pointer)** (`POST /pair-sum`):  
    Dados dos arreglos (ascendente y descendente) y un número objetivo, cuenta cuántos pares (uno de cada arreglo) suman exactamente dicho valor.

15. **Channel Pair Sum** (`POST /channel-pair-sum`):  
    Similar al anterior pero la entrada es simulada mediante canales. Usa goroutines y sincronización para procesar pares que suman un objetivo.

16. **Browser Navigator** (`POST /browser-navigator`):  
    Simula navegación con comandos como `hopTo`, `backtrack`, `leapForward`, manteniendo historial hacia atrás y hacia adelante.

---

## 📦 Colección para Bruno API Client

Incluye una carpeta `collections/` con archivos `.bru` compatibles con [Bruno API Client](https://www.usebruno.com/), permitiendo probar todos los endpoints fácilmente.

Para usarla:
1. Abre Bruno.
2. Importa la carpeta `collections/`.
3. Ejecuta cada request directamente desde Bruno.

---

## 🛠️ Requisitos Previos

1. **Go**:
   - Descarga e instala Go desde [https://go.dev/dl/](https://go.dev/dl/).
   - Verifica la instalación:
     ```bash
     go version
     ```

2. **Gin Gonic**:
   - El framework se instalará automáticamente al ejecutar el proyecto gracias a `go mod tidy`.

---

## 🚀 Cómo Ejecutar el Proyecto

1. **Clona este repositorio**:
   ```bash
   git clone https://github.com/giojimen3z/test_for-job.git
   cd test_for-job

2. **Instala dependencias**:
   ```bash
   go mod tidy 
   
3. **Ejecuta el servidor**: 
    ```bash
   go run main.go
4. **Accede a la API en http://localhost:8080**