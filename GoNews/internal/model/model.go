package model

import "encoding/xml"

// Post - модель публикации в БД, получаемая из RSS ленты
type Post struct {
	ID      int
	Title   string
	Content string
	PubTime int64
	Link    string
}

// Feed - структура для RSS канала
type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel - структура для RSS-потока
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

// Item - структура для записи одной публикации
type Item struct {
	Title   string `xml:"title"`
	Desc    string `xml:"description"`
	PubTime int64  `xml:"pubDate"`
	Link    string `xml:"link"`
}
