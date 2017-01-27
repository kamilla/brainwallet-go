// scanner.go - File Scanner
// Copyright (c) 2015 Kamilla Productions Uninc. Author Joonas Greis  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package brainwallet

import (
        "log"
        "os"
	"bufio"
	"sync"
)

// File Scanner
func Scanner(inputFile string, input chan string, done chan int, wg *sync.WaitGroup) {
        defer wg.Done()
        counter := 0				// Counter for read lines

        file, err := os.Open(inputFile)		// Open file
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()
        scanner := bufio.NewScanner(file)	// Scanner

        for scanner.Scan() {
                counter += 1
                input <- scanner.Text()		// Send read line to input channel
        }

        done <- counter				// Everything is done. Send shutdown command to channel.

}

