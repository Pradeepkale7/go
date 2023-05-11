// package main

// import "fmt"

// func main() {
// 	var a, b, c int

// 	fmt.Print("1:add")
// 	fmt.Println("2: sub")
// 	fmt.Println("3: mul")
// 	fmt.Println("4: div")
// 	fmt.Println("5: rem")
// 	fmt.Println("6: exit")
// 	fmt.Println("whats your choice ")
// 	fmt.Scanln(&c)

// 	switch c {
// 	case 1:
// 		fmt.Println("Enter the 2 nos")
// 		fmt.Scanln(&a, &b)
// 		z := a + b
// 		fmt.Println("\n sum =", z)

// 	case 2:
// 		fmt.Println("Enter the 2 nos")
// 		fmt.Scanln(&a, &b)
// 		z := a - b
// 		fmt.Println("\n sum =", z)

// 	case 3:
// 		fmt.Println("Enter the 2 nos")
// 		fmt.Scanln(&a, &b)
// 		z := a * b
// 		fmt.Println("\n sum =", z)
// 	case 4:
// 		fmt.Println("Enter the 2 nos")
// 		fmt.Scanln(&a, &b)
// 		z := a / b
// 		fmt.Println("\n sum =", z)
// 	case 5:
// 		fmt.Println("Enter the 2 nos")
// 		fmt.Scanln(&a, &b)
// 		z := a % b
// 		fmt.Println("\n rem =", z)
// 	case 6:
// 		break
// 	}

// }

// ///////////2222222222222222222222222222222222//////////////
// package main

// import "fmt"

// func main() {
// 	var n, t3 int
// 	fmt.Printf("enter terms: ")
// 	fmt.Scanf("%d", &n)
// 	t1 := 0
// 	t2 := 1
// 	fmt.Printf("%d%d", t1, t2)
// 	for i := 1; i <= n-2; i++ {
// 		t3 = t1 + t2
// 		fmt.Printf("%d\t", t3)
// 		t1 = t2
// 		t2 = t3
// 	}
// }
/////////////////////////333333333333333333333333333

// package main

// import "fmt"

// func pal(num int) int {
// 	var i, rem, n int
// 	for i = num; i > 0; {
// 		rem = i % 10
// 		n = n*10 + rem
// 		i = i / 10
// 	}
// 	return n
// }

// func main() {
// 	var z, num int

// 	fmt.Print("Enter the number ")
// 	fmt.Scanf("%d", &num)
// 	z = pal(num)

// 	if z == num {
// 		fmt.Print("The number is paldrome ")
// 	} else {
// 		fmt.Print("The numvetr is not paldrom e")
// 	}
// }

/////////////////////////////444444444444
//sum od digit using recursion
// package main

// import "fmt"

// func rsod(num int) int {

// 	if num == 0 {

// 		return 0
// 	}

// 	return (num%10 + rsod(num/10))
// }

// func main() {
// 	var num int
// 	var r int
// 	fmt.Printf("enter no: ")
// 	fmt.Scanf("%d", &num)
// 	r = rsod(num)
// 	fmt.Printf("sod is: %d\n", r)
// }

///////////////////////////////55555555555555555555555555

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	f1, e1 := os.Create("Sample.txt")
// 	if e1 != nil {
// 		fmt.Println("Unable to open file: %s", e1)
// 	}
// 	len, e1 := f1.WriteString("Welcome")
// 	if e1 != nil {
// 		fmt.Println("Unable to write data: %s", e1)
// 	}
// 	f1.Close()
// 	fmt.Printf("%d character written successfully into file", len)
// }

//////////////////////////////////6666666666666666666

////////////////////////7777777777777777777

// package main

// import "fmt"

// func main() {
// 	var i, j, rows, columns int

// 	var orgMat [10][10]int
// 	var transposeMat [10][10]int

// 	fmt.Print("Enter the Matrix rows and Columns = ")
// 	fmt.Scan(&rows, &columns)

// 	fmt.Println("Enter Matrix Items to Transpose = ")
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			fmt.Scan(&orgMat[i][j])
// 		}
// 	}
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			transposeMat[j][i] = orgMat[i][j]
// 		}
// 	}
// 	fmt.Println("*** The Transpose Matrix Items are ***")
// 	for i = 0; i < columns; i++ {
// 		for j = 0; j < rows; j++ {
// 			fmt.Print(transposeMat[i][j], "  ")
// 		}
// 		fmt.Println()
// 	}
// }

//////////////////888888888888888888888

// package main

// import "fmt"

// type Book struct {
// 	BookID int
// 	Title  string
// 	Author string
// 	Price  float64
// }

// func main() {
// 	var n int
// 	fmt.Print("Enter the number of books: ")
// 	fmt.Scan(&n)
// 	// Create a slice of n books
// 	books := make([]Book, n)
// 	// Accept book details from user
// 	for i := 0; i < n; i++ {
// 		fmt.Printf("Enter the details of book %d:\n", i+1)
// 		fmt.Print("Book ID: ")
// 		fmt.Scan(&books[i].BookID)
// 		fmt.Print("Title: ")
// 		fmt.Scan(&books[i].Title)
// 		fmt.Print("Author: ")
// 		fmt.Scan(&books[i].Author)
// 		fmt.Print("Price: ")
// 		fmt.Scan(&books[i].Price)
// 		fmt.Println()
// 	}
// 	// Display book details to user
// 	fmt.Println("Book details:")
// 	for i, book := range books {
// 		fmt.Printf("Book %d:\n", i+1)
// 		fmt.Printf("Book ID: %d\n", book.BookID)
// 		fmt.Printf("Title: %s\n", book.Title)
// 		fmt.Printf("Author: %s\n", book.Author)
// 		fmt.Printf("Price: %.2f\n", book.Price)
// 		fmt.Println()
// 	}
// }

// ///////////////////////////////////////////////////////
// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Shape interface {
// 	area() float64
// 	perimeter() float64
// }
// type Circle struct {
// 	radius float64
// }

// func (c Circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c Circle) perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// type Rectangle struct {
// 	length float64
// 	width  float64
// }

// func (r Rectangle) area() float64 {
// 	return r.length * r.width
// }
// func (r Rectangle) perimeter() float64 {
// 	return 2 * (r.length + r.width)
// }
// func main() {
// 	c := Circle{radius: 5}
// 	r := Rectangle{length: 4, width: 6}
// 	fmt.Printf("Area of circle: %.2f\n", c.area())
// 	fmt.Printf("Perimeter of circle: %.2f\n", c.perimeter())
// 	fmt.Printf("Area of rectangle: %.2f\n", r.area())
// 	fmt.Printf("Perimeter of rectangle: %.2f\n", r.perimeter())
// }

/////////////////////////999999999999999999999999

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func isPalindrome(num int) bool {
// 	// Convert the number to a string
// 	str := strconv.Itoa(num)
// 	// Check if the string is equal to its reverse
// 	for i := 0; i < len(str)/2; i++ {
// 		if str[i] != str[len(str)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }
// func main() {
// 	var num int
// 	fmt.Print("Enter a number: ")
// 	fmt.Scanln(&num)
// 	if isPalindrome(num) {
// 		fmt.Println(num, "is a palindrome")
// 	} else {
// 		fmt.Println(num, "is not a palindrome")
// 	}
// }

////////////1010101010101010101010

// package main

// import "fmt"

// type geometry interface {
// 	area() float64
// 	perimeter() float64
// }
// type rectangle struct {
// 	width, height float64
// }

// func (r rectangle) area() float64 {
// 	return r.width * r.height
// }
// func (r rectangle) perimeter() float64 {
// 	return 2*r.width + 2*r.height
// }
// func main() {
// 	var g geometry = rectangle{width: 3, height: 4}
// 	fmt.Println(g)
// 	if r, ok := g.(rectangle); ok {
// 		fmt.Println("Width:", r.width)
// 		fmt.Println("Height:", r.height)
// 		fmt.Println("Area:", r.area())
// 		fmt.Println("Perimeter:", r.perimeter())
// 	} else {
// 		fmt.Println("Not a rectangle")
// 	}
// }

/////////////////////11 11 11 11 11 11 11 11

// package main

// import "fmt"

// func main() {
// 	var num int
// 	fmt.Print("Enter a number: ")
// 	fmt.Scanln(&num)
// 	if num >= 10 && num <= 99 {
// 		fmt.Println("The number is two digits")
// 	} else {
// 		fmt.Println("The number is not two digits")
// 	}
// }

// //////12 12 12 12 12 12 12 12
// package main

// import "fmt"

// func swap(a *int, b *int) {
// 	temp := *a
// 	*a = *b
// 	*b = temp
// }
// func main() {
// 	var num1, num2 int
// 	fmt.Print("Enter the first number: ")
// 	fmt.Scanln(&num1)
// 	fmt.Print("Enter the second number: ")
// 	fmt.Scanln(&num2)
// 	fmt.Printf("Before swapping: num1 = %d, num2 = %d\n", num1, num2)
// 	swap(&num1, &num2)
// 	fmt.Printf("After swapping: num1 = %d, num2 = %d\n", num1, num2)
// }

// ////13 13 13 13 13 13 13 13
// package main

// import "fmt"

// func main() {
// 	evenSum := 0
// 	oddSum := 0
// 	for i := 1; i <= 100; i++ {
// 		if i%2 == 0 {
// 			evenSum += i
// 		} else {
// 			oddSum += i
// 		}
// 	}
// 	fmt.Printf("Sum of all even numbers between 1 to 100 is %d\n", evenSum)
// 	fmt.Printf("Sum of all odd numbers between 1 to 100 is %d\n", oddSum)
// }

///////////// 14 14 14 14 14 14 14

// package main

// import "fmt"

// func main() {
// 	// Creating a slice
// 	nums := []int{1, 2, 3, 4, 5}
// 	fmt.Println("Original slice:", nums)
// 	// Appending elements to the slice
// 	nums = append(nums, 6, 7, 8)
// 	fmt.Println("Slice after appending:", nums)
// 	// Removing an element from the slice
// 	nums = append(nums[:3], nums[4:]...)
// 	fmt.Println("Slice after removing 4th element:", nums)
// 	// Copying a slice
// 	newNums := make([]int, len(nums))
// 	copy(newNums, nums)
// 	fmt.Println("New slice after copying:", newNums)
// }

///////////15 15 15 15 15 15 15 15 15

// package main

// import "fmt"

// // Function to calculate sum and product of two numbers
// func calcSumAndProduct(a, b int) (int, int) {
// 	sum := a + b
// 	product := a * b
// 	return sum, product
// }
// func main() {
// 	// Calling the calcSumAndProduct() function and storing the results in two

// 	sum, product := calcSumAndProduct(5, 10)
// 	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
// }

// /////////////16 16 16 16 16 16 16
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		fmt.Println(i)
// 		time.Sleep(time.Duration(250) * time.Millisecond)
// 	}
// }

// ////////////17 17 17 17 17 17 17 17 17 17
// package main

// import "fmt"

// func AddSubtractMultiplyDivide(a, b float64) (float64, float64, float64, float64) {
// 	return a + b, a - b, a * b, a / b
// }
// func main() {
// 	x := 10.0
// 	y := 5.0
// 	add, subtract, multiply, divide := AddSubtractMultiplyDivide(x, y)
// 	fmt.Printf("Addition of %.2f and %.2f is %.2f\n", x, y, add)
// 	fmt.Printf("Subtraction of %.2f from %.2f is %.2f\n", y, x, subtract)
// 	fmt.Printf("Multiplication of %.2f and %.2f is %.2f\n", x, y, multiply)
// 	fmt.Printf("Division of %.2f by %.2f is %.2f\n", x, y, divide)
// }

//////////18 18 18 18 18 18 18 18 18 18

// package main

// import "fmt"

// func printMultiplicationTable(num int) {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Printf("%d x %d = %d\n", num, i, num*i)
// 	}
// }
// func main() {
// 	var num int
// 	fmt.Println("Enter a number to print its multiplication table:")
// 	fmt.Scanln(&num)
// 	printMultiplicationTable(num)
// }

///////////////19 19 19 19 19 19 19 19 19 19 19

// package main

// import "fmt"

// func cal(m int, n int) (int, int) {
// 	c := m + n
// 	d := m - n
// 	return c, d
// }
// func main() {
// 	var a, b, z1, z2 int
// 	fmt.Printf("enter 2 number")
// 	fmt.Scanf("%d%d", &a, &b)
// 	z1, z2 = cal(a, b)
// 	fmt.Printf("\n add=%d", z1)
// 	fmt.Printf("\n sub=%d", z2)
// }

// /////////20 20 20 20 20 20 20 20 20 20
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	message := "Today is Thursday"
// 	filename := "sample.txt"
// 	f1, e1 := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
// 	if e1 != nil {
// 		fmt.Println(e1)
// 		os.Exit(-1)
// 	}
// 	defer f1.Close()
// 	fmt.Fprintf(f1, "%s\n", message)
// }
