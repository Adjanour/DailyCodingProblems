package day09

func largestSumNonAdjacent(data []int) (prev int) {
	prev1 := 0
	prev2 := 0

	for i := len(data) - 1; i >= 0; i-- {

		// Option A: take data[i], then you can only add prev2
		take := data[i] + prev2

		// Option B: skip data[i], stick with prev1
		skip := prev1

		// Best choice at i:
		current := max(take, skip)

		// Slide the window:
		prev2 = prev1
		prev1 = current // store best sum from i to end

	}
	prev = prev1
	return
}
