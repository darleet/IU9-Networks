package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

func main() {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL("https://www.aviaport.ru/digest/press-releases/rss/")
	if err != nil {
		panic(err)
	}
	fmt.Println(feed.Title)
	fmt.Println(feed.Description)
	fmt.Printf("Последнее обновление: %s\n", feed.Updated)
	fmt.Printf("Статей: %d\n", feed.Len())
	for k, v := range feed.Items {
		fmt.Println()
		fmt.Printf("Номер статьи: %d\n", k+1)
		fmt.Println(v.Title)
		fmt.Println(v.Link)
	}
}
