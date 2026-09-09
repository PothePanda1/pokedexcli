package main

func main() {
	cliConfig := &config{registry: getCommands()}
	startRepl(cliConfig)
}
