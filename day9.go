package aoc2023

func ReadOasis(path string) [][]int {
	lines := ReadFile(path)

	readings := [][]int{}
	for _, line := range lines {
		if line != "" {
			nums := ParseNumList(line)
			readings = append(readings, nums)
		}
	}
	return readings
}

func Diff(nums []int) []int {
	diffLen := len(nums) - 1
	res := make([]int, diffLen)

	for i := 0; i < diffLen; i++ {
		res[i] = nums[i+1] - nums[i]
	}
	return res
}

func IsAllZero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func Predict(nums []int) int {
	diffs := computeDiffs(nums)

	prediction := 0
	for _, diff := range diffs {
		prediction += diff[len(diff)-1]
	}
	return prediction
}

func computeDiffs(nums []int) [][]int {
	diffs := [][]int{}

	for {
		diffs = append(diffs, nums)
		if IsAllZero(nums) {
			break
		}
		nums = Diff(nums)
	}
	return diffs
}

func SumOfPredictions(readings [][]int) int {
	sum := 0
	for _, reading := range readings {
		sum += Predict(reading)
	}
	return sum
}

func history(nums []int) int {
	diffs := computeDiffs(nums)

	prediction := 0
	for i := len(diffs) - 2; i >= 0; i-- {
		prediction = diffs[i][0] - prediction
	}
	return prediction
}

func SumOfHistory(readings [][]int) int {
	sum := 0
	for _, reading := range readings {
		sum += history(reading)
	}
	return sum
}
