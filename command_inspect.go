package main

import (
	"errors"
	"fmt"
)

func commandInspect(cliConfig *config, args ...string) error {
	name := args[0]
	pokemon, ok := cliConfig.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}
