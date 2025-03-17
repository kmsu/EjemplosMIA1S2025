package main

import (
	Comandos "Proyecto/Comandos"
	DM "Proyecto/Comandos/AdministradorDiscos" //DM -> DiskManagement (Administrador de discos)
	FS "Proyecto/Comandos/SistemaDeArchivos"   //FS -> FileSystem (sistema de archivos)

	"bufio"
	"fmt"
	"os"
	"strings"
	//instalar go get -u github.com/rs/cors en la raiz del proyecto
)

type Entrada struct {
	Text string `json:"text"`
}

type StatusResponse struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

/*func main() {
	//EndPoint
	http.HandleFunc("/analizar", getCadenaAnalizar)

	// Configurar CORS con opciones predeterminadas
	//Permisos para enviar y recir informacion
	c := cors.Default()

	// Configurar el manejador HTTP con CORS
	handler := c.Handler(http.DefaultServeMux)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", handler)

}*/

func main() {
	// MENSAJES DE INICIO
	Ms_inicio := "Bienvenido escriba un comando..."
	Ms_info := "(si desea salir escriba el comando: exit)"
	fmt.Println(Ms_inicio)
	fmt.Println(Ms_info)
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n$: ")
		reader.Scan()

		entrada := strings.TrimRight(reader.Text(), " ") //Quitar espacios vacios a la derecha
		linea := strings.Split(entrada, "#")             //para ignorar comentarios desde la consola manual
		//entrada := execute -path=script.txt
		if strings.ToLower(linea[0]) != "exit" {
			analizar(linea[0])
		} else {
			fmt.Println("Salida exitosa")
			break
		}
	}
}

/*
func getCadenaAnalizar(w http.ResponseWriter, r *http.Request) {
	var respuesta string
	// Configurar la cabecera de respuesta
	w.Header().Set("Content-Type", "application/json")

	var status StatusResponse
	//verificar que sea un post
	if r.Method == http.MethodPost {
		var entrada Entrada
		if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
			http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
			status = StatusResponse{Message: "Error al decodificar JSON", Type: "unsucces"}
			json.NewEncoder(w).Encode(status)
			return
		}

		//creo un lector de bufer para el archivo
		lector := bufio.NewScanner(strings.NewReader(entrada.Text))
		//leer el archivo linea por linea
		for lector.Scan() {
			//Elimina los saltos de linea
			if lector.Text() != "" {
				//Divido por # para ignorar todo lo que este a la derecha del mismo
				linea := strings.Split(lector.Text(), "#") //lector.Text() retorna la linea leida
				if len(linea[0]) != 0 {
					fmt.Println("\n*********************************************************************************************")
					fmt.Println("Comando en ejecucion: ", linea[0])
					respuesta += "***************************************************************************************************************************\n"
					respuesta += "Comando en ejecucion: " + linea[0] + "\n"
					respuesta += analizar(linea[0]) + "\n"
				}
				//Comentarios
				if len(linea) > 1 && linea[1] != "" {
					fmt.Println("#" + linea[1] + "\n")
					respuesta += "#" + linea[1] + "\n"
				}
			}

		}

		//fmt.Println("Cadena recibida ", entrada.Text)
		w.WriteHeader(http.StatusOK)

		status = StatusResponse{Message: respuesta, Type: "succes"}
		json.NewEncoder(w).Encode(status)

	} else {
		//http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		status = StatusResponse{Message: "Metodo no permitido", Type: "unsucces"}
		json.NewEncoder(w).Encode(status)
	}
}
*/

// func analizar(entrada string) string {
func analizar(entrada string) {
	//Separar los parametros -size=3000 -path=ruta (obtenemos la lista: size=3000, path=ruta)
	parametros := strings.Split(entrada, " -")

	//respuesta := ""

	//analizamos los parametros
	if strings.ToLower(parametros[0]) == "execute" {
		if len(parametros) == 2 {
			tmpParametro := strings.Split(parametros[1], "=")
			if strings.ToLower(tmpParametro[0]) == "path" && len(tmpParametro) == 2 {
				//abrir el archivo
				archivo, err := os.Open(tmpParametro[1])
				if err != nil {
					fmt.Println("Error al leer el script: ", err)
					//return
				}
				defer archivo.Close()
				//creo un lector de bufer para el archivo
				lector := bufio.NewScanner(archivo)
				//leer el archivo linea por linea
				for lector.Scan() {
					//Divido por # para ignorar todo lo que este a la derecha del mismo
					linea := strings.Split(lector.Text(), "#") //lector.Text() retorna la linea leida
					if len(linea[0]) != 0 {
						fmt.Println("\n*********************************************************************************************")
						fmt.Println("Linea en ejecucion: ", linea[0])
						analizar(linea[0])
					}
				}
			} else {
				fmt.Println("EXECUTE ERROR: parametro path no encontrado")
			}
		}

		//--------------------------------- ADMINISTRADOR DE DISCOS ------------------------------------------------
	} else if strings.ToLower(parametros[0]) == "mkdisk" {
		//MKDISK
		//crea un archivo binario que simula un disco con su respectivo MBR
		if len(parametros) > 1 {
			DM.Mkdisk(parametros)
		} else {
			fmt.Println("MKDISK ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "fdisk" {
		//FDISK
		if len(parametros) > 1 {
			DM.Fdisk(parametros)
		} else {
			fmt.Println("FDISK ERROR: parametros no encontrados")
		}
	} else if strings.ToLower(parametros[0]) == "mount" {
		//Mount
		if len(parametros) > 1 {
			DM.Mount(parametros)
		} else {
			fmt.Println("FDISK ERROR: parametros no encontrados")
		}

		//--------------------------------- SISTEMA DE ARCHIVOS ----------------------------------------------------
	} else if strings.ToLower(parametros[0]) == "mkfs" {
		//MKFS
		if len(parametros) > 1 {
			FS.Mkfs(parametros)
		} else {
			fmt.Println("MKFS ERROR: parametros no encontrados")
		}
		//--------------------------------------- OTROS ------------------------------------------------------------
	} else if strings.ToLower(parametros[0]) == "rep" {
		//REP
		if len(parametros) > 1 {
			Comandos.Rep(parametros)
			//Comandos.Rep()
		} else {
			fmt.Println("REP ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "exit" {
		fmt.Println("Salida exitosa")
		os.Exit(0)

	} else if strings.ToLower(parametros[0]) == "" {
		//para agregar lineas con cada enter sin tomarlo como error
	} else {
		fmt.Println("Comando no reconocible")
	}

	//return respuesta
}
