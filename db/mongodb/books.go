package mongodb

import (
	"context"
	"github.com/guitarpawat/bookshelf/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var booksRepo *BooksRepo

const DefaultBooksCollectionName = "books"
const MinObjectId = "000000000000000000000000"

type BooksRepo struct {
	CollectionName string
}

type book struct {
	dto.Book `bson:",inline"`
	ID       primitive.ObjectID `bson:"_id"`
}

func (r *BooksRepo) GetPaginationSortByTimeDesc(limit int, lastId string) ([]dto.Book, string, error) {
	var objectId primitive.ObjectID
	var err error
	if lastId == "" {
		lastId = MinObjectId
	}
	objectId, err = primitive.ObjectIDFromHex(lastId)
	if err != nil {
		return nil, "", err
	}

	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetSort(bson.D{{"_id", -1}})

	ctx, _ := context.WithTimeout(context.Background(), database.Timeout)
	cur, err := database.getDB().Collection(r.CollectionName).Find(ctx, bson.M{"_id": bson.M{"$gt": objectId}}, opts)
	if err != nil {
		return nil, "", err
	}

	books := make([]dto.Book, 0)
	for cur.Next(context.TODO()) {
		b := &book{}
		err = cur.Decode(&b)
		if err != nil {
			return nil, "", err
		}
		b.Book.ID = b.ID.Hex()
		b.Book.AddTime = b.ID.Timestamp()
		books = append(books, b.Book)
	}

	currentLast := ""
	if len(books) > 0 {
		currentLast = books[len(books)-1].ID
	}

	return books, currentLast, nil
}

func (r *BooksRepo) GetById(id string) (dto.Book, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.Book{}, err
	}

	ctx, _ := context.WithTimeout(context.Background(), database.Timeout)
	findResult := database.getDB().Collection(r.CollectionName).FindOne(ctx, bson.M{"_id": objectId})
	if findResult.Err() != nil {
		return dto.Book{}, findResult.Err()
	}

	result := &book{}
	err = findResult.Decode(&result)
	if err != nil {
		return dto.Book{}, err
	}
	result.Book.ID = result.ID.Hex()
	result.Book.AddTime = result.ID.Timestamp()
	return result.Book, nil
}

func (r *BooksRepo) Save(b dto.Book) (string, error) {
	bookData := BookToBson(book{
		Book: b,
		ID:   primitive.NewObjectID(),
	})

	ctx, _ := context.WithTimeout(context.Background(), database.Timeout)
	result, err := database.getDB().Collection(r.CollectionName).InsertOne(ctx, bookData)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func newBooksRepo(collectionName string) *BooksRepo {
	return &BooksRepo{CollectionName: collectionName}
}

func BookToBson(b book) bson.M {
	return bson.M{
		"_id":     b.ID,
		"title":   b.Title,
		"edition": b.Edition,
		"author":  b.Author,
		"tags":    b.Tags,
		"type":    b.Type,
		"status":  b.Status,
		"volume":  b.Volume,
	}
}
