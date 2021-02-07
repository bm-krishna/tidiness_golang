package main

import (
	"context"
	"encoding/json"
	"log"
)

func Main(ctx context.Context, payloadStream []byte) ([]byte, error) {
	payload := make(map[string]interface{})
	err := json.Unmarshal(payloadStream, &payload)
	if err != nil {
		log.Fatal("Faild to unmarshal payload", err)
	}
	return nil, nil
}
