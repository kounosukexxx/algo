package main

func distance(nums []int) []int64 {
	// prepare
	mp := make(map[int][]int, len(nums))
	indexesInArray := make([]int, 0, len(nums))
	for i, num := range nums {
		mp[num] = append(mp[num], i)
		indexesInArray = append(indexesInArray, len(mp[num])-1)
	}

	// init ans
	ans := make([]int64, len(nums))
	for i := range ans {
		ans[i] = -1
	}

	// create ans
	for i := range indexesInArray {
		if ans[i] >= 0 {
			continue
		}
		array := mp[nums[i]]
		prefixSum := generatePrefixSum(array)
		for indexInArray, indexInNums := range array { // prefixSumは一度しか生成したくないので、一気にやる
			ans[indexInNums] = getAnsValue(indexInArray, array, prefixSum)
		}
	}

	return ans
}

func generatePrefixSum(array []int) []int {
	prefixSum := make([]int, 0, len(array)+1)
	prefixSum = append(prefixSum, 0)
	for _, v := range array {
		prefixSum = append(prefixSum, prefixSum[len(prefixSum)-1]+v)
	}

	return prefixSum
}

func getAnsValue(indexInArray int, array, prefixSum []int) int64 {
	ansValue := prefixSum[len(array)] - prefixSum[indexInArray+1] - prefixSum[indexInArray] + array[indexInArray]*(-len(array)+2*indexInArray+1)
	// log.Println(prefixSum[len(array)], prefixSum[indexInArray+1], prefixSum[indexInArray], array[indexInArray]*(-len(array)+2*indexInArray+1))
	return int64(ansValue)
	// log.Println(i, ans[i])
	// log.Println(prefixSum, indexInArray)
}
