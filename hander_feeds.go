package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/combtz/gator/internal/database"
	"github.com/google/uuid"
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

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	arg := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), arg)
	if err != nil {
		log.Fatal("Feed not registered: ", err)
	}
	fmt.Println("Feed created successfully", feed)
	arg1 := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), arg1)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Println("Feed followed successfully:")
	fmt.Println(feedFollow.UserName, feedFollow.FeedName)
	return nil
}
