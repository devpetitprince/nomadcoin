package main

import (
	"github.com/devpetitprince/nomadcoin/cli"
	"github.com/devpetitprince/nomadcoin/db"
)

func main(){
	defer db.Close()
	cli.Start()
}
