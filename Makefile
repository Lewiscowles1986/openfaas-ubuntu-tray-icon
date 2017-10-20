GOCOMMAND=go
GOFLAGS=build
prefix=/usr/local

all:
	$(GOCOMMAND) $(GOFLAGS) main.go

dep:
	go get "github.com/doxxan/appindicator"
	go get "github.com/doxxan/appindicator/gtk-extensions/gotk3"
	go get "github.com/conformal/gotk3/gtk"
	go get "github.com/mitchellh/go-homedir"
	go get "github.com/mitchellh/mapstructure"
	go get "github.com/spf13/viper"
	go get "github.com/prometheus/common/model"
	go get "github.com/prometheus/client_golang/api/prometheus"
	go get "gonum.org/v1/plot"

run:
	go run main.go

install:
	install -m 0755 main $(prefix)/bin

uninstall:
	rm -f $(prefix)/bin/main
clean:
	rm -rf main
