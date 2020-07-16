package main

import "fmt"

func main() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	start := float64(0)

	for i := 0; i < len(scores); i++ {
		start = start + float64(scores[i])
	}

	fmt.Print(start / float64(len(scores)))

	// var goodScores []int

	// for _, score := range scores {
	// 	if score >= 90 {
	// 		goodScores = append(goodScores, score)
	// 	}
	// }
	// fmt.Println(goodScores)
}
