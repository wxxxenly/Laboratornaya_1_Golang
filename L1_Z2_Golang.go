package main
import "fmt"

func main() {
  var a, first, second int
    
  fmt.Println("Введите число: ")
	fmt.Scanf("%d\n", &a)
	
	first = a/10
	second = a % 10
	
	fmt.Println("Сумма двух чисел равна: ", first+second, ",а произведение их равно: ", first*second)
}
