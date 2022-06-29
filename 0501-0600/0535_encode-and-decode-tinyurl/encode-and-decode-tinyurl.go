package main

import "strconv"

// Difficulty:
// Medium

// Tags:
// Design

type Codec struct {
	a []string
}

func Constructor() Codec {
	return Codec{}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.a = append(this.a, longUrl)
	return strconv.Itoa(len(this.a) - 1)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	i, _ := strconv.Atoi(shortUrl)
	return this.a[i]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
