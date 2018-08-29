package space

const (
	EARTH_YEAR_SEC        = 31557600
	MERCURY_TO_EARTH_YEAR = 0.2408467
	VENUS_TO_EARTH_YEAR   = 0.61519726
	MARS_TO_EARTH_YEAR    = 1.8808158
	JUPITER_TO_EARTH_YEAR = 11.862615
	SATURN_TO_EARTH_YEAR  = 29.447498
	URANUS_TO_EARTH_YEAR  = 84.016846
	NEPTUNE_TO_EARTH_YEAR = 164.79132
)

type Planet string

// Age count age in years on specified planet
func Age(seconds float64, planet Planet) (years float64) {
	if seconds == 0 {
		return 0.0
	}
	earthYears := seconds / EARTH_YEAR_SEC
	var planetMultiplyer float64
	switch planet {
	case "Earth":
		planetMultiplyer = 1
	case "Mercury":
		planetMultiplyer = MERCURY_TO_EARTH_YEAR
	case "Venus":
		planetMultiplyer = VENUS_TO_EARTH_YEAR
	case "Mars":
		planetMultiplyer = MARS_TO_EARTH_YEAR
	case "Jupiter":
		planetMultiplyer = JUPITER_TO_EARTH_YEAR
	case "Saturn":
		planetMultiplyer = SATURN_TO_EARTH_YEAR
	case "Uranus":
		planetMultiplyer = URANUS_TO_EARTH_YEAR
	case "Neptune":
		planetMultiplyer = NEPTUNE_TO_EARTH_YEAR
	}
	return float64(earthYears) / planetMultiplyer
}
