package main

// 3.9 Параллелизм ч.2
// https://stepik.org/submissions/1337641?unit=329291

func calculator(argements <-chan int, done <-chan struct{}) <-chan int {
	out := make(chan int)
	sum := 0
	go func() {
		for {
			select {
			case a := <- argements:
				sum += a
			case <- done:
				out <- sum
				close(out)
				return
			}
		}
	}()
	return out
}

func main() {

}
