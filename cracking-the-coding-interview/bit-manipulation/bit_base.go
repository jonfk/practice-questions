package main

//This method shifts 1 over by i bits, creating a value that looks like 00010000.
//By performing an AND with num,we clear all bits other than the bit at bit i.
//Finally, we compare that to 0. If that new value isnot zero,then bit i must have a 1. Otherwise, bit i is a O.
func getBit(num, i int) bool {
	return (num & (1 << i)) != 0
}

//SetBit shifts 1 over by i bits, creating a value like 00010000.
//By performing an OR with num,only the value at bit i will change.
//All other bits of the mask are zero and will not affect num.
func setBit(num, i int) int {
	return num | (1 << i)
}

//This method operates in almost the reverse of setBit.
//First,we create a number like 11101111 by creating the reverse of it (00010000) and negating it.
//Then, we perform an AND with num. This will clear the ith bit and leave the remainder unchanged.
func clearBit(num, i int) int {
	mask := ^(1 << i)
	return num & mask
}

//Clears Most Significant bit through i inclusive.
func clearBitsMSBthroughI(num, i int) int {
	mask := (1 << i) - 1
	return num & mask
}

//Clears i through 0 inclusive
func clearBitsIthrough0(num, i int) int {
	mask := ^((1 << (i + 1)) - 1)
	return num & mask
}

//This method mergesthe approaches of setBit and clearBit. First, we clear the bit at position
//i by using a mask that looks like 11101111.Then,we shift the intended value, v, left by i bits.
//This will create a number with biti equal to v and all other bits equal to 0.
//Finally, we OR these two numbers, updating the ith bit if v is 1 and leaving it as 0 otherwise.
// NOTE: v should be 1 or 0
func updateBit(num, i, v int) int {
	mask := ^(1 << i)
	return (num & mask) | (v << i)
}
