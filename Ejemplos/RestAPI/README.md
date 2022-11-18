# RestAPI con GorillaMux

Aqui encontraras un ejemplo sobre como crear una restAPI con la libreria GorillaMux.
En dicho ejemplo encontraremos las herramientas basicas para la realizacion de la misma, tales como:

- Creacion de endpoints
- Solucion a problemas de CORS
- Lectura de json recibido en el body de un POST
- Lectura de datos por medio de endpoint
- Conversion de structs a formato json
- Respuestas




Para importar las librerias necesarias utilizaremos los comandos:

```console
go get -u github.com/gorilla/mux
go get -u github.com/gorilla/handlers
```

Si necesitas mas informacion sobre la libreria y sus distintas funciones visita el 
[repositorio oficial de gorilla](https://github.com/gorilla/mux)
