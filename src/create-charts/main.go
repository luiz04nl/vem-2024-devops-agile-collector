// package main

// import (
// 	"image/color"
// 	"log"

// 	"github.com/luiz04nl/devops-ic-collector/src/shared"
// 	_ "github.com/mattn/go-sqlite3"

// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/vg"
// )

// func createLineChart(xAxis []string, YAxis []float64, title string, valueText string, outFile string) {
// 	p := plot.New()
// 	p.Title.Text = title
// 	p.Y.Label.Text = valueText
// 	values := make(plotter.Values, len(YAxis))
// 	for i, val := range YAxis {
// 		values[i] = val
// 	}
// 	bar, err := plotter.NewBarChart(values, vg.Points(20))
// 	if err != nil {
// 		panic(err)
// 	}
// 	bar.LineStyle.Width = vg.Length(0)
// 	bar.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255} // cor azul
// 	p.Add(bar)
// 	p.NominalX(xAxis...)
// 	p.X.Label.Text = title

// 	if err := p.Save(10*vg.Inch, 10*vg.Inch, outFile); err != nil {
// 		panic(err)
// 	}
// }

// func createChartToDevOpsUsage() {
// 	var dataSourceName = "../../database/sqlite/repository-dataset.db"
// 	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
// 	if err != nil {
// 		log.Fatal("Was not possible connect with database:", err)
// 	}

// 	var aggregateDevOpsToolsUsageDto []shared.AggregateDevOpsToolsUsageDto

// 	aggregateDevOpsToolsUsageDto, err = newSQLiteRepository.AggregateDevOpsToolsUsage()
// 	if err != nil {
// 		log.Fatal("Error on get the repository:", err)
// 	}

// 	var ferramentas []string
// 	var utilizacao []float64

// 	for _, item := range aggregateDevOpsToolsUsageDto {
// 		ferramentas = append(ferramentas, item.Name)
// 		utilizacao = append(utilizacao, item.Value)
// 	}

// 	title := "Ferramenta de DevOps"
// 	valueText := "Utilização (%)"
// 	outFile := "../../out/charts/ferramentas-devops.png"

// 	createLineChart(ferramentas, utilizacao, title, valueText, outFile)
// }

// func main() {
// 	createChartToDevOpsUsage()
// }
