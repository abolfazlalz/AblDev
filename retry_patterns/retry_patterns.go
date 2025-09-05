package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fixed Delay Retry
func retryFixed(attempts int, delay time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		fmt.Println("retrying in", delay)
		time.Sleep(delay)
	}
	return err
}

// Exponential Backoff Retry
func retryExponential(attempts int, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		sleep := time.Duration(1<<i) * time.Second // 1s, 2s, 4s...
		fmt.Println("retrying in", sleep)
		time.Sleep(sleep)
	}
	return err
}

// Exponential Backoff + Jitter Retry
func retryExponentialJitter(attempts int, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		base := 1 << i
		jitter := rand.Intn(1000) // up to 1s random
		sleep := time.Duration(base*1000+jitter) * time.Millisecond
		fmt.Println("retrying in", sleep)
		time.Sleep(sleep)
	}
	return err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Fixed Delay Retry ===")
	count := 0
	err := retryFixed(3, 2*time.Second, func() error {
		count++
		if count < 3 {
			return fmt.Errorf("temporary error")
		}
		fmt.Println("✅ success on attempt", count)
		return nil
	})
	if err != nil {
		fmt.Println("final error:", err)
	}

	fmt.Println("\n=== Exponential Backoff Retry ===")
	count = 0
	err = retryExponential(4, func() error {
		count++
		if count < 3 {
			return fmt.Errorf("temporary error")
		}
		fmt.Println("✅ success on attempt", count)
		return nil
	})
	if err != nil {
		fmt.Println("final error:", err)
	}

	fmt.Println("\n=== Exponential Backoff + Jitter Retry ===")
	count = 0
	err = retryExponentialJitter(5, func() error {
		count++
		if count < 3 {
			return fmt.Errorf("temporary error")
		}
		fmt.Println("✅ success on attempt", count)
		return nil
	})
	if err != nil {
		fmt.Println("final error:", err)
	}
}