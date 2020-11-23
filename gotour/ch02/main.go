package main
import "fmt" 
import "strconv" 
import "strings" 

func main() {
    var f32 float32 = 2.2
    var f64 float64 = 10.3456
    fmt.Println("f32 is ", f32, ", f64 is", f64)
	
	var bf bool = false
	var bt bool = true

	fmt.Println("bf is", bf, ", bt is", bt)
	
	var s1 string = "Hello"
	var s2 string = "世界"
	fmt.Println("s1 is", s1, "s2 is", s2)

	fmt.Println("s1 + s2=", s1 + s2)

	var zi int
	var zf float64
	var zb bool
	var zs string

	fmt.Println(zi, zf, zb, zs)

	i := 10
	pi := &i
	fmt.Println(*pi)	
	fmt.Println(pi)	
	i = 20;
	fmt.Println(*pi)
	
	const name = "飞雪无情"
	const (
		one = iota + 1
		two
		three
		four
	)
	
	fmt.Println(one, two, three, four, name)


	i2s := strconv.Itoa(i)
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)

	fmt.Println(strings.HasPrefix(s1, "H"))
	fmt.Println(strings.Index(s1, "o"))
	fmt.Println(strings.ToUpper(s1))
}
