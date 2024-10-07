package plot

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"fmt"
)

// PlotArray gera um gráfico de barras com base no array fornecido
func PlotArray(arr []int, iteration int) {
	p := plot.New()
	p.Title.Text = "Quicksort - Iteração " + fmt.Sprintf("%d", iteration)
	p.Y.Label.Text = "Valores"
	p.X.Label.Text = "Índices"

	values := make(plotter.Values, len(arr))
	for i, v := range arr {
		values[i] = float64(v)
	}

	bar, err := plotter.NewBarChart(values, vg.Points(20))
	if err != nil {
		fmt.Println("Erro ao criar gráfico de barras:", err)
		panic(err)
	}

	bar.LineStyle.Width = vg.Length(0)
	bar.Color = color.RGBA{R: 0, B: 100, A: 0}

	p.Add(bar)


	if err := p.Save(6*vg.Inch, 4*vg.Inch, fmt.Sprintf("graphics/quicksort_iter_%05d.png", iteration)); err != nil {
		fmt.Println("Erro ao salvar gráfico:", err)
		panic(err)
	}
}