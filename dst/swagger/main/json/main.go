package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/podhmo/advent2016/dst/swagger/convert"
	"github.com/podhmo/advent2016/src/github"
)

func main() {
	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	var from github.AuthenticatedUser
	if err := decoder.Decode(&from); err != nil {
		log.Fatal(err)
	}

	{
		fmt.Println("------------------------------")
		fmt.Println("source:")
		fmt.Println("------------------------------")
		if err := encoder.Encode(&from); err != nil {
			log.Fatal(err)
		}
	}
	{
		to := convert.AuthenticatedUserToUser(&from)
		fmt.Println("------------------------------")
		fmt.Println("destination:")
		fmt.Println("------------------------------")
		if err := encoder.Encode(&to); err != nil {
			log.Fatal(err)
		}
	}
}
