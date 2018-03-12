package main

func test(p *int) {
	go func() {
		println(p)
	}()
}

func main() {
	x := 100
	p := &x
	test(p)
}
