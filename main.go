package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	const url = "wss://cooker-eth:cooker0910@apis-sj.ankr.com/wss/135c6de4c5e74a6e839dd9c9ee0a4425/cab324a7a627360ac56b0a03af261d11/eth/fast/main" // url string

	client, err := ethclient.Dial(url)

	if err != nil {
		panic(err)
	}

	ch := make(chan *types.Header, 1024)
	sub, err := client.SubscribeNewHead(context.Background(), ch)

	if err != nil {
		panic(err)
	}

	fmt.Println("---subscribe-----")

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("---unsubscribe-----")
		sub.Unsubscribe()
	}()

	go func() {
		for c := range ch {
			fmt.Println(c.Number)
		}
	}()

	<-sub.Err()

}
