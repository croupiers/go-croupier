package croupier

import "fmt"

func ExampleSampleFloat64() {
	uniform := NewUniform()

	sample := SampleFloat64(5, uniform)

	for _, s := range sample {
		fmt.Println(s)
	}
}

func ExampleMapDistributionFloat64() {
	uniform := NewUniform()

	constant := MapDistributionFloat64(uniform, func(_ float64) float64 {
		return 5
	})

	for i := 0; i < 10; i++ {
		fmt.Println(<-constant)
	}

	//Output: 5
	//5
	//5
	//5
	//5
	//5
	//5
	//5
	//5
	//5
}

func ExampleNewUniform() {
	uniform := NewUniform()

	for {
		if num := <-uniform; num > 0.5 {
			break
		}

	}
}

func ExampleNewUniformInRange() {
	uniform := NewUniformInRange(5, 10)

	for {
		if num := <-uniform; num > 7 {
			break
		}
	}
}

func ExampleNewNormal() {
	normal := NewNormal()

	for {
		if num := <-normal; num > 0 {
			break
		}
	}
}

func ExampleNewCustomNormal() {
	normal := NewCustomNormal(4, 8)

	for {
		if num := <-normal; num > 0 {
			break
		}
	}
}
