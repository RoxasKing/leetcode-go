package main

// Tags:
// Bit Manipulation
func reverseBits(num uint32) uint32 {
	out := uint32(0)
	for i := 0; i <= 31; i++ {
		out = (out << 1) + (num>>i)&1
	}
	return out
}
