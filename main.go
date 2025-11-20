/*
 * Author Zoe Wrubleski
 * Version 1.0.0
 * Date 2025-11-20
 * Fileoverview This program displays sentences in the oppisite order than input
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// set variables
	var line1 string
	var line2 string
	var line3 string
	var line4 string

	reader := bufio.NewReader(os.Stdin)

	// prompt user for sentences 
	fmt.Print("Enter your first sentence: ")
	line1, _ = reader.ReadString('\n')
	line1 = strings.TrimSpace(line1)
	fmt.Print("Enter your second sentence: ")
	line2, _ = reader.ReadString('\n')
	line2 = strings.TrimSpace(line2)
	fmt.Print("Enter your third sentence: ")
	line3, _ = reader.ReadString('\n')
	line3 = strings.TrimSpace(line3)
	fmt.Print("Enter your fourth sentence: ")
	line4, _ = reader.ReadString('\n')
	line4 = strings.TrimSpace(line4)

	// display sentences in opposite order
	fmt.Println(line4)
	fmt.Println(line3)
	fmt.Println(line2)
	fmt.Println(line1)

}
