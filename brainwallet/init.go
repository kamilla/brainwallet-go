// init.go - Generator Initialization
// Copyright (c) 2015 Kamilla Productions Uninc. Author Joonas Greis  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package brainwallet

import (
        "runtime"
)

// Initialize
func Init() {
	setMaxProcs()
}

// Set GOMAXPROCS Highest possible if not userdefined
func setMaxProcs() int {

    	maxProcs := runtime.GOMAXPROCS(0)	// Max Procs
    	CPUCores := runtime.NumCPU()		// Number of Cores
    	if maxProcs < CPUCores {
 		return maxProcs
    	}
    	runtime.GOMAXPROCS(CPUCores)
    	return CPUCores

}

// Get GOMAXPROCS value
func GetMaxProcs() int {

	return runtime.GOMAXPROCS(0)

}
