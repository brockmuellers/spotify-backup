package main

import (
	"fmt"
	"log"
	"brockmuellers.com/spotify-backup/auth"
	"brockmuellers.com/spotify-backup/hello"
)

func main() {
	fmt.Println(hello.Hello())

	// try out app client

	appClient := auth.AppClient()

	msg, page, err := appClient.FeaturedPlaylists()
	if err != nil {
		log.Fatalf("couldn't get features playlists: %v", err)
	}

	fmt.Println(msg)
	for _, playlist := range page.Playlists {
		fmt.Println("  ", playlist.Name)
	}

	// try out user client

	userClient := auth.UserClient()

	user, err := userClient.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are:", user.DisplayName)
}
