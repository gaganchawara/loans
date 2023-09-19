package helper

func BreakAmount(amount int64, term int) []int64 {
	// Calculate the base amount for each term.
	baseAmount := amount / int64(term)

	// Calculate the remaining amount.
	remainingAmount := amount % int64(term)

	// Create a slice to store the result.
	parts := make([]int64, term)

	// Distribute the base amount evenly among terms.
	for i := 0; i < term; i++ {
		parts[i] = baseAmount
	}

	// Add the remaining amount to the last part.
	for i := 0; i < int(remainingAmount); i++ {
		parts[term-i-1]++
	}

	return parts
}
