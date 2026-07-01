package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	
	"github.com/BurntSushi/toml"
)

func main() {
	value := make(map[string]any)

	tomlText, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not read from stdin:", err.Error())
		os.Exit(1)
	}

	err = toml.Unmarshal(tomlText, &value)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not unmarshal toml:", err.Error())
		os.Exit(1)
	}

	jsonText, err := json.Marshal(value)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not marshal json:", err.Error())
		os.Exit(1)
	}

	os.Stdout.Write(jsonText)
}
