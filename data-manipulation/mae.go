package main

import (
	"fmt"
	"math"
)

/*
	Validacion de un modelo.
	Variables continuas (temperatura, porcentajes, cantidad de personas, desviacion, ganancias)
	Varibales discretas (estados: activo, inacativo, interrumpido, etc)
	Valores observados			valores pronosticados			margen de error
		10.5							8							3
		12.66							7.88						2.33
		09.2							6							1.76
		11.3							9.03						2
	MAE (Mean absolut error): error absoluto medio
	Podemos comparar el resultado con la media de los valores observados.
	formula:
		MAE = sum(abs(Vo - Vp)) / N
	Vo: Valor original
	Vp: Valor predicho
	N: Total de valores que tenemos.
*/

func main() {
	res := MAE([]int{10, 12, 15, 19, 22, 25}, []int{9, 11, 19, 18, 26, 19})
	fmt.Println(res)
}

func MAE(observedValues []int, predictedValues []int) float64 {
	var errorMargin []float64
	var stats int
	count := len(observedValues)
	var _mae float64
	for i := 0; i < count; i++ {
		_mae += math.Abs(float64(observedValues[i]) - float64(predictedValues[i]))
		errorMargin = append(errorMargin, math.Abs(float64(observedValues[i])-float64(predictedValues[i])))
	}

	_mae = _mae / float64(count)

	for i := 0; i < count; i++ {
		if errorMargin[i] < _mae {
			stats += 1
		}
	}
	stats = (stats * 100) / count
	fmt.Printf("Cantidad de valores por debajo de la media: %d%% \n", stats)
	return _mae

}
