# About

This is a bitcoin brainwallet address generator library package for golang. It can be used to masscreate brainwallet bitcoin addresses. It uses Optimized C library for EC operations on curve secp256k1 and base58 encoding library extracted from btcutil. Generator is designed to be fast and efficient.

# Generation Process

Passphrase -> ECDSA Private Key -> ECDSA Public Key

  |  |  |  |  | | | |  |  |  |  |  |  |  |  |
:-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-:
  |  |  |  |  | ↱ | SHA256<br/>→<br/>*crypto/sha256* | **Hash** | RIPEMD160<br/>→<br/>*golang.org/x/crypto/ripemd160* | **Hash** | Add version byte<br/>→<br/>*(0x00 for Main Network)* | **Extented Hash** | ⬎ | |  |  |  |
**Passphrase** | → | SHA256<br/>→<br/>*crypto/sha256* | **ECDSA Private Key** | UncompressedPubkeyFromSeckey<br/>→<br/>*github.com/haltingstate/secp256k1-go* | **ECDSA Public Key** |  |  |  |  |  |  | ↓<br/>Concatenate→<br/>↑ | **25-Byte Binary Bitcoin Address** | Base58Check encoding<br/>→<br/>*github.com/jbenet/go-base58* | **Base58 Encoded Bitcoin Address**
  |  |  |  |  | ↳ | SHA256<br/>→<br/>*crypto/sha256* | **Hash** | SHA256<br/>→<br/>*crypto/sha256* | **Hash** | Crop 4 first bytes<br/>→<br/> | **Address Checksum** | ⬏ | |  |  |  |
  |  |  |  |  | | | |  |  |  |  |  | |  |  | 


 Usage
 -----

 # Run from command line

 $ brainwallet-go


 # Command line arguments

 -h	--help		prints help
 -v (bool)		verbose
 -i <path>		input file
 -o <path>		output file


 # Example
 
 $ brainwallet-go -v -i passphrases.txt -o output.txt




 Files
 -----

 bin/						// Binaries
 dummyinstaller.sh				// Installer script
 INSTALL					// Installation instructions
 passphrases.txt				// Sample passphrase list
 pkg/						// Package directory
 README						// This file
 src/
        kamillaproductions.com/			// Copyright (c) 2015 Kamilla Productions Uninc. Author Joonas Greis (joonas.greis@kamillaproductions.com)
                brainwallet-go/
                        main.go			// Main loop for generator
                        brainwallet/		// Package brainwallet
                                generator.go	// Bitcoin Address Generator routines
                                init.go		// Initialization functions
                                logger.go	// Log writer
                                scanner.go	// File scanner routines
                                statistics.go	// Statistic routines
                                writer.go	// Writer routines
	github.com/
		haltingstate/
			secp256k1-go/		// Copyright (c) 2013 Pieter Wuille. MIT License
				...
		jbenet/
			go-base58/		// Copyright (c) 2013 Conformal Systems LLC. ISC License. Modified by Juan Benet (juan@benet.ai)
				...
	golang.org/
		x/
			crypto/
				ripemd160/	// Copyright 2010 The Go Authors.




 Installation
 ------------

 Please see the file called INSTALL.



 Updating from Sources
 ---------------------

 If you want to rebuild brainwallet from sources, just run "go get kamillaproductions.com/brainwallet-go"
 and it does the work for you.



 Licensing
 ---------

 Please see the file called LICENSE.




 Cryptographic Software Notice
 -----------------------------

 This distribution may include software that has been designed for use
 with cryptographic software.  The country in which you currently reside
 may have restrictions on the import, possession, use, and/or re-export
 to another country, of encryption software.  BEFORE using any encryption
 software, please check your country's laws, regulations and policies
 concerning the import, possession, or use, and re-export of encryption
 software, to see if this is permitted.  See <http://www.wassenaar.org/>
 for more information.
