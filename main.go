package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	name      string
	completed bool
}

var tasks []Task

func addTask(task string) {
	newTask := Task{
		name:      task,
		completed: false,
	}
	tasks = append(tasks, newTask)
	fmt.Println("Task Added")
}

func listTask() {
	for i, v := range tasks {
		status := "not"
		if v.completed {
			status = "done"
		}
		fmt.Printf("%d : %s[%s]\n", i+1, v.name, status)
	}
}

func markCompleter(i int) {
	if i >= 1 && i <= len(tasks) {
		tasks[i-1].completed = true
		fmt.Println("Task Completed")
	} else {
		fmt.Printf("invalide Task")
	}
}

func editTask(i int, s string) {
	if i >= 1 && i <= len(tasks) {
		tasks[i-1].name = s
		tasks[i-1].completed = false
		fmt.Println("Task Updated")
	} else {
		fmt.Printf("invalide Task")
	}
}

func deletTask(i int) {
	if i >= 1 && i <= len(tasks) {
		tasks = append(tasks[:i-1], tasks[i:]...)
		fmt.Println("Task deleted")
	} else {
		fmt.Printf("invalide Task")
	}
}

func main() {
	var indexInput int
	var taskINput, NewTaskInput string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Start Todo Options:")
		fmt.Println("1 : ADD task")
		fmt.Println("2 : List task")
		fmt.Println("3 : Mark as completed task")
		fmt.Println("4 : Edit task")
		fmt.Println("5 : Delite task")
		fmt.Println("Exit")
		fmt.Println("Enter The Option =")
		scanner.Scan()
		input := scanner.Text()
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Number")
			continue
		}
		switch option {
		case 1:
			fmt.Println("Enter The Task :")
			scanner.Scan()
			NewTaskInput = scanner.Text()
			addTask(NewTaskInput)
		case 2:
			fmt.Println("Listing The Tasks :")
			listTask()
		case 3:
			fmt.Println("completed the Tasks")
			scanner.Scan()
			indexInput, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid Number")
				continue
			}
			markCompleter(indexInput)
		case 4:
			fmt.Println("Edit Task No :")
			scanner.Scan()
			indexInput, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid Number")
				continue
			}
			fmt.Println("enter new Task  :")
			scanner.Scan()
			taskINput = scanner.Text()

			editTask(indexInput, taskINput)
		case 5:
			fmt.Println("Delet Task No :")
			scanner.Scan()
			indexInput, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid Number")
				continue
			}
			deletTask(indexInput)
		case 6:
			fmt.Println("Exiting......")
			return
		}

	}

}
