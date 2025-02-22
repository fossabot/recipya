package units

import (
	"errors"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/neurosnap/sentences"
	"github.com/reaper47/recipya/internal/utils/regex"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	pluralizeClient *pluralize.Client
	tokenizer       *sentences.DefaultSentenceTokenizer
)

type Measurement struct {
	Quantity float64
	Unit     Unit
}

func (m Measurement) Convert(to Unit) (Measurement, error) {
	q := m.Quantity
	isCannotConvert := false

	switch m.Unit {
	case Celsius:
		switch to {
		case Celsius:
			break
		case Fahrenheit:
			q = math.Round(1.8*q + 32)
		default:
			isCannotConvert = true
		}
	case Centimeter:
		switch to {
		case Centimeter:
			break
		case Feet:
			q *= 0.03280839895
		case Inch:
			q *= 0.3937008
		case Meter:
			q *= 0.01
		case Millimeter:
			q *= 10
		case Yard:
			q *= 0.01093613
		default:
			isCannotConvert = true
		}
	case Cup:
		switch to {
		case Cup:
			break
		case Decilitre:
			q *= 2.365882
		case FlOz:
			q *= 8
		case Gallon:
			q *= 0.0625
		case Litre:
			q *= 0.2365882
		case Millilitre:
			q *= 236.5882
		case Pint:
			q *= 0.5
		case Quart:
			q *= 0.25
		case Tablespoon:
			q *= 16
		case Teaspoon:
			q *= 48
		default:
			isCannotConvert = true
		}
	case Decilitre:
		switch to {
		case Cup:
			q *= 0.4226753
		case Decilitre:
			break
		case FlOz:
			q *= 3.381402
		case Gallon:
			q *= 0.264172
		case Litre:
			q *= 0.1
		case Millilitre:
			q *= 100
		case Pint:
			q *= 0.2113376
		case Quart:
			q *= 0.1056688
		case Tablespoon:
			q *= 6.7628045404
		case Teaspoon:
			q *= 20.288413621
		default:
			isCannotConvert = true
		}
	case Inch:
		switch to {
		case Centimeter:
			q *= 2.54
		case Feet:
			q *= 0.08333333
		case Inch:
			break
		case Meter:
			q *= 0.0254
		case Millimeter:
			q *= 25.4
		case Yard:
			q *= 0.02777778
		default:
			isCannotConvert = true
		}
	case Fahrenheit:
		switch to {
		case Celsius:
			q = math.Round(0.555555555556 * (q - 32))
		case Fahrenheit:
			break
		default:
			isCannotConvert = true
		}
	case Feet:
		switch to {
		case Centimeter:
			q *= 30.48
		case Feet:
			break
		case Inch:
			q *= 0.0833333333333
		case Meter:
			q *= 0.3048
		case Millimeter:
			q *= 30480
		case Yard:
			q *= 0.3333333
		default:
			isCannotConvert = true
		}
	case FlOz:
		switch to {
		case Cup:
			q *= 0.125
		case Decilitre:
			q *= 0.2957353
		case FlOz:
			break
		case Gallon:
			q *= 0.0078125
		case Litre:
			q *= 0.02957353
		case Millilitre:
			q *= 29.57353
		case Pint:
			q *= 0.0625
		case Quart:
			q *= 0.03125
		case Tablespoon:
			q *= 2
		case Teaspoon:
			q *= 6
		default:
			isCannotConvert = true
		}
	case Gallon:
		switch to {
		case Cup:
			q *= 16
		case Decilitre:
			q *= 37.85412
		case FlOz:
			q *= 128
		case Gallon:
			break
		case Litre:
			q *= 3.785412
		case Millilitre:
			q *= 3785.412
		case Pint:
			q *= 8
		case Quart:
			q *= 4
		case Tablespoon:
			q *= 256
		case Teaspoon:
			q *= 768
		default:
			isCannotConvert = true
		}
	case Gram:
		switch to {
		case Gram:
			break
		case Kilogram:
			q *= 0.001
		case Milligram:
			q *= 1000
		case Ounce:
			q *= 0.03527396
		case Pound:
			q *= 0.002204623
		default:
			isCannotConvert = true
		}
	case Kilogram:
		switch to {
		case Gram:
			q *= 1e3
		case Kilogram:
			break
		case Milligram:
			q *= 1e6
		case Ounce:
			q *= 35.27396
		case Pound:
			q *= 2.204623
		default:
			isCannotConvert = true
		}
	case Litre:
		switch to {
		case Cup:
			q *= 4.226753
		case Decilitre:
			q *= 10
		case FlOz:
			q *= 33.81402
		case Gallon:
			q *= 0.264172
		case Litre:
			break
		case Millilitre:
			q *= 1000
		case Pint:
			q *= 2.113376
		case Quart:
			q *= 1.056688
		case Tablespoon:
			q *= 67.628045404
		case Teaspoon:
			q *= 202.88413621
		default:
			isCannotConvert = true
		}
	case Meter:
		switch to {
		case Centimeter:
			q *= 100
		case Feet:
			q *= 3.28084
		case Inch:
			q *= 39.37008
		case Meter:
			break
		case Millimeter:
			q *= 1000
		case Yard:
			q *= 1.093613
		default:
			isCannotConvert = true
		}
	case Milligram:
		switch to {
		case Gram:
			q *= 0.001
		case Kilogram:
			q *= 0.000001
		case Pound:
			q *= 2.204623e-6
		case Milligram:
			break
		case Ounce:
			q *= 0.00003527396
		default:
			isCannotConvert = true
		}
	case Millilitre:
		switch to {
		case Cup:
			q *= 0.00422675283773
		case Decilitre:
			q *= 1e-2
		case FlOz:
			q *= 0.03381402
		case Gallon:
			q *= 0.000264172037284
		case Litre:
			q *= 1e-3
		case Millilitre:
			break
		case Pint:
			q *= 0.00211337641887
		case Quart:
			q *= 0.0010566882608
		case Tablespoon:
			q *= 0.0666666666667
		case Teaspoon:
			q *= 0.2
		default:
			isCannotConvert = true
		}
	case Millimeter:
		switch to {
		case Centimeter:
			q *= 0.1
		case Feet:
			q *= 0.003280839895
		case Inch:
			q *= 0.039370
		case Meter:
			q *= 0.001
		case Millimeter:
			break
		case Yard:
			q *= 0.0010936133
		default:
			isCannotConvert = true
		}
	case Ounce:
		switch to {
		case Gram:
			q *= 28.34952
		case Kilogram:
			q *= 0.02834952
		case Milligram:
			q *= 28349.52
		case Ounce:
			break
		case Pound:
			q *= 0.0625
		default:
			isCannotConvert = true
		}
	case Pint:
		switch to {
		case Cup:
			q *= 2
		case Decilitre:
			q *= 4.731765
		case FlOz:
			q *= 16
		case Gallon:
			q *= 0.125
		case Litre:
			q *= 0.4731765
		case Millilitre:
			q *= 473.1765
		case Pint:
			break
		case Quart:
			q *= 0.5
		case Tablespoon:
			q *= 32
		case Teaspoon:
			q *= 96
		default:
			isCannotConvert = true
		}
	case Pound:
		switch to {
		case Gram:
			q *= 453.59237
		case Kilogram:
			q *= 0.45359237
		case Milligram:
			q *= 453592.37
		case Ounce:
			q *= 16
		case Pound:
			break
		default:
			isCannotConvert = true
		}
	case Tablespoon:
		switch to {
		case Cup:
			q *= 0.0625
		case Decilitre:
			q *= 0.1478677
		case FlOz:
			q *= 0.5
		case Gallon:
			q *= 0.00390625
		case Litre:
			q *= 0.01478677
		case Millilitre:
			q *= 14.78677
		case Pint:
			q *= 0.03125
		case Quart:
			q *= 0.015625
		case Tablespoon:
			break
		case Teaspoon:
			q *= 3
		default:
			isCannotConvert = true
		}
	case Teaspoon:
		switch to {
		case Cup:
			q *= 0.02083333
		case Decilitre:
			q *= 0.049289216
		case FlOz:
			q *= 0.1666667
		case Gallon:
			q *= 0.001302083
		case Litre:
			q *= 0.0049289216
		case Millilitre:
			q *= 5
		case Pint:
			q *= 0.01041667
		case Quart:
			q *= 0.005208333
		case Tablespoon:
			q *= 0.333333
		case Teaspoon:
			break
		default:
			isCannotConvert = true
		}
	case Quart:
		switch to {
		case Cup:
			q *= 4
		case Decilitre:
			q *= 9.46353
		case FlOz:
			q *= 32
		case Gallon:
			q *= 0.25
		case Litre:
			q *= 0.946353
		case Millilitre:
			q *= 946.353
		case Pint:
			q *= 2
		case Quart:
			break
		case Tablespoon:
			q *= 64
		case Teaspoon:
			q *= 192
		default:
			isCannotConvert = true
		}
	case Yard:
		switch to {
		case Centimeter:
			q *= 91.44
		case Feet:
			q *= 3
		case Inch:
			q *= 36
		case Meter:
			q *= 0.9144
		case Millimeter:
			q *= 914.4
		case Yard:
			break
		default:
			isCannotConvert = true
		}
	default:
		isCannotConvert = true
	}

	if isCannotConvert {
		return Measurement{}, errors.New("cannot convert " + m.Unit.String() + " to " + to.String())
	}
	return Measurement{Quantity: q, Unit: to}, nil
}

// String represents the Measurement as a string.
func (m Measurement) String() string {
	v := fmt.Sprintf("%.2f", m.Quantity)
	v = strings.TrimRight(strings.TrimRight(v, "0"), ".")

	unit := m.Unit.String()
	if math.Round(m.Quantity*10)*0.1 > 1 {
		unit = pluralizeClient.Plural(unit)
	}
	return v + " " + unit
}

// NewMeasurement creates a Measurement from a quantity of type int or float64
// and a unit. The creation fails when the unit is invalid.
func NewMeasurement(quantity float64, unit string) (Measurement, error) {
	unit = pluralizeClient.Singular(strings.ToLower(unit))
	unit = strings.TrimSuffix(unit, ".")

	var u Unit
	switch unit {
	case "°c", "c", "celsius", "degrees celsius", "degree celsius", "degrees c", "degree c":
		u = Celsius
	case "cm", "centimeter", "centimetre":
		u = Centimeter
	case "cup":
		u = Cup
	case "dl", "dL", "deciliter", "decilitre":
		u = Decilitre
	case "°f", "f", "fahrenheit", "degrees farenheit", "degree farenheit", "degrees fahrenheit", "degree fahrenheit", "degrees f":
		u = Fahrenheit
	case "foot", "feet", "ft", "′":
		u = Feet
	case "fluid ounce", "fl oz", "fl. oz", "floz.", "floz":
		u = FlOz
	case "gallon", "gal":
		u = Gallon
	case "g", "gram", "gramme":
		u = Gram
	case "inche", "inch", "in", `"`, `”`:
		u = Inch
	case "kg", "kilogram", "kilogramme":
		u = Kilogram
	case "l", "litre", "liter":
		u = Litre
	case "m", "meter", "metre":
		u = Meter
	case "mg", "milligram", "milligramme":
		u = Milligram
	case "ml", "milliliter", "millilitre", "cc":
		u = Millilitre
	case "mm", "millimeter", "millimetre":
		u = Millimeter
	case "ounce", "oz":
		u = Ounce
	case "pint", "pt", "fl pt", "fl. pt":
		u = Pint
	case "lb", "#", "pound":
		u = Pound
	case "quart", "qt", "fl qt", "fl. qt":
		u = Quart
	case "tablespoon", "tbl", "tbs", "tb", "tbsp":
		u = Tablespoon
	case "teaspoon", "tsp":
		u = Teaspoon
	case "yard":
		u = Yard
	default:
		return Measurement{}, errors.New("unit " + unit + " is unsupported")
	}
	return Measurement{Quantity: quantity, Unit: u}, nil
}

// ConvertParagraph converts the paragraph to the desired System.
func ConvertParagraph(paragraph string, from, to System) string {
	tokens := tokenizer.Tokenize(paragraph)
	xs := make([]string, len(tokens))
	for i, sentence := range tokens {
		s, err := ConvertSentence(sentence.Text, from, to)
		if err != nil {
			xs[i] = sentence.Text
			continue
		}
		xs[i] = s
	}
	return strings.Join(xs, "")
}

// ConvertSentence converts the sentence to the desired System.
func ConvertSentence(input string, from, to System) (string, error) {
	if from == to {
		return input, errors.New("the measurement system is unchanged")
	}

	input = replaceKeyboardFractions(input)

	matches := regex.Unit.FindStringSubmatch(input)
	if matches == nil {
		return "", errors.New("sentence has no measurement")
	}

	q, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return parseIrregularQuantity(input, matches, regex.Unit, to)
	}

	m, err := NewMeasurement(q, matches[len(matches)-1])
	if err != nil {
		return "", err
	}

	converted := convertMeasurement(m, to)
	return regex.Unit.ReplaceAllString(input, converted.String()), nil
}

func replaceKeyboardFractions(input string) string {
	keyboardFractions := map[string]string{
		"½": "1/2",
		"⅓": "1/3",
		"⅔": "2/3",
		"¼": "1/4",
		"¾": "3/4",
		"⅕": "1/5",
		"⅖": "2/5",
		"⅗": "3/5",
		"⅘": "4/5",
		"⅙": "1/6",
		"⅚": "5/6",
		"⅐": "1/7",
		"⅛": "1/8",
		"⅜": "3/8",
		"⅝": "5/8",
		"⅞": "7/8",
		"⅑": "1/9",
		"⅒": "1/10",
	}

	for key, value := range keyboardFractions {
		input = strings.Replace(input, key, value, -1)
	}
	return input
}

func parseIrregularQuantity(input string, matches []string, re *regexp.Regexp, to System) (string, error) {
	match := strings.Replace(matches[1], "-", " ", 1)
	parts := strings.Split(match, " ")
	convertedParts := make([]string, len(parts))

	errCount := 0
	for i, part := range parts {
		q, err := strconv.ParseFloat(part, 64)
		if err != nil {
			if strings.Contains(part, "/") {
				subParts := strings.Split(part, "/")

				numerator, err := strconv.ParseFloat(subParts[0], 64)
				if err != nil {
					convertedParts[i] = subParts[i]
					continue
				}

				denominator, err := strconv.ParseFloat(subParts[1], 64)
				if err != nil {
					convertedParts[i] = subParts[i]
					continue
				}

				m, _ := NewMeasurement(numerator/denominator, matches[len(matches)-1])
				converted := convertMeasurement(m, to)
				if i > 0 {
					prev, err := strconv.ParseFloat(convertedParts[i-1], 64)
					if err == nil {
						convertedParts[i-1] = ""
						converted.Quantity += prev
					}
				}

				if i == len(parts)-1 {
					convertedParts[i] = converted.String()
				} else {
					v := fmt.Sprintf("%.2f", converted.Quantity)
					v = strings.TrimRight(strings.TrimRight(v, "0"), ".")
					convertedParts[i] = v
				}
				continue
			}

			convertedParts[i] = parts[i]
			errCount++
			continue
		}

		m, _ := NewMeasurement(q, matches[len(matches)-1])
		converted := convertMeasurement(m, to)
		if i == len(parts)-1 {
			convertedParts[i] = converted.String()
		} else {
			v := fmt.Sprintf("%.2f", converted.Quantity)
			v = strings.TrimRight(strings.TrimRight(v, "0"), ".")
			convertedParts[i] = v
		}
	}

	if errCount == len(parts) {
		return "", errors.New("unable to parse")
	}

	xs := slices.DeleteFunc(convertedParts, func(s string) bool { return s == "" })
	return re.ReplaceAllString(input, strings.Join(xs, " ")), nil
}

func convertMeasurement(m Measurement, to System) Measurement {
	q := m.Quantity
	var converted Measurement
	switch to {
	case ImperialSystem:
		switch m.Unit {
		case Celsius:
			converted, _ = m.Convert(Fahrenheit)
		case Centimeter:
			if q < 100 {
				converted, _ = m.Convert(Inch)
			} else if q < 30480 {
				converted, _ = m.Convert(Feet)
			} else {
				converted, _ = m.Convert(Yard)
			}
		case Decilitre:
			if q < 0.1478676 {
				converted, _ = m.Convert(Teaspoon)
			} else if q < 0.2957353 {
				converted, _ = m.Convert(Tablespoon)
			} else if q < 1.182941 {
				converted, _ = m.Convert(FlOz)
			} else if q < 4.731765 {
				converted, _ = m.Convert(Cup)
			} else if q < 9.46353 {
				converted, _ = m.Convert(Pint)
			} else if q < 37.85412 {
				converted, _ = m.Convert(Quart)
			} else {
				converted, _ = m.Convert(Gallon)
			}
		case Gram:
			if q < 2834 {
				converted, _ = m.Convert(Pound)
			} else {
				converted, _ = m.Convert(Ounce)
			}
		case Kilogram:
			if q < 35.27396 {
				converted, _ = m.Convert(Pound)
			} else {
				converted, _ = m.Convert(Ounce)
			}
		case Litre:
			if q < 0.01478676 {
				converted, _ = m.Convert(Teaspoon)
			} else if q < 0.02957353 {
				converted, _ = m.Convert(Tablespoon)
			} else if q < 0.1182941 {
				converted, _ = m.Convert(FlOz)
			} else if q < 0.4731765 {
				converted, _ = m.Convert(Cup)
			} else if q < 0.946353 {
				converted, _ = m.Convert(Pint)
			} else if q < 3.785412 {
				converted, _ = m.Convert(Quart)
			} else {
				converted, _ = m.Convert(Gallon)
			}
		case Meter:
			if q < 10 {
				converted, _ = m.Convert(Inch)
			} else if q < 305 {
				converted, _ = m.Convert(Feet)
			} else {
				converted, _ = m.Convert(Yard)
			}
		case Milligram:
			if q < 28349.52 {
				converted, _ = m.Convert(Pound)
			} else {
				converted, _ = m.Convert(Ounce)
			}
		case Millilitre:
			if q < 15 {
				converted, _ = m.Convert(Teaspoon)
			} else if q < 30 {
				converted, _ = m.Convert(Tablespoon)
			} else if q < 118.2941 {
				converted, _ = m.Convert(FlOz)
			} else if q < 946.353 {
				converted, _ = m.Convert(Cup)
			} else if q < 1892.706 {
				converted, _ = m.Convert(Pint)
			} else if q < 3785.412 {
				converted, _ = m.Convert(Quart)
			} else {
				converted, _ = m.Convert(Gallon)
			}
		case Millimeter:
			if q < 1000 {
				converted, _ = m.Convert(Inch)
			} else if q < 304800 {
				converted, _ = m.Convert(Feet)
			} else {
				converted, _ = m.Convert(Yard)
			}
		}
	case MetricSystem:
		switch m.Unit {
		case Cup:
			if q < 0.4226753 {
				converted, _ = m.Convert(Millilitre)
			} else if q < 4.226753 {
				converted, _ = m.Convert(Decilitre)
			} else {
				converted, _ = m.Convert(Litre)
			}
		case Fahrenheit:
			converted, _ = m.Convert(Celsius)
		case Feet:
			if q < 1 {
				converted, _ = m.Convert(Millimeter)
			} else if q < 304.8 {
				converted, _ = m.Convert(Centimeter)
			} else {
				converted, _ = m.Convert(Meter)
			}
		case FlOz:
			if q < 0.8453506 {
				converted, _ = m.Convert(Millilitre)
			} else if q < 8.453506 {
				converted, _ = m.Convert(Decilitre)
			} else {
				converted, _ = m.Convert(Litre)
			}
		case Gallon:
			if q < 0.0264172 {
				converted, _ = m.Convert(Millilitre)
			} else if q < 0.264172 {
				converted, _ = m.Convert(Decilitre)
			} else {
				converted, _ = m.Convert(Litre)
			}
		case Inch:
			if q < 1 {
				converted, _ = m.Convert(Millimeter)
			} else if q < 39.37008 {
				converted, _ = m.Convert(Centimeter)
			} else {
				converted, _ = m.Convert(Meter)
			}
		case Ounce:
			if q < 0.03527396 {
				converted, _ = m.Convert(Milligram)
			} else if q < 35.27396 {
				converted, _ = m.Convert(Gram)
			} else {
				converted, _ = m.Convert(Kilogram)
			}
		case Pint:
			if q < 0.2113376 {
				converted, _ = m.Convert(Millilitre)
			} else if q < 2.113376 {
				converted, _ = m.Convert(Decilitre)
			} else {
				converted, _ = m.Convert(Litre)
			}
		case Pound:
			if q < 0.002204623 {
				converted, _ = m.Convert(Milligram)
			} else if q < 2 {
				converted, _ = m.Convert(Gram)
			} else {
				converted, _ = m.Convert(Kilogram)
			}
		case Quart:
			if q < 0.1056688 {
				converted, _ = m.Convert(Millilitre)
			} else if q < 1.056688 {
				converted, _ = m.Convert(Decilitre)
			} else {
				converted, _ = m.Convert(Litre)
			}
		case Tablespoon:
			if q < 6.762804 {
				converted, _ = m.Convert(Millilitre)
			} else if q < 67.62804 {
				converted, _ = m.Convert(Decilitre)
			} else {
				converted, _ = m.Convert(Litre)
			}
		case Teaspoon:
			if q < 20.28841 {
				converted, _ = m.Convert(Millilitre)
			} else if q < 202.8841 {
				converted, _ = m.Convert(Decilitre)
			} else {
				converted, _ = m.Convert(Litre)
			}
		case Yard:
			if q < 0.1093613 {
				converted, _ = m.Convert(Millimeter)
			} else if q < 1.093613 {
				converted, _ = m.Convert(Centimeter)
			} else {
				converted, _ = m.Convert(Meter)
			}
		}
	}
	return converted
}

// DetectMeasurementSystemFromSentence determines the System used in the text.
func DetectMeasurementSystemFromSentence(text string) System {
	for _, part := range strings.Split(text, " ") {
		switch pluralizeClient.Singular(strings.ToLower(part)) {
		case "°c", "celsius", "cm", "centimeter", "centimetre", "dl", "dL", "deciliter", "decilitre", "g", "gram", "gramme", "kg", "kilogram", "kilogramme", "l", "litre", "liter", "m", "meter", "metre", "mg", "milligram", "milligramme", "ml", "milliliter", "millilitre", "cc", "mm", "millimeter", "millimetre":
			return MetricSystem
		case "cup", "°f", "f", "farenheit", "foot", "feet", "ft", "′", "ounce", "oz", "gallon", "gal", "gill", "inche", "inch", `"`, "pint", "pt", "lb", "#", "pound", "quart", "qt", "tablespoon", "tbl", "tbs", "tb", "tbsp", "teaspoon", "tsp", "yard":
			return ImperialSystem
		}
	}
	return InvalidSystem
}

func init() {
	pluralizeClient = pluralize.NewClient()

	rules := map[string]string{
		Celsius.String():    Celsius.String(),
		Centimeter.String(): Centimeter.String(),
		Decilitre.String():  Decilitre.String(),
		Fahrenheit.String(): Fahrenheit.String(),
		FlOz.String():       FlOz.String(),
		Gram.String():       Gram.String(),
		Kilogram.String():   Kilogram.String(),
		Litre.String():      Litre.String(),
		Meter.String():      Meter.String(),
		Milligram.String():  Milligram.String(),
		Millilitre.String(): Millilitre.String(),
		Millimeter.String(): Millimeter.String(),
		Ounce.String():      Ounce.String(),
		Pound.String():      Pound.String(),
		Tablespoon.String(): Tablespoon.String(),
		Teaspoon.String():   Teaspoon.String(),
		Quart.String():      Quart.String(),
	}
	for k, v := range rules {
		pluralizeClient.AddIrregularRule(k, v)
	}

	b, err := fs.ReadFile("lang/english.json")
	if err != nil {
		panic(err)
	}

	training, err := sentences.LoadTraining(b)
	if err != nil {
		panic(err)
	}

	tokenizer = sentences.NewSentenceTokenizer(training)
}
