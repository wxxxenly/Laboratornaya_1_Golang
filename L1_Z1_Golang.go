package main
// импортируем библиотеки "fmt" и "math"
import "fmt"
import "math"

func main() {
    // создаем константу числа Пи
    const PI float64 = 3,1415926535
    // вводим число X
    var x float64
    fmt.Println("Введите число X: ")
	fmt.Scanf("%x\n", &x)
	// вводим число Y
	var y float64
    fmt.Println("Введите число Y: ")
	fmt.Scanf("%y\n", &y)
	// вводим число Z
	var z float64
    fmt.Println("Введите число Z: ")
	fmt.Scanf("%z\n", &z)
    // создаем переменные a,b,c, присваиваем им тип float64 и производим операции над ними
    var a float64 = (2.4 * math.Pow(10,-3) * math.Pow(x,(4/x))*(x-PI/math.Sqrt(3)))/(0.21*math.Pow(10,-5)*math.Sqrt(math.Abs(y)/math.Abs(z))+math.Pow(x,2)*math.Abs(y-math.Pow(math.Tan(math.Pow(x,2)),3)))
    var b float64 = ((y-x)/(z-x) + (y-(y/z)-x)/(math.Pow(math.Sin(x/y),2)))
    var c float64 = (math.Pow(math.Sqrt(math.Log(math.Abs(z-x))*math.Cos(a)*x-math.Sin((a-b)/(a*b)),3)))
    // выводим результат в строчку
    fmt.Println("Результаты вычислений: A = ", a, "; B = ", b,"; C = ", c)
}
