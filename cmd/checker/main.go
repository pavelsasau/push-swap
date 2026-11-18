package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"push-swap/stack"
)

func main() {
	args := os.Args[1:] // получаем числа из терминала

	if len(args) == 0 { // если аргументов нет
		return
	}

	text := strings.Join(args, " ") // объединяем аргументы
	parts := strings.Fields(text)   // разделяем числа

	if len(parts) == 0 { // если чисел нет
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	a := []int{} // создаём стек A
	b := []int{} // создаём пустой стек B

	for i := 0; i < len(parts); i++ { // перебираем числа
		num, err := strconv.Atoi(parts[i]) // переводим текст в число

		if err != nil { // если значение неправильное
			fmt.Fprintln(os.Stderr, "Error")
			return
		}

		for j := 0; j < len(a); j++ { // проверяем повторения
			if a[j] == num {
				fmt.Fprintln(os.Stderr, "Error")
				return
			}
		}

		a = append(a, num) // добавляем число в A
	}

	scanner := bufio.NewScanner(os.Stdin) // читаем команды из терминала

	for scanner.Scan() { // читаем каждую команду
		command := scanner.Text() // получаем название команды

		switch command { // выбираем нужную операцию
		case "sa":
			stack.Sa(a)
		case "sb":
			stack.Sb(b)
		case "ss":
			stack.Ss(a, b)
		case "pa":
			stack.Pa(&a, &b)
		case "pb":
			stack.Pb(&a, &b)
		case "ra":
			stack.Ra(&a)
		case "rb":
			stack.Rb(&b)
		case "rr":
			stack.Rr(&a, &b)
		case "rra":
			stack.Rra(&a)
		case "rrb":
			stack.Rrb(&b)
		case "rrr":
			stack.Rrr(&a, &b)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}

	if scanner.Err() != nil { // если ошибка чтения команд
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	if len(b) != 0 { // если стек B не пустой
		fmt.Println("KO")
		return
	}

	for i := 0; i < len(a)-1; i++ { // проверяем порядок чисел
		if a[i] > a[i+1] { // если числа стоят неправильно
			fmt.Println("KO")
			return
		}
	}

	fmt.Println("OK") // все числа отсортированы, B пустой
}
