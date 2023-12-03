package mall

import "fmt"

type UnitType string

//region Константы

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

//endregion

// region Интерфейсы
type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

//endregion

// region Структуры

type Unit struct {
	Value float64
	T     UnitType
}

type dimensions struct {
	unitType UnitType
	length   float64 //Длина
	width    float64 //Ширина
	height   float64
}

type auto struct {
	dimensions
	brand       string
	model       string
	maxSpeed    int
	enginePower int
}

//endregion

//region Конструктор

func NewAuto(unitType UnitType, length float64, width float64, height float64, brand string, model string, maxSpeed int, enginePower int) *auto {
	return &auto{
		dimensions{
			unitType,
			length,
			width,
			height,
		},
		brand,
		model,
		maxSpeed,
		enginePower,
	}
}

//endregion

// region Методы

func (d dimensions) Length() Unit {
	return Unit{
		Value: d.length,
		T:     d.unitType,
	}
}

func (d dimensions) Width() Unit {
	return Unit{
		Value: d.width,
		T:     d.unitType,
	}
}

func (d dimensions) Height() Unit {
	return Unit{
		Value: d.height,
		T:     d.unitType,
	}
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		switch t {
		case CM:
			value *= 2.54
		case Inch:
			value /= 2.54
		}
	}
	return value
}

//region Авто

func (a auto) Brand() string { return a.brand }

func (a auto) Model() string { return a.model }

func (a auto) Dimensions() Dimensions {
	return a.dimensions
}

func (a auto) MaxSpeed() int { return a.maxSpeed }

func (a auto) EnginePower() int { return a.enginePower }

//endregion

func Print(a *auto, flag bool) {
	var dim string
	if flag {
		dim = fmt.Sprintf("Длина: %v; Ширина: %v; Высота: %v\n",
			a.Length().Get(CM), a.Width().Get(CM), a.Height().Get(CM))
	} else {
		dim = fmt.Sprintf("Длина: %v; Ширина: %v; Высота: %v\n",
			a.Length().Get(Inch), a.Width().Get(Inch), a.Height().Get(Inch))
	}
	fmt.Printf("- Бренд: %v; Модель: %v; "+
		"Максимальная скорость: %v; Мощность двигателя, лс: %v\n"+
		dim, a.brand, a.model, a.maxSpeed, a.enginePower)
}

// endregion
