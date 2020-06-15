package main

// 3.9 Параллелизм ч.2
// https://stepik.org/lesson/345547/step/13?unit=329291

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		select {
		case msg1 := <-firstChan:
			out <- msg1*msg1
		case msg2 := <-secondChan:
			out <- msg2*3
		case <-stopChan:
			close(out)
			return
		}
		close(out)
	}()
	return out
}

func main() {

}
