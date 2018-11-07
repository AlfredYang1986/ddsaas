package teacher

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmTeacherProps struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	TeacherProps []BMTeacherProp `json:"teacherProps" jsonapi:"relationships"`
}

//type BmTeacherPropResults struct {
//	Id          string   `json:"id"`
//	TeacherIds []string `json:"teacherids"`
//}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmTeacherProps) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmTeacherProps) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmTeacherProps) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmTeacherProps) QueryId() string {
	return bd.Id
}

func (bd *BmTeacherProps) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmTeacherProps) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmTeacherProps) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "teacherProps":
		var rst []BMTeacherProp
		for _, item := range v.([]interface{}) {
			rst = append(rst, item.(BMTeacherProp))
		}
		bd.TeacherProps = rst
	}
	return bd
}

func (bd BmTeacherProps) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmTeacherProps) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmTeacherProps) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmTeacherProps) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmTeacherProps) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.TeacherProps)
	for i, r := range bd.TeacherProps {
		r.ResetIdWithId_()
		bd.TeacherProps[i] = r
	}
	return err
}