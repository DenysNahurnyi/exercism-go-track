package scale

func Scale(tonic, interval string) []string {
	// CAP letters withoud pair are diatonic, len(diatonic) == 7
	// Smallest interval is "m" (half-step), "M" -> whole step
	chromaticScale := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	return chromaticScale
}
