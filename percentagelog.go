package percentagelog

import (
	"fmt"
	"strings"
	"time"
)

var width, height int

const (
	rightSpacing         = 4
	percentageLength     = 4
	barPercentageSpacing = 2
	nameBarSpacing       = 4
)

func getMaxWordLength(ps ...Printable) int {
	mLen := 0
	for _, p := range ps {
		if l := len(p.String()); l > mLen {
			mLen = l
		}
	}
	return mLen
}

func print(ps ...Printable) {
	fmt.Printf("\r\033[%dA", len(ps))
	fmt.Print(sprint(ps...))
}

func sprint(ps ...Printable) string {
	output := ""
	maxWordLength := getMaxWordLength(ps...)
	for _, p := range ps {
		barLengths := int(p.Percentage() / 2)
		bar := strings.Repeat("=", barLengths)
		bar += ">"
		output += fmt.Sprintf("%-*s | %-*s %6.2f %%\n", maxWordLength, p.String(), 51, bar, p.Percentage())
	}
	return output
}

//Blocking function to print output until finished
func PrintUntilFinished(ps ...Printable) {
	fmt.Print(strings.Repeat("\n", len(ps)))
	for {
		print(ps...)
		time.Sleep(time.Millisecond * 200)
		complete := true
		for _, p := range ps {
			if p.Percentage() < 100 {
				complete = false
			}
		}
		if complete {
			print(ps...)
			break
		}
	}
}
