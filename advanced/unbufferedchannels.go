func unbuffered() {
	ch := make(chan int) // Creates an unbuffered channel

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println(<-ch) // Receives from channel after 3 seconds
		fmt.Println("3 second Goroutine finished")
	}()

	ch <- 1 // Sends 1 to channel — BLOCKS here until someone receives

	fmt.Println("End of program")
}

func example1() {
	done := make(chan bool)

	go func() {
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
		fmt.Println("Done working")
		done <- true // Signal completion
	}()

	<-done // Wait for the signal
	fmt.Println("Worker finished, main continues")
}

func example2() {
	ping := make(chan string)
	pong := make(chan string)

	// Player 1
	go func() {
		msg := <-ping // Wait for ping
		fmt.Println("Player 1 received:", msg)
		pong <- "pong" // Send pong back
	}()

	// Player 2 (main goroutine)
	ping <- "ping" // Send ping
	msg := <-pong  // Wait for pong
	fmt.Println("Player 2 received:", msg)
}

func example3() {
	jobs := make(chan int)
	results := make(chan int)

	// Worker goroutine
	go func() {
		for job := range jobs {
			results <- job * 2 // Process and send result
		}
	}()

	// Send jobs
	jobs <- 5
	fmt.Println("Result:", <-results) // Output: 10

	jobs <- 10
	fmt.Println("Result:", <-results) // Output: 20

	close(jobs)
}

func example4() {
	ch := make(chan int)

	// This would deadlock without the goroutine!
	// ch <- 1  // DEADLOCK - no one to receive

	go func() {
		ch <- 1 // This works because main is waiting to receive
	}()

	value := <-ch // Main waits here
	fmt.Println(value)
}