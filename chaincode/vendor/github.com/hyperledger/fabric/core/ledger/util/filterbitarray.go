/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

// FilterBitArray is an array of bits based on byte unit, so 8 bits at each
// index. The array automatically increases if the set index is larger than the
// current capacity. The bit index starts at 0.
type FilterBitArray []byte

// NewFilterBitArray creates an array with the specified bit-size. This is an
// optimization to make array once for the expected capacity rather than
// using Set function to auto-increase the array.
func NewFilterBitArray(size uint) FilterBitArray {
	ba := make(FilterBitArray, (size-1)/8+1)
	return ba
}

// Capacity returns the number of bits in the FilterBitArray.
func (ba *FilterBitArray) Capacity() uint {
	return uint(len(*ba) * 8)
}

// Set assigns 1 to the specified bit-index, which is starting from 0.
// Set automatically increases the array to accommodate the bit-index.
func (ba *FilterBitArray) Set(i uint) {
	// Location of i in the array index is floor(i/8) + 1. If it exceeds the
	// current byte array, we'll make a new one large enough to include the
	// specified bit-index
	if i >= ba.Capacity() {
		array := make([]byte, i/8+1)
		copy(array, *ba)
		*ba = array
	}
	(*ba)[i/8] |= 1 << (i % 8)
}

// Unset assigns 0 the specified bit-index. If bit-index is larger than capacity,
// do nothing.
func (ba *FilterBitArray) Unset(i uint) {
	if i < ba.Capacity() {
		(*ba)[i/8] &^= 1 << (i % 8)
	}
}

// ValueAt returns the value at the specified bit-index. If bit-index is out
// of range, return 0. Note that the returned value is in byte, so it may be
// a power of 2 if not 0.
func (ba *FilterBitArray) ValueAt(i uint) byte {
	if i < ba.Capacity() {
		return (*ba)[i/8] & (1 << (i % 8))
	}
	return 0
}

// IsSet returns true if the specified bit-index is 1; false otherwise.
func (ba *FilterBitArray) IsSet(i uint) bool {
	return (ba.ValueAt(i) != 0)
}

// ToBytes returns the byte array for storage.
func (ba *FilterBitArray) ToBytes() []byte {
	return *ba
}

// FromBytes accepts a byte array.
func (ba *FilterBitArray) FromBytes(bytes []byte) {
	*ba = bytes
}
