package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PostConMap(w http.ResponseWriter, r *http.Request) {
	// la siguiente linea se utiliza para leer el contenido que recibimos por medio del body
	// (comunmente un json)
	// y lo guardaremos en una variable body de tipo []byte
	body, _ := io.ReadAll(r.Body) // en versiones de go < 1.16 utilizamos ioutil en ligar de io

	// creamos el mapa donde guardaremos la informacion
	var data map[string]interface{}
	// la libreria json con su funcion Unmarshal nos permiten leer el contenido y transformar
	// la informacion en un mapa o en un struct
	json.Unmarshal(body, &data)

	// por ultimo vamos a recorrer el mapa y almacenaremos los datos en un string para regresar dicho
	// string como respuesta
	response := ""
	for key, val := range data {
		response += key + " : " + fmt.Sprintf("%v", val) + "\n"
	}

	fmt.Fprintln(w, response)
}
