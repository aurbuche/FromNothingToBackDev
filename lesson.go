package main

import (
	"fmt"
	// "bufio" // -> Use for I/O
	"os" // -> Fonction OS
	"io"
	// "strconv"
)

func CreateLogFile() *os.File {
	logFile, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	return logFile
}

func WriteLog(logMsg string, logFile io.Writer) {
	fmt.Fprint(logFile, logMsg)
}

func addition(a int, b int) int {
	return a + b
}

/*func main() {
	a := 50
	b := 5
	c := addition(a, b)
	fmt.Println(a, " + ", b, " = ", c)
}*/

// Fonction anonyme:

func rajouterDix(a float64, fAnonyme func(float64) float64) /**/ {
    operation := fAnonyme(a) // Appel à notre fonction anonyme
    result := operation + 10
    fmt.Println(result)
}

/*func main() {
    // stockage de notre fonction anonyme dans une variable
    racineCarree := func(x float64) float64 { return math.Sqrt(x) }
    rajouterDix(9, racineCarree)
	rajouterDix(5, func(x float64) float64 { return math.Pow(x, 2) })
}*/