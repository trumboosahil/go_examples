package main

import (
	"fmt"
	"sync"
)

type Customer struct {
	mutex sync.RWMutex
	id    string
	age   int
}

func (c *Customer) UpdateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("age should be positive for customer %v", c)
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.age = age
	return nil
}

func (c *Customer) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return fmt.Sprintf("id %s, age %d", c.id, c.age)
}

func main() {
	var wg sync.WaitGroup

	customer := &Customer{id: "1234", age: 25}

	// Multiple concurrent reads
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(customer.String()) // safe concurrent reads
		}()
	}

	// Concurrent write
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := customer.UpdateAge(-1) // safe concurrent write
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}
