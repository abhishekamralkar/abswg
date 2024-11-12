package main

import "fmt"

func main() {
	paperPrices := []float32{1.00, 2.00, 4.00, 8.00, 16.00, 32.00, 64.00, 128.00}
	inflatedPaperPrices := make([]float32, len(paperPrices))

	for i, v := range paperPrices {
		inflatedPaperPrices[i] = (v * 0.08) + v
	}

	fmt.Println(inflatedPaperPrices)
}
