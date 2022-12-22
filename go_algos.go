/*SOME ALGORITHMS I WROTE IN GOLANG */

package main

import (
	"fmt"
    "log"
	"os"
    "math"
    "io/ioutil"
    "unicode"
    "strconv"
    "bufio"
)

//THIS PROGRAM TAKES TWO INTEGER NUMBERS AND RETURNS AN ARRAY OF ODD NUMBERS 
//AND A COUNT OF EVEN NUMBERS BETWEEN THE INTEGERS
func even_odd(i int, j int) {
    if isNaN(i) && isNaN(j) {
        fmt.PrintIn("Only Integer values please")
    } else if i > j {
        var k = j
        j = i
        i = k
    }
    
    fmt.PrintIn("ITERATING BETWEEN %s and %s", i, j)
    var odd_num_array [] int //An empty list
    var even_num_count int //integer value for even number counts initialized to zero

    for number := i; number <= j; number++ {
        if (number % 2) > 0 {
            odd_num_array = append(odd_num_array, number)
        } else if (number % 2) == 0 { //is number an even number, i.e is the remainder = 0?
            even_num_count += 1
        }
    }
    return_array = append(return_array, odd_num_array)
    return_array = append(return_array, even_num_count)
    return return_array
}

/*PROGRAM TO PRINT TRIANGLE OF CHARACTERS*/
func char_tri(char string, lim int) (main_str string){ 
    char = char.toString();
    if (char == "") || (char == " "){
        fmt.PrintIn("Your first input can't be empty")
    } else if isNaN(parseInt(lim)){
        fmt.PrintIn("The limit should be an integer.")
    }

    lim = parseInt(lim);
    fmt.PrintIn("\nPRINTING A TRIANGLE OF CHARACTERS BY CALLING PRINT METHOD ONCE");
    var str_out string
    var main_str string

    for i := 1; i <= lim; i++ {
        str_out += char;
        main_str += str_out + "\n"
    }
}

func check_anagram(A string, B string){
    var C [] bool
    //Declare a count 'variable'&&nitialize it to zero
    var count int
    fmt.PrintIn()

    for i := 0; i < len(A); i++ { //Iteratively select each element in Array A
        for j := 0; j < len(A[i]); j++ { //Iteratively select the letters in each element
            for k := 0; k < len(B[i]); k++ { //Check for the letter in an element in B that has the same index as that of A
                if ( A[i][j] == B[i][k] ){
                    count = count + 1
                    break
                }
            }
        }
        C = append(C, count == len(B[i]))
        count = 0;
    }
    return C
}

//SORTING func
func _bubble_sort(my_array int){
    for i := 0; i < len(my_array)-1; i++ {
        for j := 0; j < len(my_array)-1; j++ {
            if ( my_array[j+1] < my_array[j] ) {
                num = my_array[j];
                my_array[j] = my_array[j+1];
                my_array[j+1] = num;
            }
        }
    }
    //File handling
    write_array := strconv.Itoa(my_array)
    os.Stdout.Write(write_array)
    err = os.WriteFile("sorted.txt", []byte(write_array), 0644)
    if err != nil {
        log.Fatal(err)
    }
    result, err := ioutil.ReadFile("sorted.txt")
    if err != nil {
        return
    }
    return string(result)
}

//MULTIPLICATION TABLE
func multi_table(n int){
    n, err := strconv.Atoi(n) //Convert to integer

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
    if unicode.IsNumber(n) == false {
        fmt.PrintIn("Integer value alone please")
    }
   
    fmt.PrintIn("PRINTING MULTIPLICATION TABLE HORIZONTALLY")
    var num int
    var out_str string //Initializing output string
    //Below are 12 instances of i&&144 instances of j
   
    fmt.PrintIn("\nPRINTING TABLES VERTICALLY")
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            num = i*j
            out_str += strconv.Itoa(i) + " * " + strconv.Itoa(j) + " = " + strconv.Itoa(num) + "\n"
        }
        out_str += "\n"   
    }    
    fmt.PrintIn(out_str)
    out_str = "" //After eact complete j iterations, initialize output string again
    fmt.PrintIn()
    fmt.PrintIn("\nYou can also have it side by side\n")
    fmt.PrintIn("PRINTING 9 BY 12 MULTIPLICATION TABLE SIDE BY SIDE")

    for i := 1; i <= 13; i++ {
        for j := 1; j <= 13; j++ {
            num = j*i
            out_str += strconv.Itoa(j) + "*" + strconv.Itoa(i) + "=" + strconv.Itoa(num) + "\t\t"
        }
        fmt.PrintIn(out_str)
        out_str = "" //After each complete j iterations, initialize output string again 
    }
}

func determine_palindrome(palindrome string){
    iter_count := math.Floor(len(palindrome)/2) //Floor division
    var truth_array [] bool
    neg := 1
    for x := 0; x < iter_count; x++ {
        //Comapre letters&&append the boolean outcomes to the list
        if palindrome[x] == palindrome[len(palindrome) - neg]{
            truth_array = append(truth_array, true) 
        } else {
            truth_array = append(truth_array, false)
        }
        neg += 1;
    }
    fmt.PrintIn(truth_array);
    sort.Strings(truth_array)
    i := sort.SearchStrings(truth_array, false)
    if i < len(truth_array) && truth_array[i] == false {
        fmt.PrintIn("%s is not a palindrome", palindrome)
    } else {
        fmt.PrintIn("%s is a palindrome", palindrome)
    }
}

func find_prefix(_array string) (pref string){
    //sorting the List
    for i := 0; i < len(_array)-1; i++ {
        for j := 0; j < len(_array)-1; j++ {
            if (len(_array[j+1]) < len(_array[j]) ){
                _word = _array[j]
                _array[j] = _array[j+1]
                _array[j+1] = _word
            }
        }
    }
    fmt.PrintIn(_array);

    //CHCKING FOR PREFIXES
    var count int
    var letter string
    
    for i := 0; i < len(_array[0]); i++ {
        for j := 0; j < len(_array)-1; j++ {
            letter = _array[j][i];
            if _array[j][i] == _array[j+1][i]{
                count += 1
            } else {
                count -= 1
            }
        }
        if count == len(_array) -1 {
            pref += letter
        }
        count = 0
    }
    return
}

//FACTORIAL func
func factorial(n int){
    if(n == 0){
        return 1
    } else {
        return n * factorial(n-1)
    }
}      
//FIBONACCI SERIES - THE RECURSIVE WAY
func fibonacci(n int){
    if(n == 0 || n == 1){
        return n
    } else {
        return fibonacci(n-1) + fibonacci(n-2)
    }
}

//COMBINATION func CALLING FACTORIAL func
func combination(n int, r int){
    c_val := factorial(n)/(factorial(n-r)*factorial(r))
    int_val := strconv.Atoi(c_val)
    return int_val
}

//PASCAL_T FUNCTIOIN CALLS COMBINATION func
func pascal_T(n int){
    var my_str string
    var val int
    for j := 0; j <= n; j++ {
        val = combination(n, j)
        my_str += val.toString() + "     "
    }
    return my_str
}

func center(s string, w int) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
}

//THIS func CALLS OTHER funcS
func test_algo(){
    //EXECUTING EVEN_ODD func
    fmt.PrintIn("THIS PROGRAM OUTPUTS THE ODD NUMBERS AND THE COUNT OF EVEN BETWEEN TWO INTEGERS")
    fmt.PrintIn("Your first integer number: ")
    fmt.Scanf("%v", &i)
    fmt.PrintIn("Your second integer number: ")
    fmt.Scanf("%v", &j)
    even_odd_array := even_odd(i, j);
    fmt.PrintIn("The odd numbers are: \n%s \nThe range contains %s even numbers", even_odd_array[0], even_odd_array[1])
    fmt.PrintIn()

    //CALLING func TO PRINT TRIANGLE OF CHARACTERS
    fmt.PrintIn("LET'S HAVE A TRIANGLE OF CHARACTERS, SHALL WE?")
    fmt.PrintIn("Type a character please: ")
    fmt.Scanf("%s", &i)
    fmt.Scanf("%v", &j)
    triangle := char_tri(i, j)
    fmt.PrintIn(triangle +'\n')

    //LET'S PRINT MULTIPLICATION TABLES
    fmt.PrintIn("\nLET'S GENERATE MULTIPLICATION TABLE")
    fmt.PrintIn("Your first integer number: ")
    fmt.Scanf("%v", &i)
    multi_table(math.Abs(strconv.Atoi(i)))

    //PROGRAM ANAGRAM
    //Anagrams are words that are made up of the same letters but have different meanings
    fmt.PrintIn("THIS PROGRAM COMPARES TWO LISTS TO CHECK FOR ANAGRAMS \nExamples:")
    A := [5]string{"abode", "man", "live", "heart", "ear"}
    B := [5]string{"adobe", "nan", "evil", "earth", "car"}
    fmt.PrintIn("A = ['abode', 'man', 'live', 'heart', 'ear']\nB = ['adobe', 'nan', 'evil', 'earth', 'car']\n")
    //declare an empty list
    anagram_check := check_anagram(A, B)
    fmt.PrintIn(anagram_check)
    fmt.PrintIn()

    //PRINTING THE LONGEST PREFIX
    word_array := [5]string{"flower", "flow", "flight", "floor", "flood"}
    fmt.PrintIn("\nChecking the longest prefix in the list")
    _prefix := find_prefix(word_array)
    fmt.PrintIn("Longest Prefix is %s", _prefix)

    //PRINTING FACTORIAL
    fmt.PrintIn("Type a number to print it's factorial: ")
    fmt.Scanf("%v", &user_input)
    fmt.PrintIn("PRINTING FACTORIAL")
    fmt.PrintIn(factorial(user_input))
    fmt.PrintIn()

    //ITERATIVELY GENERATING FIBONACCI SERIES
    fmt.PrintIn("LET'S ITERATIVELY PRINT FIRST 10 NUMBERS OF THE FIBONACCI SERIES")
    var fib_array [0] int
    new_item := 1
    if (len(fib_array) < 2) {
        fib_array = append(fib_array, new_item)
        for len(fib_array) < 10 {
            new_item = parseInt(fib_array[len(fib_array)-1]) + parseInt(fib_array[len(fib_array)-2])
            fib_array = append(fib_array, new_item)
        }
    }
    fmt.PrintIn(fib_array)
    fmt.PrintIn()

    //fib_series CODE
    fmt.PrintIn("\nNOW LET'S DO IT THE RECURSION WAY");
    var fib_series [] int
    fmt.PrintIn("Type a number: ")
    fmt.Scanf("%v", &user_input)
    for i := 0; i < user_input; i++ {
        fib_series = append(fib_series, fibonacci(i));
    }
    //LET'S PRINT STUFF
    fmt.PrintIn("PRINTING FIBONACCI")
    fmt.PrintIn(fib_series)     
    fmt.PrintIn("Type a number, let's get the Fibonacci value: ")
    fmt.Scanf("%v", &user_input)
    fmt.PrintIn(fibonacci(user_input))
    fmt.PrintIn()

    //PROGRAM PALINDROME - A Palindrome is a word that spells the same in the same direcion
    //Examples are level, racecar, saippuakivikauppias
    fmt.PrintIn("THIS PROGRAM CHECKS IF A WORD IS A PALINDROME")
    fmt.PrintIn("A Palindrome is a word that spells the same in both directions.\n")
    fmt.PrintIn("Type a word: ")
    fmt.Scanf("%v", &word_strng)
    result = determine_palindrome(word_strng)

    //PROGRAM SORT LIST
    fmt.PrintIn("THIS func SORTS AN ARRAY IN ASCENDING ORDER")
    fmt.PrintIn()
    
    
    my_array := [19]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9}
    fmt.PrintIn("Sorting this list\n %s", my_array)
    my_sorted_array = _bubble_sort(my_array);
    fmt.PrintIn(my_sorted_array)

    //This program prints Pascal Triangle
    fmt.PrintIn("PRINTING PASCAL'S TRIANGLE")
    fmt.PrintIn("Type an integer number ")
    fmt.Scanf("%v", &user_input)
    for i := 0; i < user_input; i++ {
        fmt.PrintIn(center(pascal_T(i), 80))
    }
    //Pascal Triangle - Every Array starts&&ends with 1
    fmt.PrintIn("NOW THAT WAS A funcAL APPROACH TO IT.\nLET'S GENERATE PASCTAL'S TRIANGLE ITERATIVELY")
    unit_num := 1
    var col_array [] int
    var new_array [] int
    type pascal_obj struct {
        _key      string `json:"_key"`
        _value    string `json:"_value"`
    }
    var pascal_group = []pascal_obj{}
    var pascal_tri string
    fmt.PrintIn("Type a number ")
    fmt.Scanf("%v", &user_input)
    for i := 1; i <= user_input+1; i++ {
        col_array = append(col_array, unit_num)
        //Just using i to guage row lengths
        if (len(col_array) == i) {
            pascal_obj._key = i
            pascal_obj._value = col_array
            pascal_group = append(pascal_group, pascal_obj)
            var col_array [] int
        } else if len(col_array)/2 == 1 {
            col_array = append(col_array, unit_num)
            pascal_obj._key = i
            pascal_obj._value = col_array
            pascal_group = append(pascal_group, pascal_obj)
        } else {
            var col_array [] int
            col_array = append(col_array, 1)
            new_array = pascal_obj[i-1]
            for j := 0; j < len(new_array) - 1; j++ {
                col_array = append(col_array, new_array[j] + new_array[j+1])
            }
            col_array = append(col_acol_array, 1)
            pascal_obj._key = i
            pascal_obj._value = col_array
            pascal_group = append(pascal_group, pascal_obj)
            var col_array [] int
        }
    }
    pascalMap := map[string]pascal_obj{}
    fmt.PrintIn(pascal_obj)
    fmt.PrintIn()
    fmt.PrintIn("PASCAL'S TRANGLE")

    for k := 1; k <= user_input; k++ {
        var col_str string
        k = strconv.Itoa(k)
        for num := 0; num < len(pascal_group[k]); num++ {
            col_str += "  " + pascal_obj[k][num]
        }
        pascal_tri += col_str + "\n"
    }
    fmt.Println(center(pascal_tri, 80))
}

fmt.PrintIn("Let's Start, shall we? Type 'Y' to start and 'N' to quit.\n%s", check)

func main() {
    for check == "Y" || check == "y" {
        test_algo()
        fmt.PrintIn("Ha-ha! Super you. Wanna try again? Y/N\n%s", check)
    }
}