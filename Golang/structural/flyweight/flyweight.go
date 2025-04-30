// 享元模式

package flyweight

import "fmt"

// ColorFlyweightFactory 享元工厂
type ColorFlyweightFactory struct {
	maps map[string]*ColorFlyweight
}

var colorFactory *ColorFlyweightFactory

func GetColorFlyweightFactory() *ColorFlyweightFactory {
	if colorFactory == nil {
		colorFactory = &ColorFlyweightFactory{
			maps: make(map[string]*ColorFlyweight),
		}
	}
	return colorFactory
}

func (f *ColorFlyweightFactory) Get(filename string) *ColorFlyweight {
	color := f.maps[filename]
	if color == nil {
		color = NewColorFlyweight(filename)
		f.maps[filename] = color
	}

	return color
}

type ColorFlyweight struct {
	data string
}

// NewColorFlyweight 存储color对象
func NewColorFlyweight(filename string) *ColorFlyweight {
	// Load color file
	data := fmt.Sprintf("color data %s", filename)
	return &ColorFlyweight{
		data: data,
	}
}

type ColorViewer struct {
	*ColorFlyweight
}

func NewColorViewer(name string) *ColorViewer {
	color := GetColorFlyweightFactory().Get(name)
	return &ColorViewer{
		ColorFlyweight: color,
	}
}

func PrintIsSameAddress(a, b interface{}, aName string, bName string) {
	fmt.Printf("%v == %v: %v\n", aName, bName, a == b)
}
