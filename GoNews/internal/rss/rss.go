package rss

import (
	"GoNews/internal/model"
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
)

func Parse() ([]model.Post, error) {
	url := "https://habr.com/ru/rss/hub/go/all/?fl=ru"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var f model.Feed

	d := xml.NewDecoder(bytes.NewReader(body))
	err = d.Decode(&f)
	if err != nil {
		panic(err)
	}

	//err = xml.Unmarshal(body, &f)
	//if err != nil {
	//	panic(err)
	//}

	//var data []model.Post
	//
	//for _, item := range f.Channel.Items {
	//	var post model.Post
	//	post.Title = item.Title
	//	post.Link = item.Link
	//	post.Content = item.Desc
	//
	//	data = append(data, post)
	//}

	return f, nil
}
