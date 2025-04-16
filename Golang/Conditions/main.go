// package main
//
// import "fmt"
//
// func main() {
// 	messageLen := 10
// 	maxMessageLen := 20
// 	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)
//
// 	// don't touch above this line
//
// 	if messageLen <= maxMessageLen {
// 		fmt.Println("Message sent")
// 	} else {
// 		fmt.Println("Message not sent")
// 	}
// }
// package main
//
// import "fmt"
//
// func billingCost(plan string) float64 {
// 	switch plan {
// 	case "basic":
// 		return 10.0
// 	case "pro":
// 		return 20.0
// 	case "enterprise":
// 		return 50.0
// 	default:
// 		return 0.0
// 	}
// }
//
// // don't touch below this line
//
// func main() {
// 	plan := "basic"
// 	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
// 	plan = "pro"
// 	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
// 	plan = "enterprise"
// 	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
// 	plan = "free"
// 	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
// 	plan = "unknown"
// 	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
// }
