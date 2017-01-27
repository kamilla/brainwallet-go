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

# Usage

## Run from command line
```bash
$ brainwallet-go
```

## Command line arguments
    |     |     | 
:-: | :-: | :-: | :-: 
-h | --help | prints help |
-v (bool) | -verbose (bool) | verbose
-i \<path\> | --input \<path\> | input file
-o \<path\> | --output \<path\> | output file

## Example
```bash
$ brainwallet-go -v -i passphrases.txt -o generatedaddresses.txt
```

# Files
    |     |     | 
:-- | :-- | :-- | :-- 
**INSTALL.MD** |  | *Installation instructions* |
**LICENSE** |  | *MIT License* |
**README.MD** |  | *This document* |
**dummyinstaller.sh** |  | *Installer script* |
**main.go** |  | *Main loop for generator* |
**passphrases.txt** |  | *Sample passphrase list* |
brainwallet/ |  | 
 | **generator.go** | *Bitcoin Address Generator routines*  | 
 | **init.go** | *Initialization functions* | 
 | **logger.go** | *Log writer* | 
 | **scanner.go** | *File scanner routines* | 
 | **statistics.go** | *Statistic routines* | 
 | **writer.go** | *Writer routines* | 

## Dependencies
    |     |     | 
:-- | :-- | :-- | :-- 
**[github.com/haltingstate/secp256k1-go](http://github.com/haltingstate/secp256k1-go)** |  | *Copyright (c) 2013 Pieter Wuille. MIT License* |
**[github.com/jbenet/go-base58/](http://github.com/jbenet/go-base58/)** |  | *Copyright (c) 2013 Conformal Systems LLC. ISC License. Modified by Juan Benet (juan@benet.ai)* |
**[golang.org/x/crypto/ripemd160/](http://golang.org/x/crypto/ripemd160/)** |  | *Copyright 2010 The Go Authors* |

# Installation

Please see [INSTALL](INSTALL.MD).

# Updating from Sources
 ---------------------

If you want to rebuild brainwallet from sources, just run "go get github.com/kamilla/brainwallet-go" and it does the work for you.

# Licensing

This software is licensed under [MIT license](LICENSE).  
Copyright (c) 2017 Kamilla Productions Uninc.

# Cryptographic Software Notice

This distribution may include software that has been designed for use with cryptographic software.  The country in which you currently reside may have restrictions on the import, possession, use, and/or re-export to another country, of encryption software.  BEFORE using any encryption software, please check your country's laws, regulations and policies concerning the import, possession, or use, and re-export of encryption software, to see if this is permitted.  See <http://www.wassenaar.org/> for more information.
