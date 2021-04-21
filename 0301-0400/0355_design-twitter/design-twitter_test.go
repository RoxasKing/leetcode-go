package main

import (
	"fmt"
)

func ExampleTwitter() {
	t := Constructor()
	t.PostTweet(1, 5)
	fmt.Println(t.GetNewsFeed(1))
	t.Follow(1, 2)
	t.PostTweet(2, 6)
	fmt.Println(t.GetNewsFeed(1))
	t.Unfollow(1, 2)
	fmt.Println(t.GetNewsFeed(1))
	// Output:
	// [5]
	// [6 5]
	// [5]
}
