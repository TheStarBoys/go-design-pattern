package singleton

import "testing"

func TestSingleton(t *testing.T) {
	ch := make(chan *singleton, 10)
	go func() {
		// producer
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- Instance()
		}
	}()

	// consumer
	pre := <- ch
	for s := range ch {
		if pre != s {
			t.Fatal("instance is not equal")
		}
	}
}