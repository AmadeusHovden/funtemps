package main

import (
	"flag"
	"fmt"

	"github.com/AmadeusHovden/funtemps/conv"
	"github.com/AmadeusHovden/funtemps/funfacts"
)

// Flere flag-variabler? Om de uliek funfacts?
var Fahrenheit float64
var Celsius float64
var Kelvin float64

var out string

var funFacts string
var t string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&Fahrenheit, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&Celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&Kelvin, "K", 0.0, "temperartur i grader kelvin")

	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - fahrenheit, K- Kelvin")
	flag.StringVar(&funFacts, "funfacts", "", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "", "temperaturskala for å vise funfacts om Solen, Månen eller Jorden")

}

func main() {

	flag.Parse()

	// sjekker om flaggene F, C og K er angitt samtidig
	if isFlagPassed("F") && isFlagPassed("C") || isFlagPassed("F") && isFlagPassed("K") || isFlagPassed("C") && isFlagPassed("K") {
		fmt.Println("Bare ett av flaggene F, C og K kan angis samtidig")
		return
	}

	// sjekker om flaggene F, C og K er angitt med flagget out
	if isFlagPassed("F") || isFlagPassed("C") || isFlagPassed("K") {
		if !isFlagPassed("out") {
			fmt.Println("Flaggene F, C og K må angis med flagget out")
			return
		}
		// Varsler om ugydlig argument, hvis skrive feil, eller ikke store bokstaver
		if out != "C" && out != "F" && out != "K" {
			fmt.Println("Ugyldig argument for flagget out, må være enten C, F eller K")
			return
		}
		/*
			// conv flaggene kan ikke oppgis med flagget funfacts.
			if isFlagPassed("funfacts") {
				fmt.Println("Flaggene F, C og K kan ikke angis samtidig med flagget funfacts")
				return
			}
		*/

		// Beregner og viser resultatet
		if out == "C" {
			if isFlagPassed("F") {
				result := conv.FahrenheitToCelsius(Fahrenheit)
				fmt.Printf("%.2f°F er %.2f°C\n", Fahrenheit, result)
			} else if isFlagPassed("K") {
				result := conv.KelvinToCelsius(Kelvin)
				fmt.Printf("%.2f°K er %.2f°C\n", Kelvin, result)
			}
		} else if out == "F" {
			if isFlagPassed("C") {
				result := conv.CelsiusToFahrenheit(Celsius)
				fmt.Printf("%.2f°C er %.2f°F\n", Celsius, result)
			} else if isFlagPassed("K") {
				result := conv.KelvinToFahrenheit(Kelvin)
				fmt.Printf("%.2f°K er %.2f°F\n", Kelvin, result)
			}
		} else if out == "K" {
			if isFlagPassed("C") {
				result := conv.CelsiusToKelvin(Celsius)
				fmt.Printf("%.2f°C er %.2f°K\n", Celsius, result)
			} else if isFlagPassed("F") {
				result := conv.FahrenheitToKelvin(Fahrenheit)
				fmt.Printf("%.2f°F er %.2f°K\n", Fahrenheit, result)
			}
		}
	}

	if isFlagPassed("funfacts") && isFlagPassed("t") && t == "C" {
		facts := funfacts.GetFunFacts(funFacts)
		for i, fact := range facts {
			fmt.Printf("%d: %s\n", i+1, fact)
		}
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje.
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// input terminal
// go run main.go -(C, F, K) GRADER -out (C, F, K)
// go run main.go -funfacts (Sun, Luna, Terra) -t C
