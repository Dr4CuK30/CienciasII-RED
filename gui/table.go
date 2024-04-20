package gui

import (
	"CienciasII-RED/utils"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)

func PrintSendResults(results *[]*utils.PackageResultInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Mensaje", "Tiempo", "Estimado", "Ruta", "Estado"})

	// Crear una goroutine para actualizar la tabla periódicamente
	go func() {
		for {
			table.ClearRows()
			for _, d := range *results {
				table.Append([]string{fmt.Sprintf("%d", d.Id), d.Message, fmt.Sprintf("%d", d.Shortest), fmt.Sprintf("%d", d.Estimated), d.Path, d.Status})
			}
			table.Render()
			time.Sleep(time.Second)
		}
	}()
	select {}
}
