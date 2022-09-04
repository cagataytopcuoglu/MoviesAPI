package movie

import (
	"MovieAPI/pkg/log"
	"MovieAPI/pkg/mongoHelper"
	"MovieAPI/pkg/pagination"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	FindOne(params map[string]string) (*Movie, error)
	GetById(id string) (*Movie, error)
	Find(pageOptions *pagination.Pages) ([]Movie, error)
	Create(entity *Movie) error
	Update(entity *Movie) error
	Delete(id string) error
}

type repository struct {
	collection *mongo.Collection
}

func (r *repository) FindOne(params map[string]string) (movie *Movie, err error) {

	query := mongoHelper.BuildQuery(params)
	r.collection.FindOne(context.TODO(), query).Decode(&movie)
	return movie, nil
}
func (r *repository) Find(pageOptions *pagination.Pages) (results []Movie, err error) {

	findOptions := mongoHelper.SetFindOptions(pageOptions)

	cur, err := r.collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Logger.Error(err)
		return results, err
	}
	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem Movie
		err := cur.Decode(&elem)
		if err != nil {
			log.Logger.Fatal(err)
		}
		results = append(results, elem)
	}
	return results, nil
}
func (r *repository) Create(entity *Movie) error {

	_, err := r.collection.InsertOne(context.TODO(), entity)

	return err
}
func (r *repository) Update(entity *Movie) error {

	replaceOptions := options.Replace()
	id := entity.BaseModel.Id
	entity.BaseModel.Id = ""
	_, err := r.collection.ReplaceOne(context.TODO(), mongoHelper.CastToId(id), entity, replaceOptions)

	return err
}

func (r *repository) Delete(id string) error {

	_, err := r.collection.DeleteOne(context.TODO(), mongoHelper.CastToId(id))

	return err
}

func (r *repository) GetById(id string) (*Movie, error) {
	return r.FindOne(map[string]string{"_id": id})
}

func NewRepository(db *mongo.Database) Repository {

	col := db.Collection("movies")
	return &repository{
		collection: col,
	}
}
