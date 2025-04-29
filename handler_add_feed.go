package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/combtz/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		log.Fatal("User not registered:", err)
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
	fmt.Print("Feed created successfully", feed)
	return nil
}
