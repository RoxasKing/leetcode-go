package codinginterviews

import "testing"

func TestMaxQueue(t *testing.T) {
	mq := NewMaxQueue()
	mq.Push_back(1)
	mq.Push_back(2)
	if mq.Max_value() != 2 {
		t.Errorf("Max_value() = %v, want %v", mq.Max_value(), 2)
	}
	if front := mq.Pop_front(); front != 1 {
		t.Errorf("Max_value() = %v, want %v", front, 1)
	}
	if mq.Max_value() != 2 {
		t.Errorf("Max_value() = %v, want %v", mq.Max_value(), 2)
	}
}
