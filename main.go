// main.go

package main

func main() {
	a := App{}
	a.Initialize("root", "", "goforbeginner")

	a.Run(":8084")
}
