package mongo

import (
	"github.com/pjvds/counter"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type mgoCountService struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func NewCountService(url string) (counter.CountService, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	collection := session.DB("").C("counter")
	return &mgoCountService{
		session:    session,
		collection: collection,
	}, nil
}

func (m *mgoCountService) Increase(name counter.Name) error {
	col := m.collection
	_, err := col.Upsert(bson.M{"name": name}, bson.M{"$inc": bson.M{"count": 1}})

	return err
}

func (m *mgoCountService) Get(name counter.Name) (int, error) {
	var result bson.M
	col := m.collection

	err := col.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return 0, nil
		}
		return 0, err
	}
	return result["count"].(int), nil
}
