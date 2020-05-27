package main

import ( . "fmt"; "sync" )

func printWorker(str string, waitgroup *sync.WaitGroup) {
        defer waitgroup.Done()
        Println(str)
}

func main() {
        var waitgroup sync.WaitGroup

        waitgroup.Add(11)
        go printWorker("You", &waitgroup)
        go printWorker("are", &waitgroup)
        go printWorker("not", &waitgroup)
        go printWorker("seeing", &waitgroup)
        go printWorker("these", &waitgroup)
        go printWorker("strings", &waitgroup)
        go printWorker("in", &waitgroup)
        go printWorker("the", &waitgroup)
        go printWorker("proper", &waitgroup)
        go printWorker("order", &waitgroup)
        go printWorker("either.", &waitgroup)

        waitgroup.Wait()
}
