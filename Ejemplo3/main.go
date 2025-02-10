package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

// Structs a utilizar
type Entrada struct {
	Text string `json:"text"`
}

type StatusResponse struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func main() {
	//Registrar la ruta y asociarla con la función
	http.HandleFunc("/analizar", getCadenaAnalizar)
	c := cors.Default() //configurar CORS con la politica predeterminada

	/*
		CORS: Es para la comunicacion entre dominios (frontend - backend)
	*/
	handler := c.Handler(http.DefaultServeMux) //manejador HTTP con CORS

	//Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}

func getCadenaAnalizar(w http.ResponseWriter, r *http.Request) {
	//el apuntador r es para pasar la direccion de memoria y evitar hacer
	//duplicas innecesarias de la estructura hhtp.Request ya que dicha estructura es muy grande
	//En resumen: para optimizar el uso de memoria y rendimiento

	// Configurar la cabecera de respuesta
	w.Header().Set("Content-type", "application/json")

	//Para responder al cliente que esta haciendo la solicitud
	var status StatusResponse

	//Verificar que sea un POST
	if r.Method == http.MethodPost {
		var entrada Entrada //struct que contendrá los datos enviados desde el cliente

		//intenta decodificar el cuerpo del json y almacernaro en la variable entrada (estructura)
		if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
			//si ocurre un error, responder con un error 400
			http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
			//mensaje personalizado del error
			status = StatusResponse{Message: "Error al decodificar JSON", Type: "unsucces"}
			//envíar el JSON de la respuesta de error
			json.NewEncoder(w).Encode(status)
			return
		}

		//creo un lector de bufer para el archivo
		lector := bufio.NewScanner(strings.NewReader(entrada.Text))
		//leer la entrada linea por linea
		for lector.Scan() {
			linea := lector.Text()
			analizar(linea)
		}

		//fmt.Println("Cadena recibida ", entrada.Text)
		w.WriteHeader(http.StatusOK)

		status = StatusResponse{Message: "recibido correctamente", Type: "succes"}
		json.NewEncoder(w).Encode(status)
	} else {
		//http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		status = StatusResponse{Message: "Metodo no permitido", Type: "unsucces"}
		json.NewEncoder(w).Encode(status)
	}
}

func analizar(entrada string) {
	if entrada == "fdisk" {
		fmt.Println("fdisk")

	} else if entrada == "comando2" {
		fmt.Println("comando 2")

	} else if entrada == "exit" {
		fmt.Println("Buenas tardes")
		os.Exit(0)
	}

}
