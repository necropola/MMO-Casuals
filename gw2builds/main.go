package main

import (
	"fmt"
	"gw2builds/gw2api"
	"log"
	"os"
	"time"
)

func main() {
	var (
		account  gw2api.Account
		chars    []string
		anything interface{}
		err      error
	)

	// Initialize the HTTP client with an API_KEY (read from environment)
	// and pass a Logger to enable diagnostic output
	key := os.Getenv("GW2API_CASUAL_02")
	api := gw2api.New(gw2api.WithAuth(key), gw2api.WithLogger(log.Default()))

	// Fetch Account information
	if account, err = api.Account(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account Name: %s\n", account.Name)

	// Fetch Core information for all characters
	if chars, err = api.Characters(); err != nil {
		log.Fatal(err)
	}
	for _, cname := range chars {
		core, err := api.CharacterCore(cname)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("    %s, %s, %s, %s, %d, %s, %s\n",
			core.Name, core.Race, core.Gender, core.Profession, core.Level, core.GuildID,
			time.Duration(core.Age)*time.Second)
	}

	// Free Style GW2API Wrestling
	if anything, err = api.Anything("/v2/account/masteries"); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Anything: %v\n", anything)
}
