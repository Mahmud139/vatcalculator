# Golang vatcalculator package
 **a simple vat calculator package**


## Features

 - Lightweight and fast
 - Native Go implementation. No C-bindings, just pure Go
 - all time support
## Requirements
 - Golang any vertion 
## Installation
Simple install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go tool](https://golang.org/cmd/go/ "go command") from shell:

    $ go get -u github.com/Mahmud139/vatcalculator
 
## usage

    package main
    import (
    "fmt"
    "github.com/Mahmud139/vatcalculator"
    )
    func  main() {
    exclusive := vatcalculator.ExclusiveTax(150, 15)
    fmt.Println(exclusive)
    inclusive := vatcalculator.InclusiveTax(150, 20)
    fmt.Println(inclusive)
    
    }

