# hdwallet-export-key
export hdwallet the specified path ethereum key

# build
go get github.com/miguelmota/go-ethereum-hdwallet
go build

# running
```
./hdwallet-export-key --help
Usage of ./hdwallet-export-key:
  -mnemonic string
    	Mnemonic
  -path string
    	HD path


./hdwallet-export-key -mnemonic 'tag volcano eight thank tide danger coast health above argue embrace heavy' -path "m/44'/60'/0'/0/0"
```
