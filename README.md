Steps to run the presentation locally: 

1. `go get -u golang.org/x/tools/present`
2. Run `present` command
3. Open your web browser and visit http://127.0.0.1:3999

Note:
If you are using multiple versions of golang, and the present command does not work
as expected, ensure the gopath is setup properly with the correct go version in use.

Example:
1. Find your go path, wherever go is installed:
   1. If you are using GVM, set the path to respective version as follows:
      1. ~/.gvm/1.19/go (If using gvm) OR
      2. export GOBIN=$GOPATH/bin
      3. export PATH=$GOBIN:$PATH
   2. If you have installed go using homebrew, ensure the bin folder is in your path:
        1. export GOPATH=$HOME/go
        2. export GOROOT="$(brew --prefix golang)/libexec"
        3. export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
2. You can add the above steps in your initialization file .bash_profile, .bashrc or .zshrc.