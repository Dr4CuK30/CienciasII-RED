package utils

// PackageResultInfo Tipo que contiene los resultados de un envio de paquete
type PackageResultInfo struct {
	Id        int
	Message   string
	Status    string
	Estimated int
	PathData
}

// PathsTable Tipo auxiliar utilizado para algoritmo dijkstra
type PathsTable map[string]*PathData

// PathData Tipo auxiliar que almacena la información del camino más corto en el algoritmo dijkstra
type PathData struct {
	Locked   bool
	Shortest int
	Path     string
}
