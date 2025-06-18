package main

import (
	"testing"
)

// task   -> Struct
// tasks  -> list of task
// getNextId()

//type task struct {
//	id int
//	taskName string
//	completed bool
//}

func TestGenerateID(t *testing.T) {
	getID := generateID()
	// a := getNextId() // why this line is giving me error
	id1 := getID()
	id2 := getID()

	if id1 != 1 || id2 != 2 {
		t.Errorf("Expected ids 1 and 2, got %d and %d", id1, id2)
	}
}

func TestListTasks(t *testing.T) {
	// will have to think about Capturing the Output to somewhere and then
	// doing the Comparision
}

func TestMarkTaskAsCompleted(t *testing.T) {
	tasks := []task{}
	var completed map[int]bool
	completed = make(map[int]bool)
	currTask := task{
		1, "demotask", false,
	}
	tasks = addTask(currTask, tasks)
	markTaskAsCompleted(2, tasks, completed) // this will raise an error
	markTaskAsCompleted(1, tasks, completed) // testcase passed
	if !completed[1] {
		t.Errorf("expected id with %d as completed,found not Completed", currTask.id)
	}
}

func TestAddTask(t *testing.T) {
	tasks := []task{}
	initialSize := len(tasks)
	currTask := task{
		1, "demotask", false,
	}

	tasks = addTask(currTask, tasks)
	afterSize := len(tasks)

	if initialSize+1 != afterSize {
		t.Errorf("expected %d, got %d", initialSize+1, afterSize)
	}
}
