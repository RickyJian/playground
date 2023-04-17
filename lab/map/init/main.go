package main

func main() {

}

type data struct {
	number int
}

func initMapWithSize(size int) {
	m := make(map[int]*data, size)
	for i := 0; i < size; i++ {
		m[i] = &data{number: i}
	}
}

func initMapWithoutSize(size int) {
	m := make(map[int]*data)
	for i := 0; i < size; i++ {
		m[i] = &data{number: i}
	}
}

func initMapWithInterface(size int) {
	m := make(map[int]interface{})
	for i := 0; i < size; i++ {
		m[i] = nil
	}
}

func initMapWithEmptyStruct(size int) {
	m := make(map[int]struct{})
	for i := 0; i < size; i++ {
		m[i] = struct{}{}
	}
}
