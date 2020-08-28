package main

import (
	"fmt"
	"testing"
)

func TestDesignTwitter(t *testing.T) {
	T := NewTwitter()
	T.PostTweet(1, 5)
	T.Follow(1, 1)
	fmt.Println(T.GetNewsFeed(1))
}
