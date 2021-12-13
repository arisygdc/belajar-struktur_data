package fibonacci

type Fib struct {
	cache map[int]float64
}

func (fib Fib) fibRec(n int) (result float64) {
	if val, ok := fib.cache[n]; ok {
		result = val
		return
	} else if n <= 2 {
		result = 1
	} else {
		result = fib.fibRec(n-1) + fib.fibRec(n-2)
	}
	fib.cache[n] = result
	return
}

func Count(n int) float64 {
	fib := Fib{
		cache: make(map[int]float64),
	}
	result := fib.fibRec(n)
	return result
}
