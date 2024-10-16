/*
Калькулятор умеет выполнять операции сложения строк, вычитания строки из строки, умножения строки на число и деления строки на число: "a" + "b", "a" - "b", "a" * b, "a" / b. Данные передаются в одну строку (смотри пример ниже). Решения, в которых каждая строка, число и арифметическая операция передаются с новой строки, считаются неверными.

Значения строк, передаваемых в выражении, выделяются двойными кавычками.

Результатом сложения двух строк является строка, состоящая из переданных.

Результатом деления строки на число n является строка в n раз короче исходной (смотри пример).

Результатом умножения строки на число n является строка, в которой переданная строка повторяется ровно n раз.

Результатом вычитания строки из строки является строка, в которой удалена переданная подстрока или сама исходная строка, если в неё нет вхождения вычитаемой строки (смотри пример).

Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более, и строки длиной не более 10 символов. Если строка, полученная в результате работы приложения, длиннее 40 символов, то в выводе после 40 символа должны стоять три точки (...).

Калькулятор умеет работать только с целыми числами.

Первым аргументом выражения, подаваемым на вход, должна быть строка. При вводе пользователем выражения вроде 3 + "hello", калькулятор должен выдать панику и прекратить свою работу.

При вводе пользователем неподходящих чисел, строк или неподдерживаемых операций (например, деление строки на строку) приложение выдаёт панику и завершает свою работу.

При вводе пользователем выражения, не соответствующего одной из вышеописанных арифметических операций, приложение выдаёт панику и завершает свою работу.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main(){
	// var text string

	// fmt.Print("Введите выражение: ")
	// fmt.Fscan(os.Stdin, &vvod)	
	// splitFunc := func(r rune) bool {
	// 	return strings.ContainsRune("*%,/%,-%,+", r)
	// }
	// nummms := strings.FieldsFunc(vvod, splitFunc)
	// fmt.Println(vvod)

	// fmt.Print("Введите выражение: ")
	// fmt.Scan(&text)
	// fmt.Println(text)


	
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')  // прием строки полностью, с пробелами
	text = strings.TrimSpace(text)

	// splitFunc := func(r rune) bool {
	// 	return strings.ContainsRune("*%,/%,+", r)
	// }
	// nums := strings.FieldsFunc(text, splitFunc)
	// _ = nums
	//nummms_2 := strings.SplitN(text, " - ", 2 )

	// fmt.Println(nummms_2)
	// fmt.Println(len(nummms_2))
	// for i := 0; i < len(nummms_2); i++{
	// 	fmt.Println(nummms_2[i])
	// }
	//s := len(text)
	//fmt.Println(utf8.RuneCountInString(text))      //узнать длину строки 

	//value = strconv.Quote(value)				    //поместить текст в ковычки
	// stringss = append(stringss, num2)    не нужно!!!!
	// nummms := strings.Split(text, " ")           //Разбить по пробелам, нам не подошло

		if strings.Contains(text, " + ") || strings.Contains(text, " - ") || strings.Contains(text, " / ") || strings.Contains(text, " * "){
			if strings.Contains(text, "+"){
				nummms := strings.Split(text, " + ")
				if len(nummms) != 2 {
					panic("Ошибка")
				}
				n1, err := strconv.Atoi(nummms[0])
				if err == nil{
					panic("Первое не может быть число")
				}
				_ = n1
				//stringss := []string{}
				t1 := strings.TrimSpace(nummms[0])
				nummms[0] = t1

				n2, err := strconv.Atoi(nummms[1])
				if err == nil{
					panic("складывать строки с числами нельзя")
				}
				_ = n2

				t2 := strings.TrimSuffix(nummms[1], "\r\n")      //убрать "\r\n"
				t2 = strings.TrimSpace(t2)          		     //убрать пробелы

				if len(nummms) == 2{
					fmt.Println(utf8.RuneCountInString(t1))
					fmt.Println(utf8.RuneCountInString(t2))
					if utf8.RuneCountInString(t1) > 12{
						panic("Слишком длинная строка!!!")
					}
					if utf8.RuneCountInString(nummms[1]) > 12{
							panic("Слишком длинная строка!!!")
					} 
					num1, err := strconv.Unquote(t2)
					if err != nil{
						panic("Ошибка ввода")
					}
					num2, err := strconv.Unquote(nummms[1])
					if err != nil{
						panic("Ошибка ввода")
					}				
					res := num1 + num2
					fmt.Println(strconv.Quote(res))
				} else {
					panic("Не верное выражение")
				}
			}else if strings.Contains(text, "-"){
				nummms := strings.Split(text, " - ")
				if len(nummms) != 2 {
					panic("Ошибка ввода")
				}
				n1, err := strconv.Atoi(nummms[0])
				if err == nil{
					panic("Первое не может быть число")
				}
				_ = n1
				t1 := strings.TrimSpace(nummms[0])
				nummms[0] = t1				
				n2, err := strconv.Atoi(nummms[1])
				if err == nil{
					panic("Нельзя из строки вычетать числа")
				}
				_ = n2
				t2 := strings.TrimSuffix(nummms[1], "\r\n")      //убрать "\r\n"
				t2 = strings.TrimSpace(t2)          		     //убрать пробелы
				if utf8.RuneCountInString(t1) > 12{
					panic("Слишком длинная строка!!!")
				}
				if utf8.RuneCountInString(t2) > 12{
						panic("Слишком длинная строка!!!")
				}
				num1, err := strconv.Unquote(nummms[0])
				if err != nil{
					panic("Ошибка ввода")
				}	
				num2, err := strconv.Unquote(nummms[1])	
				if err != nil{
					panic("Ошибка ввода")
				}	
				res := strings.TrimSuffix(num1, num2)
				fmt.Println(strconv.Quote(res))
				
			}else if strings.Contains(text, " * "){
				nummms := strings.Split(text, " * ")
				if len(nummms) != 2 {
					panic("Ошибка ввода")
				}
				n1, err := strconv.Atoi(nummms[0])
				if err == nil{
					panic("Первое не может быть число")
				}
				_ = n1
				num2, err := strconv.Atoi(nummms[1])
				if err != nil{
					panic("умножать на строку нельзя")
				}
				if utf8.RuneCountInString(nummms[0]) > 12{
					panic("Слишком длинная строка!!!")
				}
				if num2 > 11 || num2 < 1{
					panic("Число должно быть от 1 до 10")
				}
				num1, err := strconv.Unquote(nummms[0])
				if err != nil{
					panic("Ошибка ввода")
				}				
				fmt.Println(strconv.Quote(strings.Repeat(num1, num2)))

			}else if strings.Contains(text, " / "){
				nummms := strings.Split(text, " / ")
				if len(nummms) != 2 {
					panic("Ошибка ввода")
				}
				num2, err := strconv.Atoi(nummms[1])
				if err != nil{
					panic("делить на строку нельзя")
				}

				num1, err := strconv.Unquote(nummms[0])
				if err != nil{
					panic("Первое не может быть число")
				}				
				numLen := len(num1)/num2
				fmt.Println(strconv.Quote(num1[:numLen]))
			}
		}else {
				panic("не верный ввод")
			}
	} 