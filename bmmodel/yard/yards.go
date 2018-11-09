package yard

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmYards struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Yards []BmYard `json:"Yards" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmYards) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmYards) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmYards) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmYards) QueryId() string {
	return bd.Id
}

func (bd *BmYards) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmYards) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmYards) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmYards) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmYards) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmYards) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmYards) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmYards) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Yards)
	for i, r := range bd.Yards {
		r.ResetIdWithId_()
		bd.Yards[i] = r
	}
	return err
}
