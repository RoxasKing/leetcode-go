package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Priority Queue(Heap Sort)

type Twitter struct {
	users map[int]*user
	times uint
}

type user struct {
	followees map[int]struct{}
	twitters  []*twitter
}

type twitter struct {
	tId  int
	time uint
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		users: make(map[int]*user),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.users[userId]; !ok {
		this.users[userId] = &user{
			followees: make(map[int]struct{}),
		}
	}
	this.times++
	this.users[userId].twitters = append(this.users[userId].twitters, &twitter{tId: tweetId, time: this.times})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	if _, ok := this.users[userId]; !ok {
		return nil
	}

	ts := twitters{}
	for i := len(this.users[userId].twitters) - 1; i >= 0 && i >= len(this.users[userId].twitters)-10; i-- {
		ts = append(ts, this.users[userId].twitters[i])
	}
	heap.Init(&ts)

	for id := range this.users[userId].followees {
		for i := len(this.users[id].twitters) - 1; i >= 0 && i >= len(this.users[id].twitters)-10; i-- {
			heap.Push(&ts, this.users[id].twitters[i])
			if ts.Len() > 10 {
				heap.Pop(&ts)
			}
		}
	}

	sort.Slice(ts, func(i, j int) bool { return ts[i].time > ts[j].time })
	out := make([]int, 0, 10)
	for _, t := range ts {
		out = append(out, t.tId)
	}
	return out
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.users[followerId]; !ok {
		this.users[followerId] = &user{
			followees: make(map[int]struct{}),
		}
	}
	if followerId == followeeId {
		return
	}
	if _, ok := this.users[followeeId]; !ok {
		this.users[followeeId] = &user{
			followees: make(map[int]struct{}),
		}
	}
	this.users[followerId].followees[followeeId] = struct{}{}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.users[followerId]; ok {
		delete(this.users[followerId].followees, followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

type twitters []*twitter

func (t twitters) Len() int            { return len(t) }
func (t twitters) Less(i, j int) bool  { return t[i].time < t[j].time }
func (t twitters) Swap(i, j int)       { t[i], t[j] = t[j], t[i] }
func (t *twitters) Push(x interface{}) { *t = append(*t, x.(*twitter)) }
func (t *twitters) Pop() interface{} {
	top := t.Len() - 1
	out := (*t)[top]
	*t = (*t)[:top]
	return out
}
