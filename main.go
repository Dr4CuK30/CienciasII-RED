package main

import (
	"CienciasII-RED/model"
)

func main() {
	InitNetwork()
}

func InitNetwork() {
	// Inicializa la red (el grafo) y se le define el peso maximo de paquete por envio
	network := model.Network{MaxPackageLength: 5}

	// Creacion de nodos
	network.AddRouter("A")
	network.AddRouter("B")
	network.AddRouter("C")
	network.AddRouter("D")
	network.AddRouter("E")
	network.AddRouter("F")

	// Conexion nodos
	network.ConnectRouters("A", "B", 8)  // Los pesos se actualizaran por envio de paquete
	network.ConnectRouters("A", "C", 4)  // en este caso se setean para comprobar funcionalidad
	network.ConnectRouters("A", "D", 14) // del calculo del camino mas corto
	network.ConnectRouters("B", "D", 2)
	network.ConnectRouters("C", "D", 3)
	network.ConnectRouters("D", "E", 5)

	// Prueba de camino corto
	network.CalculateShortestPath("A", "C")

	// Mensaje end to end
	network.SendMessage("A", "E", "Lamo demasiado novia")

}
