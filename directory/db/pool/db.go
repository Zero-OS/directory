package pool

import (
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/zero-os/0-directory/directory/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION_POOL = "resourcepool" // name of the resourcepool collection in mongodb
)

//InitModels initializes models in DB, if required.
func InitModels() {
	index := mgo.Index{
		Key:      []string{"organization"},
		Unique:   false,
		DropDups: false,
	}

	db.EnsureIndex(COLLECTION_POOL, index)
}

type Manager struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func getPoolCollection(session *mgo.Session) *mgo.Collection {
	return db.GetCollection(session, COLLECTION_POOL)
}

func NewManager(r *http.Request) *Manager {
	session := db.GetDBSession(r)
	return &Manager{
		session:    session,
		collection: getPoolCollection(session),
	}
}

//Create a new ResourcePool
func (m *Manager) Create(pool *ResourcePool) error {
	pool.UID = bson.NewObjectId()
	log.Debugln("uid", pool.UID)

	return m.collection.Insert(pool)
}

func (m *Manager) List() ([]*ResourcePool, error) {
	result := []*ResourcePool{}
	err := m.collection.Find(bson.M{}).All(&result)
	return result, err
}

func (m *Manager) ListNodes(poolID string) ([]*Node, error) {
	result := &ResourcePool{}
	//FIXME: This doesn't work for some reason, can't understand why
	// err := m.collection.FindId(bson.ObjectIdHex(poolID)).Select(bson.M{"resourcepoolcreate.nodes": 1}).All(&result)
	err := m.collection.FindId(bson.ObjectIdHex(poolID)).One(&result)
	return result.Nodes, err
}

func (m *Manager) Get(id string) (pool *ResourcePool, err error) {
	err = m.collection.FindId(bson.ObjectIdHex(id)).One(&pool)
	return
}

func (m *Manager) Update(pool *ResourcePool) (err error) {
	return m.collection.UpdateId(pool.UID, pool)
}

func (m *Manager) Delete(id string) (err error) {
	return m.collection.RemoveId(bson.ObjectIdHex(id))
}
