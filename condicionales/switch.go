package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go está corriendo en ")
	so := runtime.GOOS

	if so == "darwin" {
		fmt.Println("MacOS")
	} else if so == "linux" {
		fmt.Println("Linux")
	} else {
		fmt.Println("Windows")
	}

	//Con switch
	switch so {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Windows")
	}

	//también se puede declarar dentro del switch la variable
	switch so2 := runtime.GOOS; so2 {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Windows")
	}

	//Break está declarado de forma automático no es
	//necesario colocarlo
	switch so2 := "darwin"; so2 {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Windows")
	}

	//Si queremos ejecutar varios case
	//se usa fallthrough
	switch so2 := "darwin"; so2 {
	case "darwin":
		fmt.Println("MacOS")
		fallthrough
	case "linux":
		fmt.Println("Linux")
		fallthrough
	default:
		fmt.Println("Windows")
	}

	//Para usar switch (true)
	hora := time.Now().Hour()
	switch {
	case hora < 12:
		fmt.Println("Buenos días!")
	case hora < 17:
		fmt.Println("Buenas tardes")
	default:
		fmt.Println("Buenas noches")
	}

	switch true {
	case hora < 12:
		fmt.Println("Buenos días!")
	case hora < 17:
		fmt.Println("Buenas tardes")
	default:
		fmt.Println("Buenas noches")
	}
}
