// https://adventofcode.com/2015/day/4
// solution of advent of code 2015, day4

package day4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

func findLowestNumberToGetMd5HashStartsWithXZeroes(input string, zeroCount int) int {
	for i := 1; i < math.MaxInt32; i++ {
		inputToHash := fmt.Sprintf("%s%d", input, i)
		hash := getMD5Hash(inputToHash)
		if strings.HasPrefix(hash, strings.Repeat("0", zeroCount)) {
			return i
		}
	}

	return -1
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
