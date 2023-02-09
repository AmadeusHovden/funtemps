package funfacts

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

var funFacts = FunFacts{
	Sun: []string{
		"Visste du at temperaturen i solens kjerne er 15 000 000℃",
		"Visste du at temperaturen på ytre lag av solen er 5778K",
	},
	Terra: []string{
		"Visste du at høyeste målt temperatur på jordens overflate er 56℃",
		"Visste du at laveste temperatur på jordens overflate er -89.4℃",
		"Visste du at temperaturen i jordens indre kjerne er 9392K ",
	},
	Luna: []string{
		"Visste du at temperaturen på månens overflate om natten -183℃",
		"Visste du at temperaturen på månens overflate om dagen 106℃",
	},
}

func GetFunFacts(about string) []string {
	switch about {
	case "Sun":
		return funFacts.Sun

	case "Terra":
		return funFacts.Terra

	case "Luna":
		return funFacts.Luna

	default:
		return []string{"ingen funfacts"}
	}
}
