package main

import (
	Comandos "Proyecto/Comandos"
	DM "Proyecto/Comandos/AdministradorDiscos" //DM -> DiskManagement (Administrador de discos)

	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func analizar(entrada string) {
	//Separar los parametros -size=3000 -path=ruta (obtenemos la lista: size=3000, path=ruta)
	parametros := strings.Split(entrada, " -")

	//analizamos los parametros
	if strings.ToLower(parametros[0]) == "execute" {
		if len(parametros) == 2 {
			tmpParametro := strings.Split(parametros[1], "=")
			if strings.ToLower(tmpParametro[0]) == "path" && len(tmpParametro) == 2 {
				//abrir el archivo
				archivo, err := os.Open(tmpParametro[1])
				if err != nil {
					fmt.Println("Error al leer el script: ", err)
					return
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
		//--------------------------------------- OTROS ------------------------------------------------------------
	} else if strings.ToLower(parametros[0]) == "fdisk" {
		//FDISK
		if len(parametros) > 1 {
			DM.Fdisk(parametros)
		} else {
			fmt.Println("FDISK ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "rep" {
		//REP
		if len(parametros) > 1 {
			//Comandos.Rep(parametros)
			Comandos.Rep()
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
}
