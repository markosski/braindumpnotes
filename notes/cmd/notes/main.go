package main

import (
	env "braindumpnotes.com/notes/pkg/env"
	interfc "braindumpnotes.com/notes/pkg/interfc"
)

func main() {
	envi := env.MakeInMemEnv()

	interfc.Run(envi)
}
