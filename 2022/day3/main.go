package main

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	solvePart1()
	solvePart2()
}
