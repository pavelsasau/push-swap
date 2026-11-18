package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"push-swap/stack"
)

func main() {
	args := os.Args[1:] // получаем аргументы пользователя

	if len(args) == 0 { // если аргументов нет
		return
	}

	text := strings.Join(args, " ") // объединяем аргументы
	parts := strings.Fields(text)   // разделяем текст на числа

	if len(parts) == 0 { // если чисел нет
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	a := []int{} // создаём пустой стек A

	for i := 0; i < len(parts); i++ { // перебираем все числа
		num, err := strconv.Atoi(parts[i]) // переводим текст в число

		if err != nil { // если введено не число
			fmt.Fprintln(os.Stderr, "Error")
			return
		}

		for j := 0; j < len(a); j++ { // проверяем повторения
			if a[j] == num { // если число уже есть
				fmt.Fprintln(os.Stderr, "Error")
				return
			}
		}

		a = append(a, num) // добавляем число в A
	}

	if len(a) == 2 { // если два числа
		if a[0] > a[1] {
			stack.Sa(a) // меняем числа местами
			fmt.Println("sa")
		}
	}

	if len(a) == 3 { // если три числа
		sortThree(&a)
	}

	if len(a) == 4 || len(a) == 5 { // если четыре или пять чисел
		sortFive(&a)
	}
}

// SORT THREE
func sortThree(a *[]int) { // сортируем три числа
	if len(*a) != 3 {
		return
	}

	x := (*a)[0] // первое число
	y := (*a)[1] // второе число
	z := (*a)[2] // третье число

	if x < y && y < z { // уже отсортировано
		return
	}

	if x > y && x < z { // например 2 1 3
		stack.Sa(*a)
		fmt.Println("sa")

	} else if x > y && y > z { // например 3 2 1
		stack.Sa(*a)
		fmt.Println("sa")
		stack.Rra(a)
		fmt.Println("rra")

	} else if x > z && y < z { // например 3 1 2
		stack.Ra(a)
		fmt.Println("ra")

	} else if x < y && x > z { // например 2 3 1
		stack.Rra(a)
		fmt.Println("rra")

	} else if x < z && y > z { // например 1 3 2
		stack.Sa(*a)
		fmt.Println("sa")
		stack.Ra(a)
		fmt.Println("ra")
	}
}

// FIND MIN
func findMin(a []int) int { // ищем позицию минимального числа
	min := 0 // сначала считаем первое число минимальным

	for i := 1; i < len(a); i++ { // перебираем остальные числа
		if a[i] < a[min] { // если нашли число меньше
			min = i // запоминаем его позицию
		}
	}

	return min // возвращаем позицию
}

// SORT FIVE
func sortFive(a *[]int) { // сортируем четыре или пять чисел
	b := []int{} // создаём пустой стек B

	for len(*a) > 3 { // оставляем в A три числа
		min := findMin(*a) // ищем позицию минимального числа

		for min > 0 { // перемещаем минимальное число наверх
			if min <= len(*a)/2 { // если ближе к началу
				stack.Ra(a) // прокручиваем A вперёд
				fmt.Println("ra")
			} else { // если ближе к концу
				stack.Rra(a) // прокручиваем A назад
				fmt.Println("rra")
			}

			min = findMin(*a) // обновляем позицию минимума
		}

		stack.Pb(a, &b) // переносим минимальное число в B
		fmt.Println("pb")
	}

	sortThree(a) // сортируем три оставшихся числа

	for len(b) > 0 { // возвращаем числа из B
		stack.Pa(a, &b)
		fmt.Println("pa")
	}
}
