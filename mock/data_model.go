package mock

type MockDataModel struct{}

func (d *MockDataModel) GetBlacklistUnits() (map[string]int8, error) {
	return map[string]int8{
		"¼":           0,
		"½":           0,
		"¾":           0,
		"cc":          0,
		"celsius":     0,
		"centimeter":  0,
		"centimetre":  0,
		"cm":          0,
		"cube":        0,
		"cup":         0,
		"deciliter":   0,
		"decilitre":   0,
		"dl":          0,
		"fahrenheit":  0,
		"fl":          0,
		"fluid":       0,
		"fresh":       0,
		"g":           0,
		"gal":         0,
		"gallon":      0,
		"gill":        0,
		"gram":        0,
		"gramme":      0,
		"imperial":    0,
		"in":          0,
		"inch":        0,
		"kg":          0,
		"kilogram":    0,
		"kilogramme":  0,
		"l":           0,
		"lb":          0,
		"liter":       0,
		"litre":       0,
		"m":           0,
		"meter":       0,
		"metre":       0,
		"mg":          0,
		"milligram":   0,
		"milligramme": 0,
		"milliliter":  0,
		"millilitre":  0,
		"millimeter":  0,
		"millimetre":  0,
		"ml":          0,
		"mm":          0,
		"ounce":       0,
		"optional":    0,
		"oz":          0,
		"package":     0,
		"pint":        0,
		"pound":       0,
		"quart":       0,
		"tablespoon":  0,
		"tbl":         0,
		"tbs":         0,
		"tbsp":        0,
		"teaspoon":    0,
		"tsp":         0,
	}, nil
}

func (d *MockDataModel) GetFruitsVeggies() (map[string]int8, error) {
	return map[string]int8{
		"kiwi":       0,
		"kiwis":      0,
		"onion":      0,
		"potato":     0,
		"strawberry": 0,
		"squash":     0,
		"tomato":     0,
	}, nil
}
