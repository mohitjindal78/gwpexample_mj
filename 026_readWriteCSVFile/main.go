package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	//create csv file
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "hello World!", Author: "Mohit Jindal"},
		Post{Id: 2, Content: "Purple Cow", Author: "Seth Godwin"},
		Post{Id: 3, Content: "My Job Went to India", Author: "Chad Flower"},
		Post{Id: 4, Content: "The subtle Art of not giving a fuck", Author: "Mark Manson"},
	}

	csvWriter := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := csvWriter.Write(line)
		if err != nil {
			panic(err)
		}
	}
	csvWriter.Flush()

	//reading a CSV file
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
	record, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		fmt.Println(item)
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts)
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
	fmt.Println(posts[1].Id)
	fmt.Println(posts[1].Content)
	fmt.Println(posts[1].Author)
	fmt.Println(posts[2].Id)
	fmt.Println(posts[2].Content)
	fmt.Println(posts[2].Author)
	fmt.Println(posts[3].Id)
	fmt.Println(posts[3].Content)
	fmt.Println(posts[3].Author)
}
