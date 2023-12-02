package electronic

//region Интерфейсы

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonCount() int
}

type Smartphone interface {
	OS() string
}

//endregion

//region Структуры

type ApplePhone struct {
	os    string
	brand string
	types string
	model string
}

type AndroidPhone struct {
	os    string
	brand string
	types string
	model string
}

type RadioPhone struct {
	countButton int
	brand       string
	types       string
	model       string
}

//endregion

//region Методы

func (a ApplePhone) Brand() string {
	a.brand = "Apple"
	return a.brand
}

func (a ApplePhone) Model() string {
	return a.model
}

func (a ApplePhone) Type() string {
	a.types = "smartphone"
	return a.types
}

func (a ApplePhone) OS() string {
	return a.os
}

func (a AndroidPhone) Brand() string {
	return a.brand
}

func (a AndroidPhone) Model() string {
	return a.model
}

func (a AndroidPhone) Type() string {
	a.types = "smartphone"
	return a.types
}

func (a AndroidPhone) OS() string {
	return a.os
}

func (r RadioPhone) Brand() string {
	return r.brand
}

func (r RadioPhone) Model() string {
	return r.model
}

func (r RadioPhone) Type() string {
	r.types = "station"
	return r.types
}

func (r RadioPhone) ButtonCount() int {
	return r.countButton
}

//endregion

// region Конструкторы

func NewApplePhone(os string, model string) *ApplePhone {
	return &ApplePhone{os: os, model: model}
}

func NewAndroidPhone(os string, brand string, model string) *AndroidPhone {
	return &AndroidPhone{os: os, brand: brand, model: model}
}

func NewRadioPhone(countButton int, brand string, model string) *RadioPhone {
	return &RadioPhone{countButton: countButton, brand: brand, model: model}
}

//endregion
