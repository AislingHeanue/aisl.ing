#!/bin/bash

# Install Go
GO_VERSION="go1.22.4"
wget https://go.dev/dl/${GO_VERSION}.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf ${GO_VERSION}.linux-amd64.tar.gz
echo "export PATH=usr/local/go/bin" >> ~/.bash_profile
echo "export GOPATH=\$HOME/go" >> ~/.bash_profile
echo "export PATH=\$GOPATH/bin:\$GOROOT/bin:\$PATH" >> ~/.bash_profile
source ~/.bash_profile
rm ${GO_VERSION}.linux-amd64.tar.gz

vite build
