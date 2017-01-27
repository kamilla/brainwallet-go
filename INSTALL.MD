# Installation instructions
 
## Install with go get \*PREFERRED\*
 
```bash
$ go get github.com/kamilla/brainwallet-go
$ go build github.com/kamilla/brainwallet-go
$ $GOPATH/bin/brainwallet-go -v -i passphrases.txt -o output.txt
```
Can be also installed using dummy-script, though it is not recommended!

## dummyinstaller.sh \*NOT RECOMMENDED\*

```bash
$ ./dummyinstaller.sh
```
Dummyinstaller.sh checks if the depencies and golang are installed and installs them if necessary. Then it writes the needed environment variables to ~/.profile and downloads and compiles the brainwallet with all required depencies.

Use it at your own risk! Mainly used in development purposes, it is preferred to use manual installation.

# Manual Installation 

\# Steps 

**\#1** Install Golang and dependencies (libgmp-dev) \n  
**\#2** Declare environment variables (GOROOT, GOPATH, PATH)  
**\#3** Use 'go get' and/or 'go build' to automatically download dependencies and compile the brainwallet  

# Command Listing

```bash
$ sudo apt-get install libgmp-dev											// Install depencies (gmp-dev autoconf autogen libtool)
$ wget https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz		// Download Golang binaries
$ tar -C /usr/local -xzf go1.5.linux-amd64.tar.gz							// Unpack Golang binaries to /usr/local
$ nano ~/.profile															// Edit .profile and add exports for Golang
	-> export GOROOT=/usr/local/go											// Golang installation directory
	-> export GOPATH=~/workspace/go											// We want to use local workspace for our project
	-> export PATH=$PATH:$GOROOT/bin:$GOPATH/bin							// Set the PATH environment variable
$ source ~/.profile															// Reload .profile settings
$ mkdir -p workspace/go														// Create our workspace directory
$ wget http://kamillaproductions.com/brainwallet/brainwallet.tar.gz			// Download Brainwallet-go package
$ tar -C ~/workspace/go -xzf brainwallet.tar.gz								// Unpack Brainwallet-go to our workspace
$ go get kamillaproductions.com/brainwallet-go								// Use go get to automatically download and compile depencies
$ brainwallet-go -v -i passphrases.txt -o output.txt						// Run Brainwallet-go!
```
