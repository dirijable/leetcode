package medium

import "strconv"

func compress(chars []byte) int {
	read, write := 0, 0
	for read < len(chars) {
		currentChar := chars[read]
		count := 0
		for read < len(chars) && currentChar == chars[read] {
			count++
			read++
		}
		chars[write] = currentChar
		write++
		if count > 1 {
			d := strconv.AppendInt(chars[:write], int64(count), 10)
			write = len(d)
		}
	}

	return write
}

// func compress(chars []byte) int {
// 	oldLength := len(chars)
// 	indexOfNextCharForWriting := 0
// 	currentCharCount := 0
// 	idxOfNewChar := 0
// 	resultCompressLength := 0
// 	for range oldLength {
// 		currentChar := chars[idxOfNewChar]
// 		for j := idxOfNewChar; j < oldLength && currentChar == chars[j]; j++ {
// 			currentCharCount++
// 		}
// 		writedSymbols := writeCountAsSymbols(currentCharCount, resultCompressLength+1, chars)
// 		indexOfNextCharForWriting += writedSymbols + 1
// 		resultCompressLength += writedSymbols + 1
// 		idxOfNewChar += currentCharCount
// 		if idxOfNewChar >= oldLength {
// 			break
// 		}
// 		chars[indexOfNextCharForWriting] = chars[idxOfNewChar]
// 		currentCharCount = 0
// 		fmt.Println(string(chars))
// 	}
// 	return resultCompressLength
// }

// func writeCountAsSymbols(count, startIdx int, chars []byte) int {
// 	if count < 2 {
// 		return 0
// 	}
// 	lastIdx := startIdx
// 	for count > 0 {
// 		symbol := uint8(count%10) + '0'
// 		chars[lastIdx] = symbol
// 		lastIdx++
// 		count /= 10
// 	}
// 	writedSymbols := lastIdx - startIdx
// 	lastIdx--
// 	for ; lastIdx > startIdx; startIdx, lastIdx = startIdx+1, lastIdx-1 {
// 		chars[startIdx], chars[lastIdx] = chars[lastIdx], chars[startIdx]
// 	}
// 	return writedSymbols
// }
