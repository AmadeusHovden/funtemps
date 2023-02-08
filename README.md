# funtemps
brukes til første testing oppgave (13.feb)

## MÅL
 I denne oppgaven skal du planlegge og implementere et program (kommando, fordi at den skal starte fra kommandolinje med inn-data på kommandolinje og skrive ut svar/ut-data til kommandolinje), som regner om fra en temperaturskala til en annen. I tillegg skal ditt program kunne skrive ut noen "fun-facts" om temperaturer på Solen, Månen og Jorden.



## fra janis read.me

Notat angående testing av Float64

Når man sammenlignet to verdier av typen float64, spiller presisjon er rolle. Alle reelle tall, representert i en datamaskin, har begrenset antall desimaler, dvs. verdien er avrundet.

I testene kan man ikke teste på absolutte verdier til flyttall, så man må bestemme hvilke nøyaktighet man trenger og teste mot den. Man kan sammenligne to flyttall med funksjonen Abs() fra pakken math.

difference := math.Abs(a - b)
Intuitivt virker dette bra, men en liten forskjell på verdien for to store tall, kan være en stor forskjell for to små tall. Metoden man bruker er beregning av en relativ forskjell ved å dele differansen med den absolutte verdien til den andre input (b) => difference/math.Abs(b). Da blir funksjonen slik:

func withinTolerance(a, b, error float64) bool {
  // Først sjekk om tallene er nøyaktig like
  if a == b {
    return true
  }

  difference := math.Abs(a - b)

  // Siden vi skal dele med b, må vi sjekke om den er 0
  // Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
  // som vi aksepterer
  if b == 0 {
    return difference < error
  }

  // Tilslutt sjekk den relative differanse mot feilmargin
  return (difference/math.Abs(b)) < error
}

### Det er anbefalt å bruke denne funksjonen i testene hvor float64 er innvolvert. Testen vi hadde foreslått var

for _, tc := range tests {
  got := FarhenheitToCelsius(tc.input)
  if !reflect.DeepEqual(tc.want, got) {
    t.Errorf("expected: %v, got: %v", tc.want, got)
  }
}

### Og hvis vi erstatter !reflect.DeepEqual(tc.want, got) med !withinTolerance(tc.want, got, 1e-12) så får vi

for _, tc := range tests {
  got := FarhenheitToCelsius(tc.input)
  if !withinTolerance(tc.want, got, 1e-12) {
    t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
  }
}


# Referanser
Gerardi, R. (2021, December 21). Testing Floating Point Numbers in Go - The Pragmatic Programmers - Medium. Medium; The Pragmatic Programmers. https://medium.com/pragmatic-programmers/testing-floating-point-numbers-in-go-9872fe6de17f
