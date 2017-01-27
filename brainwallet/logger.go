// logger.go - Log Writer
// Copyright (c) 2015 Kamilla Productions Uninc. Author Joonas Greis  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package brainwallet

import (
	"log"
)

// Logger Struct
type Logger struct {
	Verbose bool
}

// Log Writer function
func (l *Logger) Write(msg string) {

	if l.Verbose { log.Println(msg) }

}
