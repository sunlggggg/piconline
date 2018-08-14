# struct
数据的组合

## 首字母大小写控制可见性

## 内部方法
``` go
type Rect struct {
	width, length float64
}

func (rect Rect) area() float64 {
	return rect.width * rect.length
}

func main() {
	var rect = new(Rect)
	rect.width = 100
	rect.length = 200
	fmt.Println("Width:", rect.width, "Length:", rect.length,
		"Area:", rect.area())
}


```
