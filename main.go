package main

func main() {
	taskStore, _ := NewTaskStore("test.json")
	taskStore.Delete(1)
	// fmt.Println(taskStore.List())
	// taskStore.Add("Testowy1")
	// fmt.Println(taskStore.List())
	// taskStore.Print()
	// taskStore.Toggle(2)
	// taskStore.Print()
}
