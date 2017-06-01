package provider

import (
	"net/http"

	"github.com/zero-os/0-directory/directory/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION_PROVIDER = "provider" // name of the provider collection in mongodb
)

//InitModels initializes models in DB, if required.
func InitModels() {
	index := mgo.Index{
		Key:      []string{"name"},
		Unique:   true,
		DropDups: true,
	}

	db.EnsureIndex(COLLECTION_PROVIDER, index)
}

type Manager struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func getProviderCollection(session *mgo.Session) *mgo.Collection {
	return db.GetCollection(session, COLLECTION_PROVIDER)
}

func NewManager(r *http.Request) *Manager {
	session := db.GetDBSession(r)
	return &Manager{
		session:    session,
		collection: getProviderCollection(session),
	}
}

//Create a new ResourceProvider
func (m *Manager) Create(provider *Provider) error {
	n, err := m.collection.Find(bson.M{"name": provider.Name}).Count()
	if err != nil {
		return err
	}

	if n > 0 {
		// already present
		return nil
	}

	return m.collection.Insert(provider)
}

func (m *Manager) List() ([]*Provider, error) {
	result := []*Provider{}
	err := m.collection.Find(bson.M{}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
