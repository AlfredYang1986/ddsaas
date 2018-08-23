package student

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmmodel"
)

type BMStudents struct {
	Id  				string        		`json:"id"`
	Id_ 				bson.ObjectId 		`bson:"_id"`

	Students			[]BMStudent			`json:"students" jsonapi:"relationships"`
}

func (bd *BMStudents) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Students)
	for i, r := range bd.Students {
		r.ResetIdWithId_()
		bd.Students[i] = r
	}
	return err
}
