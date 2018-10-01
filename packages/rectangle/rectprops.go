package rectangle

import (
	"fmt"
	"math"
)

// Cada paquete puede contener init functions, esta función no debería retornar
// nada, y puede haber más de una. Si es así se ejecutarán como se presentan en
// el compilador
func init() {
	fmt.Println("rectangle package initialized")
}

// En go las cosas a exportar deben ir con la primera letra en mayúsculas
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
