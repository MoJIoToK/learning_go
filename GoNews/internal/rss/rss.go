package rss

import (
	"GoNews/internal/model"
	"encoding/xml"
	"errors"
	"fmt"
	strip "github.com/grokify/html-strip-tags-go"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	ErrBodyNil = errors.New("Response body is nil")
)

// Parse - функция позволяет десериализовать RSS-поток в структуру Post. Функция возвращает слайс типа Post и ошибку.
func Parse(url string) ([]model.Post, error) {
	const operation = "rss.Parse"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	if body == nil {
		return nil, fmt.Errorf("%s: %w", operation, ErrBodyNil)
	}

	var f model.Feed

	err = xml.Unmarshal(body, &f)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	var data []model.Post
	regex, err := regexp.Compile(`[\n]{2,}[\s]+`)
	if err != nil {
		log.Println("Failed to compile regular expression")
	}

	for _, item := range f.Channel.Items {
		var post model.Post
		post.ID = len(data) + 1
		post.Title = item.Title
		post.Link = item.Link
		desc := strip.StripTags(item.Desc)
		post.Content = regex.ReplaceAllString(desc, "\n")
		item.PubDate = strings.ReplaceAll(item.PubDate, ",", "")
		t, err := time.Parse("Mon 2 Jan 2006 15:04:05 -0700", item.PubDate)
		if err != nil {
			t, err = time.Parse("Mon 2 Jan 2006 15:04:05 GMT", item.PubDate)
		}
		if err == nil {
			post.PubTime = t.Unix()
		}

		data = append(data, post)
	}

	return data, nil
}
