package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/combtz/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFeedFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		log.Fatal("Failed to fetch user")
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
