package main

import (
	"context"
	"fmt"
	"log"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		log.Fatal("No feeds registered:", err)
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			log.Fatal("Could not retrieve user with ID: ", feed.UserID)
		}
		fmt.Printf("Name: %s Url: %s User: %s \n", feed.Name, feed.Url, user.Name)
	}
	return nil
}
