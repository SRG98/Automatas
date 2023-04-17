package main

import (
	"fmt"

	"github.com/SRG98/Automatas.git/controladores"
	"github.com/SRG98/Automatas.git/vista"
)

func main() {
	controlador := controladores.NewControladores()
	uinterface := vista.NewUI(controlador)
	err := uinterface.RunUI()
	if err != nil {
		fmt.Println("Error", err)
	}
}
