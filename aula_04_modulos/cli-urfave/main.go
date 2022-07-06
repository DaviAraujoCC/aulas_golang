package main

import (
	"cli-urfave-exemplo/cmd"
	"os"
)

func main() {

	minhaCli := cmd.NewCMD()
	if err := minhaCli.Run(os.Args); err != nil {
		panic(err)
	}

}
