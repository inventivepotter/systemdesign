package main

import (
	"errors"
	"fmt"
	"time"
)

var Id int32

type Post struct {
	Id        int32
	Content   string
	Author    string
	createdAt time.Time
	updatedAt time.Time
}

type Question struct {
	Post
	tags []string
}

type mongoDB struct{}

func (m mongoDB) Save(i interface{}) error {
	fmt.Println(i)
	return nil
}

func validateQuestion(content string) error {
	//Validating Content
	if len(content) > 100 {
		return errors.New("Content length exceeded")
	}
	return nil
}

func createQuestion(content string, author string, tags []string) (Question, error) {
	//Creating the Question Object
	Id++
	return Question{
		Post: Post{
			Id:        Id,
			Content:   content,
			Author:    author,
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		tags: tags,
	}, nil
}

func (q Question) Save() error {
	//Saving to database
	mongoDBConn := mongoDB{}
	if err := mongoDBConn.Save(q); err != nil {
		return err
	}
	return nil
}

func PostQuestion(content string, author string, tags []string) error {
	if err := validateQuestion(content); err != nil {
		return err
	}

	q, err := createQuestion(content, author, tags)
	if err != nil {
		return err
	}

	if err := q.Save(); err != nil {
		return err
	}

	return nil
}

func main() {
	PostQuestion("What is an Ape?", "Adam G", []string{"wiered", "animal"})
}
