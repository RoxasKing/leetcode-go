package main

import (
	"container/heap"
	"sort"
)

/*
  Design a simplified version of Twitter where users can post tweets, follow/unfollow another user, and is able to see the 10 most recent tweets in the user's news feed.

  Implement the Twitter class:
    1. Twitter() Initializes your twitter object.
    2. void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user userId. Each call to this function will be made with a unique tweetId.
    3. List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user themself. Tweets must be ordered from most recent to least recent.
    4. void follow(int followerId, int followeeId) The user with ID followerId started following the user with ID followeeId.
    5. void unfollow(int followerId, int followeeId) The user with ID followerId started unfollowing the user with ID followeeId.

  Example 1:
    Input
      ["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
      [[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
    Output
      [null, null, [5], null, null, [6, 5], null, [5]]
    Explanation
      Twitter twitter = new Twitter();
      twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
      twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
      twitter.follow(1, 2);    // User 1 follows user 2.
      twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
      twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
      twitter.unfollow(1, 2);  // User 1 unfollows user 2.
      twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.

  Constraints:
    1. 1 <= userId, followerId, followeeId <= 500
    2. 0 <= tweetId <= 10^4
    3. All the tweets have unique IDs.
    4. At most 3 * 10^4 calls will be made to postTweet, getNewsFeed, follow, and unfollow.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/design-twitter
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
