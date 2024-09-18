package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name     string
	desc     string
	complete bool
}

type TaskList struct {
	tasks []Task
}

func (t *TaskList) showTasks() {
	for i, task := range t.tasks {
		fmt.Println("--------------------------------------------------------------")
		fmt.Printf("%v. %v\n- %v\n- Completada: %v \n", i+1, task.name, task.desc, task.complete)
	}
	fmt.Println("--------------------------------------------------------------")
}

func (t *TaskList) addTask(task Task) {
	t.tasks = append(t.tasks, task)
}

func (t *TaskList) completeTask(index int) {
	t.tasks[index].complete = true
}

func (t *TaskList) updateTask(index int, task Task) {
	t.tasks[index] = task
}

func (t *TaskList) deleteTask(index int) {
	t.tasks = append(t.tasks[0:index], t.tasks[index+1:]...)
}

func main() {
	TaskList := new(TaskList)

	scan := bufio.NewReader(os.Stdin)

	for {
		var opcion int

		fmt.Println("1. Agregar tarea",
			"\n2. Marcar tarea como completada",
			"\n3. Actualizar tarea",
			"\n4. Eliminar tarea",
			"\n5. Mostrar Tareas",
			"\n6. Salir",
		)

		fmt.Print("Ingresa una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Print("Ingresa el nombre de la tarea: ")
			name, _ := scan.ReadString('\n')
			fmt.Print("Ingresa la descripcion de la tarea: ")
			desc, _ := scan.ReadString('\n')
			TaskList.addTask(Task{name: name, desc: desc})
		case 2:
			var index int
			fmt.Print("Ingresa el indice de la tarea: ")
			fmt.Scanln(&index)
			TaskList.completeTask(index)
		case 3:
			var index int
			fmt.Print("Ingresa el indice de la tarea: ")
			fmt.Scanln(&index)
			fmt.Print("Ingresa el nombre de la tarea: ")
			name, _ := scan.ReadString('\n')
			fmt.Print("Ingresa la descripcion de la tarea: ")
			desc, _ := scan.ReadString('\n')
			TaskList.updateTask(index, Task{name: name, desc: desc})
		case 4:
			var index int
			fmt.Print("Ingresa el indice de la tarea: ")
			fmt.Scanln(&index)
			TaskList.deleteTask(index)
		case 5:
			TaskList.showTasks()
		case 6:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opcion invalida")
		}
	}
}
