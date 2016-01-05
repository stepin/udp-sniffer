#dist tools
go get -v github.com/mitchellh/gox
go get -v github.com/tcnksm/ghr
gem install fpm

#deps management
brew install glide
go get -v github.com/shurcooL/gostatus

#checks
go get -v github.com/golang/lint/golint
go get -v github.com/fzipp/gocyclo
go get -v github.com/mvdan/interfacer/cmd/interfacer
go get -v golang.org/x/tools/cmd/goimports
go get -v golang.org/x/tools/cmd/oracle

#debuggers
go get -v github.com/mailgun/godebug
go get -v github.com/peterh/liner github.com/derekparker/delve/cmd/dlv
