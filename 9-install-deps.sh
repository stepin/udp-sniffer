#dist tools
go get -v -u github.com/mitchellh/gox
go get -v -u github.com/tcnksm/ghr
#gem install fpm

#deps management
brew install glide
go get -v -u github.com/shurcooL/gostatus

#checks
go get -v -u github.com/golang/lint/golint
go get -v -u github.com/fzipp/gocyclo
go get -v -u github.com/mvdan/interfacer/cmd/interfacer
go get -v -u golang.org/x/tools/cmd/goimports
go get -v -u github.com/kisielk/errcheck
go get -v -u github.com/alecthomas/gometalinter
gometalinter --install

#debuggers
go get -v -u github.com/mailgun/godebug
go get -v -u github.com/peterh/liner github.com/derekparker/delve/cmd/dlv

#ide
go get -v -u golang.org/x/tools/cmd/gorename
go get -v -u github.com/rogpeppe/godef
go get -v -u github.com/jstemmer/gotags
go get -v -u github.com/nsf/gocode
go get -v -u golang.org/x/tools/cmd/guru
#apm install go-rename
#apm install go-plus
