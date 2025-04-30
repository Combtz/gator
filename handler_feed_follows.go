package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/combtz/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedsByUrl(context.Background(), url)
	if err != nil {
		log.Fatal("Failed to fetch feed")
	}
	arg := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	s.db.CreateFeedFollow(context.Background(), arg)
	fmt.Printf("Feed: %s Followed by: %s \n", feed.Name, user.Name)
	return nil
}

func handlerFeedFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		log.Fatal("Failed to fetch feed")
	}
	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}
	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.FeedName)
	}
	return nil
}
