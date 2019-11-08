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
			if len(pipes[i].BlockedBy) > 0 {
				fmt.Printf("Blocked by: %v\n", strings.Join(pipes[i].BlockedBy, ","))
			}
		}
	}
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
