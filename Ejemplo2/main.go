package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//MENSAJES DE INICIO
	Ms_inicio := "Bienvenido escriba un comando"
	Ms_info := "Si desesa salir escriba el comando exit"

	fmt.Println(Ms_inicio)
	fmt.Println(Ms_info)

	//bufio.NewScanner(os.Stdin) -> crea un nuevo escáner que lee desde os.Stdin (entrada de teclado)
	//El escaner se usa para leer linea por linea (newreader lee varias lineas o entradas grandes)
	//reader -> variable que almacena el escáner
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\n$: ")
		reader.Scan() //Lee la entrada y la prepara para ser usada mas adelante.

		entrada := reader.Text() //guarda la entrada (preparada por Scan) como string

		//salir deteniendo el ciclo for sin el os.Exit(0)
		//if entrada == "exit" {
		//	fmt.Println("Buenas tardes")
		//	break
		//}

		//Llamada a metodo analizar
		analizar(entrada)
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
