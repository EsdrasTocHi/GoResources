package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// definimos un struct y al lado le colocamos el nombre que tendra dicho parametro en nuestros json
// tanto para lectura como para la escritura
type Prueba struct {
	Param1 int    `json:"param1"`
	Param2 string `json:"param2"`
}

func PostConStruct(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	// Podemos observar que la lectura es igual con estructuras que con mapas
	var data Prueba
	err := json.Unmarshal(body, &data)

	if err != nil {
		// si llegamos a detectar un error es porque el contenido leido desde el body no concuerda
		// con la definicion de nuestro struct
		fmt.Fprintln(w, "error en la lectura del contenido")
		return
	}

	// en esta ocasion podemos utilizar la estructura para hacernos mas sencillo el manejo de los datos
	// los siguientes datos se imprimiran en la consola desde la cual corramos el main
	fmt.Println(data.Param1)
	fmt.Println(data.Param2)

	// para efectos de este ejemplo, retornaremos el struct en formato json
	// esto haciendo uso de la libreria json y su funcion Marshall
	response, _ := json.Marshal(data)

	fmt.Fprintln(w, string(response))
}
