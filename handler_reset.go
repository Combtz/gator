package main

import (
	"context"
	"fmt"
	"log"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		log.Fatal("Failed to reset users table:", err)
	}
	fmt.Println("Users table reset")
	return nil
}
