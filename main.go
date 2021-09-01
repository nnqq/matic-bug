package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := rpc.DialContext(ctx, "https://matic-mumbai.chainstacklabs.com")
	if err != nil {
		panic(err)
	}

	client := ethclient.NewClient(conn)

	block, err := client.BlockByNumber(ctx, new(big.Int).SetUint64(18373929))
	if err != nil {
		panic(err)
	}

	// https://mumbai.polygonscan.com/block/18373929
	fmt.Println("block number        : 18373929")
	fmt.Println("expected block hash : 0x6a2564f4116b0e62c2bdbd38a8188b8d350a28b8025217376d79affc7dc23134")
	fmt.Println("actual block hash   :", block.Hash().String())
}
