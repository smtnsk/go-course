package main

import ( . "fmt"; "sync" )

func printWorker(str string, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	Println(str)
}

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(1)
	go printWorker("You", &waitgroup)

	waitgroup.Add(1)
	go printWorker("are", &waitgroup)

	waitgroup.Add(1)
	go printWorker("not", &waitgroup)

	waitgroup.Add(1)
	go printWorker("seeing", &waitgroup)

	waitgroup.Add(1)
	go printWorker("these", &waitgroup)

	waitgroup.Add(1)
	go printWorker("strings", &waitgroup)

	waitgroup.Add(1)
	go printWorker("in", &waitgroup)

	waitgroup.Add(1)
	go printWorker("the", &waitgroup)

	waitgroup.Add(1)
	go printWorker("proper", &waitgroup)

	waitgroup.Add(1)
	go printWorker("order", &waitgroup)

	waitgroup.Add(1)
	go printWorker("either.", &waitgroup)

	waitgroup.Wait()
}
