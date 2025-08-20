package main

import (
	"fmt"
	"strings"
	"context"
	"database/sql"
	"time"
	"github.com/AgoCodeBro/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	feedInfo, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to get the next feed: %v", err)
	}
	
	err = s.db.MarkFeedFetched(context.Background(), feedInfo.ID)
	if err != nil {
		return fmt.Errorf("Failed to update feed: %v", err)
	}

	feed, err := fetchFeed(context.Background(), feedInfo.Url)
	if err != nil {
		return fmt.Errorf("Failed to fetch feed: %v", err)
	}
	
	fmt.Printf("Feed Name: %v\n", feed.Channel.Title)

	for _, item := range feed.Channel.Item {
		var nullableDescription sql.NullString
		if len(item.Description) > 0 {
			nullableDescription = sql.NullString {
				String : item.Description,
				Valid  : true,
			}
		} else {
			nullableDescription = sql.NullString {
				String : "",
				Valid  : false,
			}
		}
		
		var nullablePublishedAt sql.NullTime
		pubTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Printf("Failed to format date: %v", err)
			nullablePublishedAt = sql.NullTime {Valid : false}
		} else {
			nullablePublishedAt = sql.NullTime {
				Time  : pubTime,
				Valid : true,
			}
		}

		postArgs := database.CreatePostParams{
			ID          : uuid.New(),
			CreatedAt   : time.Now(),
			UpdatedAt   : time.Now(),
			Title       : item.Title,
			Url         : item.Link,
			Description : nullableDescription,
			PublishedAt : nullablePublishedAt,
			FeedID      : feedInfo.ID,
		}

		_, err = s.db.CreatePost(context.Background(), postArgs)
		if err !=nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			continue
		} else if err != nil {
			return err
		}

	}

	return nil
}

