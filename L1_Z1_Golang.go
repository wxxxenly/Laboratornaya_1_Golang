package main
import "fmt"
import "math"

func main() {
    const PI = 3.1415926535
    
    var x float64
    fmt.Println("Введите число X: ")
	fmt.Scanf("%x\n", &x)
	
	var y float64
    fmt.Println("Введите число Y: ")
	fmt.Scanf("%y\n", &y)
	
	var z float64
    fmt.Println("Введите число Z: ")
	fmt.Scanf("%z\n", &z)
    
    var a float64 = (2.4 * math.Pow(10, -3) * math.Pow(x, (4/x))*(x-PI/math.Sqrt(3)))/(0.21*math.Pow(10, -5)*math.Sqrt((math.Abs(y)/math.Abs(z)))+math.Pow(x, 2)*math.Abs(y-math.Pow(math.Tan((math.Pow(x, 2))), 3)))
    var b float64 = ((y-x)/(z-x) + (y-(y/z)-x)/(math.Pow((math.Sin(x/y)), 2)))
    var c float64 = (math.Pow((math.Sqrt(math.Log(math.Abs(z-x))*math.Cos(a)*x-math.Sin((a-b)/(a*b)))), 3))
    
    fmt.Println("Результаты вычислений: A = ", a, "; B = ", b,"; C = ", c)
}
