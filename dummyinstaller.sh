#!/bin/bash
# Dummy Installer
# Use at your own risk! Requires awk

check() {
	printf 'Checking for %s:    	' "$1"
	if [ $(dpkg-query -W -f='${Status}' "$1" 2>/dev/null | grep -c "ok installed") -eq 0 ]; then
		printf '[NOT FOUND]\n'
		sudo apt-get --force-yes --yes install "$1"
	else
		printf '[OK]\n'
	fi
}

check_go() {
	printf 'Checking for Golang:    	'
	if hash go 2>/dev/null; then
		printf '[OK]\n'
	else
		printf '[NOT FOUND]\n'
		printf 'Downloading Golang binaries:\n'
		wget https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz -O go1.5.linux-amd64.tar.gz
		printf 'Extracting binaries to /usr/local/go:	'
		sudo tar -C /usr/local -xzf go1.5.linux-amd64.tar.gz
		printf '[OK]\n'
	fi
}

check libgmp-dev
check autoconf
check autogen
check libtool
check_go

printf 'Checking Golang environment variable exports in ~/.profile:	'
if grep -q 'export GOROOT=/usr/local/go' "$HOME/.profile" && grep -q 'export GOPATH=~/workspace/go' "$HOME/.profile" && grep -q 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' "$HOME/.profile"; then
	printf '[OK]\n'
else
	printf '[NOT FOUND]\n'
	printf 'Setting Golang environment variable exports to ~/.profile:	'
	printf 'export GOROOT=/usr/local/go\nexport GOPATH=~/workspace/go\nexport PATH=$PATH:$GOROOT/bin:$GOPATH/bin\n' >> $HOME/.profile
	printf '[OK]\n'
fi
printf '  declare GOROOT = %s\n  declare GOPATH = %s\n  declare PATH = %s\n' "$GOROOT" "$GOPATH" "$PATH"
printf 'Reloading ~/.profile:		'
source $HOME/.profile
printf '[OK]\n'

printf 'Checking if workspace directory exists:		'
if [ -d "$HOME/workspace/go" ]; then
	printf '[OK]\n'
else
	printf '[NOT FOUND]\n'
	printf 'Creating workspace:			'
	mkdir -p $HOME/workspace/go
	printf '[OK]\n'
fi

printf 'Extracting brainwallet source files to workspace:	'
tar -C $HOME/workspace/go -xzf brainwallet.tar.gz
printf '[OK]\n'
printf 'Downloading dependencies and compiling the brainwallet:	'
go get kamillaproductions.com/brainwallet-go
printf '[OK]\n'
printf '\nBrainwallet-go installed successfully!\n'
