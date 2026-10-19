package main

func main() {
	linkedList := &LinkedList[string]{}

	linkedList.TailAppend("12")
	linkedList.TailAppend("13")

	linkedList.Print()
}
