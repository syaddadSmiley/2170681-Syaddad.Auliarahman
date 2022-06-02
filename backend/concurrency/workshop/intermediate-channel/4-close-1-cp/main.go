package main

//numberWorker digunakan untuk mengirim data ke channel input
func numberWorker(input chan int) {
	for i := 0; i < 10; i++ {
		input <- i
	}
	close(input)
	// TODO: answer here
}

func receiver(output chan int) {
	input := make(chan int, 10)
	go numberWorker(input)
	for number := range input {
		output <- number
	}
	close(output)
	// TODO: answer here
}
