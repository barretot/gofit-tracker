package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"sync"
)

type Workouts struct {
	Reps   int `json:"reps"`
	Weight int `json:"weight"`
}

type User struct {
	Email    string     `json:"email"`
	Workouts []Workouts `json:"workouts"`
}

type UserScore struct {
	Email string
	Score int
}

func GetScore(user User, wg *sync.WaitGroup, ctx context.Context, ch chan<- UserScore) {
	defer wg.Done()

	var totalScore int

	for _, workout := range user.Workouts {
		score := workout.Reps * workout.Weight

		totalScore += score
	}

	ch <- UserScore{
		Email: user.Email,
		Score: totalScore,
	}
}

func main() {
	data, err := os.ReadFile("workouts.json")
	if err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		return
	}

	var users []User
	var wg sync.WaitGroup
	var ctx context.Context

	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Println("Erro ao fazer unmarshal:", err)
		return
	}

	ch := make(chan UserScore, len(users))
	wg.Add(len(users))

	for _, user := range users {
		go GetScore(user, &wg, ctx, ch)
	}

	wg.Wait()
	close(ch)

	var scores []UserScore

	for value := range ch {
		scores = append(scores, value)
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	for i := 0; i < 3 && i < len(scores); i++ {
		fmt.Printf("🥇 %s: %d pontos\n", scores[i].Email, scores[i].Score)
	}
}
