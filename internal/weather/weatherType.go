package weather

// Map of weather types with dedicated emoji symbol
var weatherTypes = map[string][]rune{
	"Thubder": {'\u26c8', '\ufe0f'},
	"Drizzle": {'\u2614'},
	"Rain":    {'\U0001f327', '\ufe0f'},
	"Snow":    {'\u2744', '\ufe0f'},
	"Clear":   {'\u2600', '\ufe0f'},
	"Fog":     {'\U0001f301'},
	"Clouds":  {'\u2601', '\ufe0f'},
}

func GetWeatherTypeSymbol(w *Weather) (symbol string) {
	return string(weatherTypes[w.WeatherType])
}
