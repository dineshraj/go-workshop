package fundamentals

// fToC ...
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func Convert(n ...float64) []float64 {
	converted := []float64{}
	for _, f := range n {
		converted = append(converted, fToC(f))
	}
	return converted
}
