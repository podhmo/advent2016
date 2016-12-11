package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/podhmo/advent2016/dst/swagger/convert"
	"github.com/podhmo/advent2016/src/github"
)

func main() {
	decoder := json.NewDecoder(os.Stdin)
	var from github.AuthenticatedUser
	if err := decoder.Decode(&from); err != nil {
		log.Fatal(err)
	}

	{
		fmt.Println("------------------------------")
		fmt.Println("source:")
		fmt.Println("------------------------------")
		spew.Dump(from)
	}
	{
		to := convert.AuthenticatedUserToUser(&from)
		fmt.Println("------------------------------")
		fmt.Println("destination:")
		fmt.Println("------------------------------")
		spew.Dump(to)
	}
}
