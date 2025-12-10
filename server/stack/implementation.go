package stack

import (
	"errors"
	"fmt"
	// "os"
)

//----->>Функции стека

// Просмотр верхнего элемента
func (stck Stack) Peek() (int, error) {
	if stck.IsEmpty() {
		return 0, errors.New("стек пуст")
	} else {
		fmt.Print("--->Пикаем ")
		return stck.data[len(stck.data)-1], nil
	}
}

// Размер стека
func (stck Stack) Size() int {
	fmt.Print("--->Количество элементов в стеке: ")
	return len(stck.data)
}

// Проверка пуст ли стек
func (stck *Stack) IsEmpty() bool {
	if len(stck.data) == 0 {
		fmt.Println("--->Стек пуст!")
	}
	return len(stck.data) == 0
}

// Добавление элемента
func (stck *Stack) Push(value int) {
	fmt.Println("--->Добавление элемента со значением:", value)
	stck.data = append(stck.data, value)
}

// Снятие верхнего элемента
func (stck *Stack) Pop() (int, error) {
	if stck.IsEmpty() {
		return 0, errors.New("стек пуст")
	} else {
		fmt.Println("--->Удаление элемента")
		var deleted int = stck.data[len(stck.data)-1]
		stck.data = stck.data[0 : len(stck.data)-1]
		return deleted, nil
	}
}

// Вывод всего стека!!!
func (stck Stack) PrintStack() {
	fmt.Println("--->Начало вывода списка")
	for i := len(stck.data) - 1; i >= 0; i-- {
		fmt.Println(stck.data[i])
	}
	fmt.Println("Конец вывода списка<---")
}

// Конструктор стека
func NewStack() *Stack {
	var stck Stack
	return &stck
}

//----->>Учимся обрабатывать ошибки

// Обработка ошибок для функций стека
func ErrorHandler(res int, err error) int {
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	} else {
		return res
	}
}
