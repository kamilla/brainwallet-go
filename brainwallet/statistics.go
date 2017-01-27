// statistics.go - Statistics routines
// Copyright (c) 2015 Kamilla Productions Uninc. Author Joonas Greis  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package brainwallet

import (
        "time"
        "sync"
	"strconv"
)

// Print Statistics to io
func PrintStatistics(start time.Time, done chan int, wg *sync.WaitGroup) {
        defer wg.Done()

	logger := Logger{true}		// Create a new logger

        waitfordone:
        for {
                select {
                case count := <- done:	// Shutdown signal received. Print Stats.
                        minute := time.Minute
                        elapsed := time.Since(start)
                        speed := int64(float64(minute)/float64(elapsed)*float64(count))
			logger.Write("------------------------")
			logger.Write("------ Statistics ------")
                        logger.Write("[STATS] Bitcoin Addresses Created: " + strconv.FormatInt(int64(count), 10))
                        logger.Write("[STATS] Generation Time: " + strconv.FormatInt(int64(elapsed), 10))
                        logger.Write("[STATS] Generation Speed: " + strconv.FormatInt(int64(speed), 10) + " Addresses / Minute")
			logger.Write("------------------------")
                        break waitfordone	// Get out of here and quit
                }
        }
}
