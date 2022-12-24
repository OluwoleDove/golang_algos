/*SOME ALGORITHMS I WROTE IN GOLANG */

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// THIS PROGRAM TAKES TWO INTEGER NUMBERS AND RETURNS AN ARRAY OF ODD NUMBERS
// AND A COUNT OF EVEN NUMBERS BETWEEN THE INTEGERS
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
func even_odd(i int, j int) (return_array [][]int) {
	if i > j {
		k := j
		j = i
		i = k
	}

	fmt.Printf("ITERATING BETWEEN %d and %d\n", i, j)
	var odd_num_array []int  //An empty list
	var even_num_array []int //integer value for even number counts initialized to zero

	for number := i; number <= j; number++ {
		if (number % 2) > 0 {
			odd_num_array = append(odd_num_array, number)
		} else if (number % 2) == 0 { //is number an even number, i.e is the remainder = 0?
			even_num_array = append(even_num_array, number)
		}
	}
	return_array = append(return_array, odd_num_array)
	return_array = append(return_array, even_num_array)
	return return_array
}

/*PROGRAM TO PRINT TRIANGLE OF CHARACTERS*/
func char_tri(char string, lim int) (main_str string) {
	if (char == "") || (char == " ") {
		fmt.Println("Your first input can't be empty")
	}

	lim = int(lim)
	fmt.Println("\nPRINTING A TRIANGLE OF CHARACTERS BY CALLING PRINT METHOD ONCE")
	var str_out string

	for i := 1; i <= lim; i++ {
		str_out += char
		main_str += str_out + "\n"
	}
	return
}

func check_anagram(A [5]string, B [5]string) (C []bool) {
	//Declare a count 'variable'&&nitialize it to zero
	var count int
	fmt.Println()

	for i := 0; i < len(A); i++ { //Iteratively select each element in Array A
		for j := 0; j < len(A[i]); j++ { //Iteratively select the letters in each element
			for k := 0; k < len(B[i]); k++ { //Check for the letter in an element in B that has the same index as that of A
				if A[i][j] == B[i][k] {
					count = count + 1
					break
				}
			}
		}
		C = append(C, count == len(B[i]))
		count = 0
	}
	return C
}

// SORTING func
func _bubble_sort(my_array [19]int) (result []string) {
	var write_array []string
	delim := ","
	for i := 0; i < len(my_array)-1; i++ {
		for j := 0; j < len(my_array)-1; j++ {
			if my_array[j+1] < my_array[j] {
				num := my_array[j]
				my_array[j] = my_array[j+1]
				my_array[j+1] = num
			}
		}
	}

	write_array = append(write_array, strings.Trim(strings.Replace(fmt.Sprint(my_array), " ", delim, -1), "[]"))
	if err := writeLines(write_array, "sorted.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
	result, err := readLines("sorted.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return result
}

// MULTIPLICATION TABLE
func multi_table(n int) {
	fmt.Println("PRINTING MULTIPLICATION TABLE HORIZONTALLY")
	var num int
	var out_str string //Initializing output string
	//Below are 12 instances of i&&144 instances of j

	fmt.Println("\nPRINTING TABLES VERTICALLY")
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			num = i * j
			out_str += strconv.Itoa(i) + " * " + strconv.Itoa(j) + " = " + strconv.Itoa(num) + "\n"
		}
		out_str += "\n"
	}
	fmt.Println(out_str)
	out_str = "" //After eact complete j iterations, initialize output string again
	fmt.Println()
	fmt.Println("You can also have it side by side")
	fmt.Println("PRINTING 9 BY 12 MULTIPLICATION TABLE SIDE BY SIDE")

	for i := 1; i <= 13; i++ {
		for j := 1; j <= 13; j++ {
			num = j * i
			out_str += strconv.Itoa(j) + "*" + strconv.Itoa(i) + "=" + strconv.Itoa(num) + "\t\t"
		}
		fmt.Println(out_str)
		out_str = "" //After each complete j iterations, initialize output string again
	}
}

func determine_palindrome(palindrome string) {
	iter_count := math.Floor(float64(len(palindrome)) / 2) //Floor division
	var truth_array []bool
	neg := 1
	for x := 0; x < int(iter_count); x++ {
		//Comapre letters&&append the boolean outcomes to the list
		if palindrome[x] == palindrome[len(palindrome)-neg] {
			truth_array = append(truth_array, true)
		} else {
			truth_array = append(truth_array, false)
		}
		neg += 1
	}
	fmt.Println(truth_array)
	if len(truth_array) < len(palindrome) {
		fmt.Printf("%s is not a palindrome", palindrome)
	} else {
		fmt.Printf("%s is a palindrome", palindrome)
	}
}

func find_prefix(_array [5]string) (pref string) {
	//sorting the List
	for i := 0; i < len(_array)-1; i++ {
		for j := 0; j < len(_array)-1; j++ {
			if len(_array[j+1]) < len(_array[j]) {
				_word := _array[j]
				_array[j] = _array[j+1]
				_array[j+1] = _word
			}
		}
	}
	fmt.Println(_array)

	//CHCKING FOR PREFIXES
	var count int
	var letter string

	for i := 0; i < len(_array[0]); i++ {
		for j := 0; j < len(_array)-1; j++ {
			letter = string(_array[j][i])
			if _array[j][i] == _array[j+1][i] {
				count += 1
			} else {
				count -= 1
			}
		}
		if count == len(_array)-1 {
			pref += letter
		}
		count = 0
	}
	return
}

// FACTORIAL func
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

// FIBONACCI SERIES - THE RECURSIVE WAY
func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

// COMBINATION func CALLING FACTORIAL func
func combination(n int, r int) int {
	c_val := factorial(n) / (factorial(n-r) * factorial(r))
	return c_val
}

// PASCAL_T FUNCTIOIN CALLS COMBINATION func
func pascal_T(n int) string {
	var my_str string
	var val int
	for j := 0; j <= n; j++ {
		val = combination(n, j)
		my_str += strconv.Itoa(val) + "     "
	}
	return my_str
}

func center(s string, w int) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
}

// THIS func CALLS OTHER funcS
func test_algo() {
	var j int
	var k int
	//EXECUTING EVEN_ODD func
	fmt.Println("THIS PROGRAM OUTPUTS THE ODD NUMBERS AND THE COUNT OF EVEN BETWEEN TWO INTEGERS")
	fmt.Println("Your first integer number: ")
	fmt.Scanf("%v", &j)

	fmt.Println("Your second integer number: ")
	fmt.Scanf("%v", &k)

	even_odd_array := even_odd(j, k)
	fmt.Printf("The odd numbers are: %v", even_odd_array[0])
	fmt.Printf("The even numbers are: %v", even_odd_array[1])
	fmt.Println()

	//CALLING func TO PRINT TRIANGLE OF CHARACTERS
	var i string
	fmt.Println("LET'S HAVE A TRIANGLE OF CHARACTERS, SHALL WE?")
	fmt.Println("Type a character please: ")
	fmt.Scanf("%v", &i)
	fmt.Println("Type an integer: ")
	fmt.Scanf("%v", &j)
	triangle := char_tri(i, j)
	fmt.Println(triangle + "\n")

	//LET'S PRINT MULTIPLICATION TABLES
	fmt.Println("\nLET'S GENERATE MULTIPLICATION TABLE")
	fmt.Println("Your first integer number: ")
	fmt.Scanf("%v", &j)
	multi_table(j)

	//PROGRAM ANAGRAM
	//Anagrams are words that are made up of the same letters but have different meanings
	fmt.Println("THIS PROGRAM COMPARES TWO LISTS TO CHECK FOR ANAGRAMS \nExamples:")
	A := [5]string{"abode", "man", "live", "heart", "ear"}
	B := [5]string{"adobe", "nan", "evil", "earth", "car"}
	fmt.Println("A = ['abode', 'man', 'live', 'heart', 'ear']\nB = ['adobe', 'nan', 'evil', 'earth', 'car']")
	//declare an empty list
	anagram_check := check_anagram(A, B)
	fmt.Println(anagram_check)
	fmt.Println()

	//Printing THE LONGEST PREFIX
	word_array := [5]string{"flower", "flow", "flight", "floor", "flood"}
	fmt.Println("\nChecking the longest prefix in the list")
	_prefix := find_prefix(word_array)
	fmt.Printf("\nLongest Prefix is %s", _prefix)

	//Printing FACTORIAL
	fmt.Println("\nType a number to print it's factorial: ")
	fmt.Scanf("%v", &j)
	fmt.Println("PRINTING FACTORIAL")
	fmt.Println(factorial(j))
	fmt.Println()

	//ITERATIVELY GENERATING FIBONACCI SERIES
	fmt.Println("LET'S ITERATIVELY PRINT FIRST 10 NUMBERS OF THE FIBONACCI SERIES")
	var fib_array []int
	new_item := 1
	if len(fib_array) < 2 {
		fib_array = append(fib_array, new_item)
		for len(fib_array) < 10 {
			new_item = fib_array[len(fib_array)-1] + fib_array[len(fib_array)-2]
			fib_array = append(fib_array, new_item)
		}
	}
	fmt.Println(fib_array)
	fmt.Println()

	//fib_series CODE
	fmt.Println("\nNOW LET'S DO IT THE RECURSION WAY")
	var fib_series []int
	fmt.Println("Type a number: ")
	fmt.Scanf("%v", &j)
	for i := 0; i < j; i++ {
		fib_series = append(fib_series, fibonacci(j))
	}
	//LET'S PRINT STUFF
	fmt.Println("PRINTING FIBONACCI")
	fmt.Println(fib_series)
	fmt.Println("Type a number, let's get the Fibonacci value: ")
	fmt.Scanf("%v", &j)
	fmt.Println(fibonacci(j))
	fmt.Println()

	//PROGRAM PALINDROME - A Palindrome is a word that spells the same in the same direcion
	//Examples are level, racecar, saippuakivikauppias
	fmt.Println("THIS PROGRAM CHECKS IF A WORD IS A PALINDROME")
	fmt.Println("A Palindrome is a word that spells the same in both directions.")
	fmt.Println("Type a word: ")
	var word_strng string
	fmt.Scanf("%v", &word_strng)
	determine_palindrome(word_strng)

	//PROGRAM SORT LIST
	fmt.Println("THIS func SORTS AN ARRAY IN ASCENDING ORDER")
	fmt.Println()

	my_array := [19]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9}
	fmt.Printf("Sorting this list %v", my_array)
	my_sorted_array := _bubble_sort(my_array)
	fmt.Println(my_sorted_array)

	//This program prints Pascal Triangle
	fmt.Println("PRINTING PASCAL'S TRIANGLE")
	fmt.Println("Type an integer number ")
	fmt.Scanf("%v", &j)
	for i := 0; i < j; i++ {
		fmt.Println(center(pascal_T(i), 80))
	}
	//Pascal Triangle - Every Array starts&&ends with 1
	fmt.Println("NOW THAT WAS A funcAL APPROACH TO IT.\nLET'S GENERATE PASCTAL'S TRIANGLE ITERATIVELY")
	unit_num := 1
	var col_array []int
	var new_array []int

	type pascal_obj struct {
		_key   int
		_value []int
	}

	var pascal_group = map[string]*pascal_obj{"1": {1, []int{1}}}
	var pascal_tri string
	fmt.Println("Type a number ")
	fmt.Scanf("%v", &j)

	for i := 1; i <= int(j)+1; i++ {
		elem := new(pascal_obj)
		col_array = append(col_array, unit_num)
		_index := strconv.Itoa(i)
		//Just using i to guage row lengths
		if len(col_array) == i {
			elem._key = i
			elem._value = col_array
			pascal_group[_index] = elem
			col_array = nil
		} else if len(col_array)/2 == 1 {
			col_array = append(col_array, unit_num)
			elem._key = i
			elem._value = col_array
			pascal_group[_index] = elem
		} else {
			col_array = nil
			col_array = append(col_array, 1)
			//new_array = pascal_obj[i-1]
			for j := 0; j < len(new_array)-1; j++ {
				col_array = append(col_array, new_array[j]+new_array[j+1])
			}
			col_array = append(col_array, 1)
			elem._key = i
			elem._value = col_array
			col_array = nil
		}
	}

	fmt.Println(pascal_group)
	fmt.Println()
	fmt.Println("PASCAL'S TRANGLE")
	for k := 1; k <= j; k++ {
		var col_str string
		j := strconv.Itoa(k)
		col_str = strings.Trim(strings.Replace(fmt.Sprint(pascal_group[j]._value), " ", ",", -1), "[]")
		pascal_tri += col_str + "\n"
	}
	fmt.Println(center(pascal_tri, 80))
}

func main() {
	var check string
	fmt.Println("Let's Start, shall we? Type 'Y' to start and 'N' to quit.")
	fmt.Scanf("%v", &check)

	for check == "Y" || check == "y" {
		test_algo()
		fmt.Println("Ha-ha! Super you. Wanna try again? Y/N")
		fmt.Scanf("%v", &check)
	}
}
