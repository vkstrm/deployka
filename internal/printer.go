package internal

import (
	"fmt"
	"strings"
)

// PrintPipes : Prints the pipes
func PrintPipes(pipes []ResponseBody) {
	i := 0
	for ; i < len(pipes); i++ {
		fmt.Printf("%v\t ", pipes[i].Pipename)
		if pipes[i].Allowed {
			fmt.Printf("%s\n", PrintWithColour("Allowed", green))
		} else {
			fmt.Printf("%s\t", PrintWithColour("Blocked", red))
			var at string
			if len(pipes[i].BlockedBy) == 1 {
				at = "At"
			} else if len(pipes[i].BlockedBy) > 1 {
				at = "Latest at"
				fmt.Printf("Blocked by: %v\t%s: %v", strings.Join(pipes[i].BlockedBy, ", "), at, datePrint(pipes[i].BlockedAt))
			}
			fmt.Print("\n")
		}
	}
}

func datePrint(date map[string]string) string {
	return fmt.Sprintf("%v/%v/%v %v:%v", date["year"], date["month"], date["dateay"], date["hour"], date["min"])
}

// Colours for the printing
const (
	green = "32"
	red   = "31"
)

// PrintWithColour : print the string in green
func PrintWithColour(str string, colour string) string {
	return fmt.Sprintf("\033[1;%sm%s\033[0m", colour, str)
}

// TitleCardPrint : Print this awesome stuff!
func TitleCardPrint() {
	fmt.Println("************************************************************************")
	fmt.Println("     //////  ////// //////  //        //////   //    //  //  //    /////")
	fmt.Println("    //   // //     //   // //       //    //   //  //   // //     //  //")
	fmt.Println("   //   // ////// //////  //       //    //    ////    ////      //   //")
	fmt.Println("  //   // //     //      //       //    //     //     //  //    ////////")
	fmt.Println(" //   // //     //      //       //    //     //     //   //   //     //")
	fmt.Println("//////  ////// //      ////////  //////      //     //	   // //      //")
	fmt.Println("************************************************************************")
}
