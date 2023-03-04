package secuencias

func SubsecuenciaSumaMaxima(arreglo []int) (int, int, int) {
	sumaMaxima := 0
	posInicial, posFinal := -1, -1
	for i := 0; i < len(arreglo); i++ {
		sumaLocal := 0
		for j := i; j < len(arreglo); j++ {
			sumaLocal += arreglo[j]
			if sumaLocal > sumaMaxima {
				sumaMaxima = sumaLocal
				posInicial = i
				posFinal = j
			}
		}
	}
	return sumaMaxima, posInicial, posFinal
}
