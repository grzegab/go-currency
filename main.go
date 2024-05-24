package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"io"
	"net/http"
	"time"
)

type InputHandler struct{}
type InputData struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
type ResponseData struct {
	AmountInWords string `json:"amount_in_words"`
	TimeTook      string `json:"time_took"`
}

var ctx = context.Background()

func main() {
	s := http.Server{
		Addr:         ":8088",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  90 * time.Second,
		Handler:      InputHandler{},
	}

	err := s.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}

func (ih InputHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input InputData

	start := time.Now()
	amountText := ""

	w.Header().Set("Content-Type", "application/json")

	body := r.Body
	bodyBytes, _ := io.ReadAll(body)

	err := json.Unmarshal(bodyBytes, &input)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		errorText := "422 - Wrong input / decode problem: " + err.Error()
		w.Write([]byte(errorText))

		return
	}

	// Store in Redis ~2ms
	amountString := fmt.Sprintf("%f", input.Amount)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "top-secret",
		DB:       0, // use default DB
	})

	redisVal, err := rdb.Get(ctx, amountString).Result()
	if err == redis.Nil {
		fmt.Println("saving to Redis...")
		amountText = Translate(input.Amount)
		err = rdb.Set(ctx, amountString, amountText, 0).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("got Redis hit...")
		amountText = redisVal
	}

	//Pure count ~100us
	//amountText = Translate(input.Amount)

	elapsed := time.Since(start)

	response := ResponseData{
		AmountInWords: amountText,
		TimeTook:      elapsed.String(),
	}
	json.NewEncoder(w).Encode(response)
}
