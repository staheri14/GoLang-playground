package main

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.



type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func (p *Philosopher) Eat(wg *sync.WaitGroup, allowChannel chan bool, doneChannel chan bool) {
	defer wg.Done()

	var pwg sync.WaitGroup
	for i := 0; i < 3; i++ {
		// Wait the host
		<-allowChannel

		// Pickup chopsticks in any order
		pwg.Add(2)
		go func() {
			defer pwg.Done()
			p.leftChopstick.Lock()
		}()
		go func() {
			defer pwg.Done()
			p.rightChopstick.Lock()
		}()
		pwg.Wait()

		fmt.Printf("Starting to eat %d\n", p.id)
		time.Sleep(3 * time.Second)
		fmt.Printf("Finishing eating %d\n", p.id)

		// Drop chopsticks in any order
		pwg.Add(2)
		go func() {
			defer pwg.Done()
			p.leftChopstick.Unlock()
		}()
		go func() {
			defer pwg.Done()
			p.rightChopstick.Unlock()
		}()
		pwg.Wait()

		// Inform the host
		doneChannel <- true
	}
}

func main() {
	// Create 5 chopsticks
	chopsticks := make([]*Chopstick, 0)
	for i := 0; i < 5; i++ {
		chopsticks = append(chopsticks, new(Chopstick))
	}

	// Create 5 philosophers
	philosophers := make([]Philosopher, 0)
	for i := 0; i < 5; i++ {
		p := Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
		}
		philosophers = append(philosophers, p)
	}

	var allowChannel = make(chan bool, 2)
	var doneChannel = make(chan bool, 2)

	// Philosophers start eating
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].Eat(&wg, allowChannel, doneChannel)
	}

	// The host allow maximum 2 philosophers to eat at the same time
	var c = 0
	go func() {
		for {
			select {
			case <-doneChannel:
				c--
			default:
				if c < 2 {
					c++
					allowChannel <- true
				}
			}
		}
	}()

	// Wait for all philosophers to finish eating
	wg.Wait()
}

 */