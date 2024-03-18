package main

import (
	"fmt"
	"github.com/ali-mohit/simple-fibonacci-server/internal/cache_fibo"
	"github.com/ali-mohit/simple-fibonacci-server/internal/process_fibo"
	"github.com/sirupsen/logrus"
	"math/big"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func getConsoleFibonacciCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "fibonacci [int]",
		Short: "return result of fibonacci(int)",
		Run: func(cmd *cobra.Command, args []string) {
			inMemoryCache := cache_fibo.New()
			fiboProcess := process_fibo.New(inMemoryCache)

			// Convert the first argument to integer
			n, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Error: The argument must be an integer")
				os.Exit(1)
			}

			request := process_fibo.FibonacciRequest{
				N: int64(n),
			}
			result, err := fiboProcess.ProcessFibonacciNumber(&request)
			if err != nil {
				logrus.Fatalf("process failed: %v", err)
			} else {
				fmt.Printf("%d\n", result.N)
			}
		},
	}
}

func getConsoleFibonacciV2Command() *cobra.Command {
	return &cobra.Command{
		Use:   "fibonacci-v2 [int]",
		Short: "receive (i)th fibonacci number then return (i+1)th fibonacci number",
		Run: func(cmd *cobra.Command, args []string) {
			inMemoryCache := cache_fibo.New()
			fiboProcess := process_fibo.New(inMemoryCache)

			x := big.Int{}
			_, convertResult := x.SetString(args[0], 10)
			if !convertResult {
				fmt.Println("Error: The argument must be an integer")
				os.Exit(1)
			}

			request := process_fibo.FibonacciRequestV2{
				N: &x,
			}
			result, err := fiboProcess.ProcessFibonacciNumberV2(&request)
			if err != nil {
				logrus.Fatalf("process failed: %v", err)
			} else {
				fmt.Printf("%d\n", result.N)
			}
		},
	}
}
