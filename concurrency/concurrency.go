package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func CountDown(n int) {
	for i := n; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func TickTock(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Tick", time.Now())
		time.Sleep(1 * time.Second)
	}
}

func callWithGoroutine() {
	go TickTock(10)
	go CountDown(10)
	time.Sleep(11 * time.Second)
}

func TickTockV2(done chan bool, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Tick", time.Now())
		time.Sleep(1 * time.Second)
	}
	done <- true
}

func CountDownV2(done chan bool, n int) {
	for i := n; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	done <- true
}

func callWithGoroutineV2() {
	done := make(chan bool)
	go TickTockV2(done, 10)
	go CountDownV2(done, 10)
	<-done
	<-done
}

func sequentialTickTock() {
	TickTock(10)
	fmt.Println("TickTock Done")
	CountDown(10)
	fmt.Println("CountDown Done")
}

func bufferredChannel() {
	ch := make(chan int, 2)
	ch <- 1
	fmt.Println("Add 1")
	ch <- 2
	fmt.Println("Add 2")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func sender(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Sender: ", i)
	}
}

func receiver(ch chan int) {
	for {
		fmt.Println("Receiver: ", <-ch)
	}
	fmt.Println("Receiver: Done")
}

func senderV2(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Sender: ", i)
	}
	close(ch)
}

func receiverV2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Receiver: ", v)
	}
	fmt.Println("Receiver: Done")
}

func receiverV3(ch chan int) {
	for v := range ch {
		fmt.Println("Receiver: ", v)
	}
	fmt.Println("Receiver: Done")
}

func sendAndReceive() {
	ch := make(chan int)
	go sender(ch)
	go receiver(ch)
	time.Sleep(1 * time.Second)
}

func sendAndReceiveV2() {
	ch := make(chan int)
	go senderV2(ch)
	go receiverV2(ch)
	time.Sleep(1 * time.Second)
}

func sendAndReceiveV3() {
	ch := make(chan int)
	go senderV2(ch)
	go receiverV3(ch)
	time.Sleep(1 * time.Second)
}

func multiplexWithSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer fmt.Println("Reciever: DONE")
		for {
			select {
			case v := <-ch1:
				fmt.Println("Channel 1: ", v)
			case v := <-ch2:
				fmt.Println("Channel 2: ", v)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch1 <- i
		ch2 <- i
	}
	time.Sleep(1 * time.Second)
}

func multiplexWithSelectV2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan bool)

	go func() {
		defer fmt.Println("Reciever: DONE")
		for {
			select {
			case v := <-ch1:
				fmt.Println("Channel 1: ", v)
			case v := <-ch2:
				fmt.Println("Channel 2: ", v)
			case <-done:
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch1 <- i
		ch2 <- i
	}
	done <- true
	time.Sleep(1 * time.Second)
}

func multiplexWithSelectV3() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func syncWithMutex() {
	var sharedCounter int
	var mutex sync.Mutex
	increment := func() {
		mutex.Lock()
		defer mutex.Unlock()
		sharedCounter++
		fmt.Printf("Incrementing: %d\n", sharedCounter)
	}

	for i := 0; i < 1000; i++ {
		go increment()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Counter: ", sharedCounter)
}

func syncDoneWithWaitGroup() {
	var sharedCounter int
	var mutex sync.Mutex
	var wg sync.WaitGroup
	increment := func() {
		mutex.Lock()
		defer mutex.Unlock()
		defer wg.Done()
		sharedCounter++
		fmt.Printf("Incrementing: %d\n", sharedCounter)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println("Counter: ", sharedCounter)
}

func sequentialRequest() {
	fmt.Println("Sequential Request")
	start := time.Now()
	for i := 0; i < 10; i++ {
		req, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
		if err != nil {
			panic(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		fmt.Println("DONE with", resp.StatusCode)
	}
	fmt.Println("use: ", time.Since(start))
}

func concurrentRequest() {
	fmt.Println("Sequential Request")
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
			if err != nil {
				panic(err)
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				panic(err)
			}
			fmt.Println("DONE with", resp.StatusCode)
		}()
	}
	wg.Wait()
	fmt.Println("use: ", time.Since(start))
}

func contextWithCancel() {
	process := func(ctx context.Context) {
		for i := 1; i <= 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Context Done")
				return
			default:
				fmt.Println(i)
				time.Sleep(1 * time.Second)
			}
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer time.Sleep(1 * time.Second)
	defer cancel()
	go process(ctx)
	time.Sleep(5 * time.Second)
}

func contextWithTimeout() {
	process := func(ctx context.Context) {
		for i := 1; i <= 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Context Done")
				return
			default:
				fmt.Println(i)
				time.Sleep(1 * time.Second)
			}
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer time.Sleep(1 * time.Second)
	defer cancel()
	go process(ctx)
	time.Sleep(10 * time.Second)
}

func contextWithValue() {
	type key string
	const countLimitKey key = "count-limit"

	process := func(ctx context.Context) {
		countLimit, ok := ctx.Value(countLimitKey).(int)
		if !ok {
			countLimit = 10
		}
		for i := 1; i <= countLimit; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Context Done")
				return
			default:
				fmt.Println(i)
				time.Sleep(1 * time.Second)
			}
		}
	}
	ctx := context.WithValue(context.Background(), countLimitKey, 2)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	go process(ctx)
	time.Sleep(10 * time.Second)
}

func main() {
	// sequentialTickTock()
	// callWithGoroutine()
	// callWithGoroutineV2()
	// bufferredChannel()
	// sendAndReceive()
	// sendAndReceiveV2()
	// sendAndReceiveV3()
	// multiplexWithSelect()
	// multiplexWithSelectV2()
	// multiplexWithSelectV3()
	// syncWithMutex()
	// syncDoneWithWaitGroup()
	// sequentialRequest()
	// concurrentRequest()
	// cancelWithContext()
	// contextWithTimeout()
	// contextWithValue()
}
