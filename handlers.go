package main

import(
	"errors"
	"fmt"
	"github.com/AgoCodeBro/gator/internal/config"
	"context"
	"github.com/AgoCodeBro/gator/internal/database"
	"github.com/google/uuid"
	"time"
	"database/sql"
)

func handlerLogin(s *state, cmd command) error{
	if len(cmd.Args) == 0 {
		return errors.New("Login expects a username as an argument")
	}
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Failed to find user %v: %v", cmd.Args[0], err)
	}

	s.cfg.CurrentUserName = cmd.Args[0]
	err = config.SetUser(s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Failed to set the User: %v", err)
	}

	fmt.Printf("%v has been set as the current user\n", s.cfg.CurrentUserName)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Register expects a username as an argument")
	}

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, cmd.Args[0])
	if err == nil {
		return fmt.Errorf("User %v already exists", cmd.Args[0])
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("Failed to register %v: %v", cmd.Args[0], err)
	}

	args := database.CreateUserParams { 
		ID 		  : uuid.New(),
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
		Name 	  :  cmd.Args[0],
	}
	
	_, err = s.db.CreateUser(ctx, args)
	if err != nil {
return fmt.Errorf("Error while adding %v to the database: %v", cmd.Args[0], err)
	}
	
	s.cfg.CurrentUserName = cmd.Args[0]
	err = config.SetUser(s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Failed to set the User: %v", err)
	}

	fmt.Printf("%v has been set as the current user\n", s.cfg.CurrentUserName)

	return nil
	
}

func handlerUsers(s *state, cmd command) error {
	names, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to get the list of users:\n", err)
	}

	for _, name := range names {
		if name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", name)
		} else {
			fmt.Printf("* %v\n", name)
		}
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("Error occured while trying to reset the users table: %v", err)
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()

	feed, err := fetchFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("Error occured while fetching feed: %v\n", err)
	}

	fmt.Println(*feed)

	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Add feed expects a feed name and a url as arguments")
	}

	ctx := context.Background()

	feedArgs := database.CreateFeedParams{
		ID 		  : uuid.New(),
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
		Name 	  : cmd.Args[0],
		Url 	  : cmd.Args[1],
		UserID 	  : user.ID,
	}

	_, err := s.db.CreateFeed(ctx, feedArgs)
	if err != nil {
		return fmt.Errorf("Failed to add the feed: %v", err)
	}
	
	fmt.Printf("Added feed\n Name: %v\n URL: %v\n Created At: %v\n Updated At: %v\n User ID: %v\n Feed ID: %v\n",
		feedArgs.Name, 
		feedArgs.Url,
		feedArgs.CreatedAt,
		feedArgs.UpdatedAt,
		feedArgs.UserID,
		feedArgs.ID,
	)
	
	followArgs := database.CreateFeedFollowParams {
		ID        : uuid.New(),
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
		UserID    : user.ID,
		FeedID    : feedArgs.ID,
	}
	
	_, err = s.db.CreateFeedFollow(ctx, followArgs)
	if err != nil {
		return fmt.Errorf("Failed to follow the feed: %v", err)
	}

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()
	
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get the feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %v\nUrl: %v\nUser: %v\n", feed.Name, feed.Url, feed.UserName)
		fmt.Println("____________________________________")
	}

	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Follow command expects a url as an argument")
	}

	url := cmd.Args[0]
	ctx := context.Background()

	feed, err := s.db.GetFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("Falied to find that feed: %v", err)
	}

	followArgs := database.CreateFeedFollowParams{
		ID        : uuid.New(),
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
		UserID    : user.ID,
		FeedID    : feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(ctx, followArgs)
	if err != nil { 
		return fmt.Errorf("Failed to follow the feed: %v", err)
	}
	
	fmt.Printf("User: %v\nFeed: %v\n", feedFollow.UserName, feedFollow.Name)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)	
	if err != nil {
		return fmt.Errorf("Failed to get the feed for %v: %v", user.Name, err)
	}

	fmt.Println("Following:")
	for _, feed := range feeds {
		fmt.Println("- ", feed.Name)
	}

	return nil
}
