package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	/*
		IntValue, ok := value.(int) // метод извлечения типа данных для более ветвистой структуры
		if ok {
			color.Red("Код ошибки: %d", IntValue)
			return
		}
		StrValue, ok := value.(string)
		if ok {
			color.Red(StrValue)
			return
		}
		ErrValue, ok := value.(error)
		if ok {
			color.Red(ErrValue.Error())
			return
		}
	*/
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	default:
		color.Red("Неизвестный тип ошибки")
	}

} // Generic
func sum[T int | float64 | float32 | int16 | int32](a, b T) T {
	/*switch d := any(a).(type) { //any for work // type switch for generic
	case string:
		fmt.Println(d)
	}
	*/
	return a + b
}

/*type List[T any] struct { // Generic on Struct
	elements []T
}
*/
