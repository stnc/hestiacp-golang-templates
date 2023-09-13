package main

import "fmt"

type Start struct {
	Number1        float64
	Number2        float64
	Text           string
	RequestHandler RequestHandler
}
type RequestHandler interface {
	Addition() float64
	Subtraction() float64
	Multiplication() float64
	Division() float64
}
type Calc struct{}

var a = &Start{Number1: 10, Number2: 2, RequestHandler: &Calc{}}

// ProcessRequest handles a request passed from Alexa
func (s *Start) ProcessRequest(whoIs string) float64 {
	var val float64
	switch whoIs {
	case "Addition":
		val = s.RequestHandler.Addition()
	case "Subtraction":
		val = s.RequestHandler.Subtraction()
	case "Multiplication":
		val = s.RequestHandler.Multiplication()
	case "Division":
		val = s.RequestHandler.Division()
	}
	return val
}

func (c *Calc) Addition() float64 {
	return a.Number1 + a.Number2
}

func (c *Calc) Subtraction() float64 {
	//return Number1 - Number2
	return a.Number1 - a.Number2
}

func (c *Calc) Multiplication() float64 {
	return a.Number1 * a.Number2
}

func (c *Calc) Division() float64 {
	return a.Number1 / a.Number2
}

func main() {
	//a.Number1 = 20
	//a.Number2 = 4
	fmt.Println("Addition")
	fmt.Println(a.ProcessRequest("Addition"))

	fmt.Println("Subtraction")
	fmt.Println(a.ProcessRequest("Subtraction"))

	fmt.Println("Multiplication")
	fmt.Println(a.ProcessRequest("Multiplication"))

	fmt.Println("Division")
	fmt.Println(a.ProcessRequest("Division"))

}
