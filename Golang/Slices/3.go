package main

func getMessageCosts(messages []string) []float64 {
	cost := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		cost[i] = float64(len(msg)) * 0.01
	}
	return cost
}
