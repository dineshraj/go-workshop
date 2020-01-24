package fundamentals

// FToC ...
func FToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func Convert(n ...float64) []float64 {
	converted := []float64{}
	for _, f := range n {
		converted = append(converted, FToC(f))
	}
	return converted
}

// ShortVars ...
func ShortVars() (float64, float64, float64) {
	boilingF := 212.0
	freezingF := 32.0
	absoluteZeroF := -459.67

	celsius := []float64(Convert(boilingF, freezingF, absoluteZeroF))

	return celsius[0], celsius[1], celsius[2]
}
