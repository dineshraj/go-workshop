package fundamentals

const boilingF float64 = 212.0
const freezingF float64 = 32.0
const absoluteZeroF float64 = -459.67

// Constants ...
func Constants() (float64, float64, float64) {
	var celsius []float64 = Convert(boilingF, freezingF, absoluteZeroF)
	return celsius[0], celsius[1], celsius[2]
}
