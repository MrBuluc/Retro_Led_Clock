package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

const (
	holder          = "█"
	seperatorHolder = "@"
)

type placeholder [5]string

func main() {
	// Step #1

	base := fmt.Sprintf("%s%s%s", holder, holder, holder)

	zero, one, two, three, four, five, six, seven, eight, nine, A, L, R, M, exclamationMark, empty :=
		placeholder{
			base,
			fillHolder("%s %s"),
			fillHolder("%s %s"),
			fillHolder("%s %s"),
			base,
		}, placeholder{
			fillHolder("%s%s "),
			fillHolder(" %s "),
			fillHolder(" %s "),
			fillHolder(" %s "),
			base,
		}, placeholder{
			base,
			fillHolder("  %s"),
			base,
			fillHolder("%s  "),
			base,
		}, placeholder{
			base,
			fillHolder("  %s"),
			base,
			fillHolder("  %s"),
			base,
		}, placeholder{
			fillHolder("%s %s"),
			fillHolder("%s %s"),
			base,
			fillHolder("  %s"),
			fillHolder("  %s"),
		}, placeholder{
			base,
			fillHolder("%s  "),
			base,
			fillHolder("  %s"),
			base,
		}, placeholder{
			base,
			fillHolder("%s  "),
			base,
			fillHolder("%s %s"),
			base,
		}, placeholder{
			base,
			fillHolder("  %s"),
			fillHolder("  %s"),
			fillHolder("  %s"),
			fillHolder("  %s"),
		}, placeholder{
			base,
			fillHolder("%s %s"),
			base,
			fillHolder("%s %s"),
			base,
		}, placeholder{
			base,
			fillHolder("%s %s"),
			base,
			fillHolder("  %s"),
			base,
		}, placeholder{
			base,
			fillHolder("%s %s"),
			base,
			fillHolder("%s %s"),
			fillHolder("%s %s"),
		}, placeholder{
			fillHolder("%s  "),
			fillHolder("%s  "),
			fillHolder("%s  "),
			fillHolder("%s  "),
			base,
		}, placeholder{
			fillHolder("%s%s "),
			fillHolder("%s %s"),
			fillHolder("%s%s "),
			fillHolder("%s %s"),
			fillHolder("%s %s"),
		}, placeholder{
			fillHolder("%s %s"),
			base,
			fillHolder("%s %s"),
			fillHolder("%s %s"),
			fillHolder("%s %s"),
		}, placeholder{
			fillHolder("%s  "),
			fillHolder("%s  "),
			fillHolder("%s  "),
			fillHolder("  "),
			fillHolder("%s  "),
		}, placeholder{
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		}

	digits, alarm := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}, [...]placeholder{empty, A, L, A, R, M, exclamationMark, empty}

	// prints digits top
	//				  to
	//				down
	/* for _, digit := range digits {
		for _, line := range digit {
			fmt.Println(line)
		}

		fmt.Println("")
	} */

	// prints digits side-by-side
	/* for line := range digits[0] {
		for digit := range digits {
			fmt.Print(digits[digit][line])
			fmt.Print("  ")
		}

		fmt.Println("")
	} */

	//Step #2
	//now := time.Now()
	filledSeperator := fillSeperator(" %s ")
	seperator := placeholder{
		"   ",
		filledSeperator,
		"   ",
		filledSeperator,
		"   ",
	}
	/* clock := [...]placeholder{
		digits[now.Hour()/10], digits[now.Hour()%10],
		seperator,
		digits[now.Minute()/10], digits[now.Minute()%10],
		seperator,
		digits[now.Second()/10], digits[now.Second()%10],
	} */

	/* for line := range clock[0] {
		for digit := range clock {
			fmt.Print(clock[digit][line])
			fmt.Print(" ")
		}

		fmt.Println("")
	} */

	//Step #3
	for {
		screen.Clear()
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			seperator,
			digits[min/10], digits[min%10],
			seperator,
			digits[sec/10], digits[sec%10],
		}

		printClock(clock, seperator, sec)

		if sec%10 == 0 {
			fmt.Println("")
			printClock(alarm, seperator, sec)
		}

		fmt.Println("")
		time.Sleep(time.Second)
	}
}

func fillHolder(format string) string {
	return strings.Replace(format, "%s", holder, -1)
}

func fillSeperator(format string) string {
	return strings.Replace(format, "%s", seperatorHolder, -1)
}

func printClock(clock [8]placeholder, seperator placeholder, sec int) {
	for line := range clock[0] {
		for i, digit := range clock {
			next := clock[i][line]
			if digit == seperator && sec%2 == 0 {
				next = "   "
			}
			fmt.Print(next)
			fmt.Print(" ")
		}

		fmt.Println("")
	}
}
