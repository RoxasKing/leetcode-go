package leetcode

import "sort"

/*
  设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

  postTweet(userID, tweetID): 创建一条新的推文
  getNewsFeed(userID): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
  follow(followerID, followeeID): 关注一个用户
  unfollow(followerID, followeeID): 取消关注一个用户
  示例:

  Twitter twitter = new Twitter();

  // 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
  twitter.postTweet(1, 5);

  // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
  twitter.getNewsFeed(1);

  // 用户1关注了用户2.
  twitter.follow(1, 2);

  // 用户2发送了一个新推文 (推文id = 6).
  twitter.postTweet(2, 6);

  // 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
  // 推文id6应当在推文id5之前，因为它是在5之后发送的.
  twitter.getNewsFeed(1);

  // 用户1取消关注了用户2.
  twitter.unfollow(1, 2);

  // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
  // 因为用户1已经不再关注用户2.
  twitter.getNewsFeed(1);
*/

type twitter struct {
	tID  int  // twitter's ID
	time uint // twitter's time
}

type user struct {
	followees map[int]struct{} // follower's user ID
	twitters  []*twitter       // user's twitters
}

// Twitter data struct
type Twitter struct {
	users map[int]*user // store user's info
	times uint
}

// NewTwitter Initialize your data structure here.
func NewTwitter() Twitter {
	return Twitter{
		users: make(map[int]*user),
	}
}

// PostTweet Compose a new tweet.
func (t *Twitter) PostTweet(userID int, tweetID int) {
	if _, ok := t.users[userID]; !ok {
		t.users[userID] = &user{
			followees: make(map[int]struct{}),
		}
	}
	t.times++
	t.users[userID].twitters = append(
		t.users[userID].twitters,
		&twitter{tweetID, t.times},
	)
}

func ajust(array []*twitter) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		son := i*2 + 1
		if son > len(array)-1 {
			continue
		}
		if son < len(array)-1 && array[son+1].time < array[son].time {
			son++
		}
		if array[son].time < array[i].time {
			array[son], array[i] = array[i], array[son]
		}
	}
}

// GetNewsFeed Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
func (t *Twitter) GetNewsFeed(userID int) []int {
	if _, ok := t.users[userID]; !ok {
		return nil
	}
	news := make([]*twitter, 0, 10)
	for i := len(t.users[userID].twitters) - 1; i >= 0; i-- {
		if len(news) == 10 {
			break
		}
		news = append(news, t.users[userID].twitters[i])
	}
	ajust(news)
	for fID := range t.users[userID].followees {
		if _, ok := t.users[fID]; !ok {
			continue
		}
		for i := len(t.users[fID].twitters) - 1; i >= 0; i-- {
			if len(news) < 10 {
				news = append(news, t.users[fID].twitters[i])
				ajust(news)
				continue
			}
			if t.users[fID].twitters[i].time > news[0].time {
				news[0] = t.users[fID].twitters[i]
				ajust(news)
			}
		}
	}
	sort.Slice(news, func(i, j int) bool {
		return news[i].time < news[j].time
	})
	out := make([]int, 0, 10)
	for i := range news {
		out = append(out, news[i].tID)
	}
	return out
}

// Follow Follower follows a followee. If the operation is invalid, it should be a no-op.
func (t *Twitter) Follow(followerID int, followeeID int) {
	if _, ok := t.users[followerID]; !ok {
		t.users[followerID] = &user{
			followees: make(map[int]struct{}),
		}
	}
	if followerID == followeeID {
		return
	}
	if _, ok := t.users[followeeID]; !ok {
		t.users[followeeID] = &user{
			followees: make(map[int]struct{}),
		}
	}
	t.users[followerID].followees[followeeID] = struct{}{}
}

// Unfollow Follower unfollows a followee. If the operation is invalid, it should be a no-op.
func (t *Twitter) Unfollow(followerID int, followeeID int) {
	if _, ok := t.users[followerID]; ok {
		delete(t.users[followerID].followees, followeeID)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := NewTwitter();
 * obj.PostTweet(userID,tweetID);
 * param_2 := obj.GetNewsFeed(userID);
 * obj.Follow(followerID,followeeID);
 * obj.Unfollow(followerID,followeeID);
 */
