package main

// Tags:
// Prefix Sum

func corpFlightBookings(bookings [][]int, n int) []int {
	out := make([]int, n+1)
	for _, b := range bookings {
		out[b[0]-1] += b[2]
		out[b[1]] -= b[2]
	}
	for i := 1; i < n; i++ {
		out[i] += out[i-1]
	}
	return out[:n]
}
