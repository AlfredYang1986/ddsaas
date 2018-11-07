package person

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmPersons struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Persons []BmPerson `json:"persons" jsonapi:"relationships"`
}

type BmTeachersResult struct {
	Id          string   `json:"id"`
	PersonIds []string `json:"personids"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmPersons) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmPersons) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmPersons) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmPersons) QueryId() string {
	return bd.Id
}

func (bd *BmPersons) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmPersons) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmPersons) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "persons":
		var rst []BmPerson
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(BmPerson))
		}
		bd.Persons = rst
	}
	return bd
}

func (bd BmPersons) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmPersons) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmPersons) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmPersons) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmPersons) FindMulti(req request.Request) error {

	err := bmmodel.FindMutil(req, &bd.Persons)
	for i, r := range bd.Persons {
		r.ResetIdWithId_()
		bd.Persons[i] = r
	}
	return err
}

