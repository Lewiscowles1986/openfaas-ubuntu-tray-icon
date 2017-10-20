## OpenFaaS Ubuntu Tray Icon

### thanks to 

- https://golang.org/cmd/cgo/
- https://askubuntu.com/questions/695796/view-list-of-all-available-unique-icons-with-their-names-and-thumbnail/695958
- https://github.com/perlw/bbinfo
- https://github.com/johnmccabe/openfaas-bitbar
    - config comes from this repo as does code to use it

### Running

> This is beyond go, git & other things you should have already

```sh
sudo apt-get install build-essential libappindicator-dev libappindicator3-dev gtk-3-examples
make dep
make run
```

