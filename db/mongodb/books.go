package mongodb

import (
	"context"
	"github.com/guitarpawat/bookshelf/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var booksRepo *BooksRepo

const DefaultBooksCollectionName = "books"

type BooksRepo struct {
	CollectionName string
}

func (r *BooksRepo) GetById(id string) (dto.Book, error) {
	ctx, _ := context.WithTimeout(context.Background(), database.Timeout)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.Book{}, err
	}

	findResult := database.getDB().Collection(r.CollectionName).FindOne(ctx, bson.M{"_id": objectId})
	if findResult.Err() != nil {
		return dto.Book{}, findResult.Err()
	}

	result := bson.M{}
	err = findResult.Decode(&result)
	if err != nil {
		return dto.Book{}, err
	}

	authorBson, ok := result["author"].(bson.A)
	var author []string
	if !ok {
		author = nil
	}
	author = make([]string, len(authorBson))
	for i := range authorBson {
		author[i] = authorBson[i].(string)
	}

	tagsBson, ok := result["tags"].(bson.A)
	var tags []string
	if !ok {
		tags = nil
	}
	tags = make([]string, len(tagsBson))
	for i := range tagsBson {
		tags[i] = tagsBson[i].(string)
	}

	volBson, ok := result["type"].(bson.A)
	var vol []int
	if !ok {
		vol = nil
	}
	vol = make([]int, len(volBson))
	for i := range volBson {
		vol[i] = volBson[i].(int)
	}

	return dto.Book{
		ID:      id,
		Title:   result["title"].(string),
		Edition: result["edition"].(string),
		Author:  author,
		Tags:    tags,
		Type:    dto.BookType(result["type"].(int32)),
		Status:  dto.BookStatus(result["status"].(int32)),
		Volume:  vol,
		Owner:   result["owner"].(string),
		AddTime: objectId.Timestamp(),
	}, nil
}

func (r *BooksRepo) Save(book dto.Book) error {
	ctx, _ := context.WithTimeout(context.Background(), database.Timeout)
	objectId := primitive.NewObjectID()
	_, err := database.getDB().Collection(r.CollectionName).InsertOne(ctx, bson.M{
		"_id":     objectId,
		"title":   book.Title,
		"edition": book.Edition,
		"author":  book.Author,
		"tags":    book.Tags,
		"type":    book.Type,
		"status":  book.Status,
		"volume":  book.Volume,
		"owner":   book.Owner,
	})
	return err
}

func newBooksRepo(collectionName string) *BooksRepo {
	return &BooksRepo{CollectionName: collectionName}
}
