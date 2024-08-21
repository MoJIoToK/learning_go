package model

type Order struct {
	ID              int       // номер заказа
	IsOpen          bool      // открыт/закрыт
	DeliveryTime    int64     //срок доставки
	DeliveryAddress string    // адрес доставки
	Products        []Product // состав заказа
}

type Product struct {
	ID    int     // артикул товара
	Name  string  //наименование товара
	Price float64 //цена
}
