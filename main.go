// main.go - Main loop for generator
// Copyright (c) 2015 Kamilla Productions Uninc. Author Joonas Greis  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package main

import (
	"kamillaproductions.com/brainwallet-go/brainwallet"
	"flag"
	"sync"
	"time"
	"strconv"
)

// Main loop for generator
func main() {

    	startTime := time.Now()

    	// Get/Set Args
    	verbose := flag.Bool("v", false, "verbose")
    	inputFile := flag.String("i", "passphrases.txt", "input file")
    	outputFile := flag.String("o", "output.txt", "output file")

	// Parse Args
    	flag.Parse()

	// Create Logger
	logger := brainwallet.Logger{*verbose}
	logger.Write("[STATE] Initializing Brainwallet Address Generator")

	// Initialize
    	brainwallet.Init()

	// Log Params
	logger.Write("------------------------")
	logger.Write("------ Parameters ------")
	logger.Write("[PARAM] VERBOSE = " + strconv.FormatBool(*verbose))
    	logger.Write("[PARAM] STARTTIME = " + startTime.String())
 	logger.Write("[PARAM] GOMAXPROCS = " + strconv.FormatInt(int64(brainwallet.GetMaxProcs()), 10))
    	logger.Write("[PARAM] INPUT FILE = " + *inputFile)
    	logger.Write("[PARAM] OUTPUT FILE = " + *outputFile)
	logger.Write("------------------------")

    	// Create channels for goroutines
    	input, output := make(chan string), make(chan string)
    	done := make(chan int)

    	// Create WaitGroup
    	var wg sync.WaitGroup
    	wg.Add(2)

	// Create output file
    	brainwallet.CreateFile(*outputFile)

	// Start goroutines
    	logger.Write("[STATE] Starting Goroutines")
    	go brainwallet.Scanner(*inputFile, input, done, &wg)
    	go brainwallet.Generator(input, output, done, &wg)
    	go brainwallet.Writer(*outputFile, output, done, &wg)
    	go brainwallet.PrintStatistics(startTime, done, &wg)

	// Wait for finish
    	logger.Write("[STATE] Waiting To Finish")
    	wg.Wait()
    	logger.Write("[STATE] Terminating Generator")

}
