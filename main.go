package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/alwindoss/morse"
)

// See this page:
// http://www.nu-ware.com/NuCode%20Help/index.html?morse_code_structure_and_timing_.htm
var (
	dotLength  = 250 * time.Millisecond
	dashLength = 3 * dotLength
	wordLength = 7 * dotLength

	ledPath   = `/sys/class/leds/tpacpi::thinklight/brightness`
	onSignal  = `255`
	offSignal = `0`

	ledFile *os.File
)

func main() {
	// Open the device file
	var err error
	ledFile, err = os.OpenFile(ledPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer ledFile.Close()
	// parse arguments (if provided)
	start := 1
	onLoop := false
	if os.Args[1] == "-l" {
		start = 2
		onLoop = true
	}
	// Start the morse process
	input := strings.Join(os.Args[start:], " ")
	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}
	morseString := string(morseCode)
	fmt.Println(morseString)
	for ok := true; ok; ok = onLoop {
		for _, v := range morseString {
			turnOn(v)
			time.Sleep(dotLength)
		}
	}
}

func turnOn(char rune) {
	switch char {
	case '.':
		alt(dotLength)
	case '-':
		alt(dashLength)
	case ' ':
		time.Sleep(dashLength)
	case '/':
		time.Sleep(dotLength)
	}
}

func alt(sl time.Duration) {
	led(true)
	time.Sleep(sl)
	led(false)
}

func led(on bool) {
	signal := offSignal
	if on {
		signal = onSignal
	}
	_, err := ledFile.Write([]byte(signal))
	if err != nil {
		log.Fatal(err)
	}
}
