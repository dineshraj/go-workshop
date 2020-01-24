package fundamentals

// ShortVars ...
func ShortVars() (float64, float64, float64) {
	boilingF := 212.0
	freezingF := 32.0
	absoluteZeroF := -459.67

	celsius := []float64(Convert(boilingF, freezingF, absoluteZeroF))

	return celsius[0], celsius[1], celsius[2]
}
