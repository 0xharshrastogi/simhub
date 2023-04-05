package main

import (
	profitcalculator "github.com/harshrastogiexe/cmd/cli/profit_calculator"
	"github.com/harshrastogiexe/lib/go/common"
)

func main() {
	calc := profitcalculator.New(common.Magnates)
	calc.Calculate("L")

}
