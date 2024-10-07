package plot

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// PlotArray gera um gráfico de barras com base no array fornecido
// compareIndices são os índices que devem ser destacados em verde
func PlotArray(arr []int, iteration int, name string, compareIndices []int) {
	p := plot.New()
	p.Title.Text = fmt.Sprintf("%s", name) + " : Iteração " + fmt.Sprintf("%d", iteration)
	p.Y.Label.Text = "Valores"
	p.X.Label.Text = "Índices"

	// Define as cores do gráfico
	p.Title.TextStyle.Color = color.White
	p.X.Label.TextStyle.Color = color.White
	p.Y.Label.TextStyle.Color = color.White
	p.X.Color = color.White
	p.Y.Color = color.White
	p.X.Tick.Label.Color = color.White
	p.Y.Tick.Label.Color = color.White

	p.BackgroundColor = color.Black

	// Prepara os valores para o gráfico de barras
	values := make(plotter.Values, len(arr))
	for i, v := range arr {
		values[i] = float64(v)
	}

	// Cria o gráfico de barras
	bar, err := plotter.NewBarChart(values, vg.Points(50))
	if err != nil {
		fmt.Println("Erro ao criar gráfico de barras:", err)
		panic(err)
	}

	// Ajusta as cores das barras
	bar.LineStyle.Width = vg.Length(10)
	for i := range values {
		if contains(compareIndices, i) {
			// bar.Color  = color.RGBA{R: 0, G: 255, B: 0, A: 255} // Verde para as barras em comparação
			bar.Color  = color.RGBA{R: 64, G: 64, B: 64, A: 255} // Cinza para as outras barras
		
		
		} else {
			bar.Color  = color.RGBA{R: 64, G: 64, B: 64, A: 255} // Cinza para as outras barras
		}

		
		// Criar uma nova barra para cada valor
	}
	
	p.Add(bar)

	// Salva o gráfico como uma imagem PNG
	if err := p.Save(20*vg.Inch, 10*vg.Inch, fmt.Sprintf("graphics/iter_%05d.png", iteration)); err != nil {
		fmt.Println("Erro ao salvar gráfico:", err)
		panic(err)
	}
}

// contains verifica se um índice está no slice de índices
func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
