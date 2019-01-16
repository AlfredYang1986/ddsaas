package courseunit

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
	"sort"
)

type BmCourseUnits struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	CourseUnits []BmCourseUnit `json:"CourseUnits" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmCourseUnits) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmCourseUnits) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmCourseUnits) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmCourseUnits) QueryId() string {
	return bd.Id
}

func (bd *BmCourseUnits) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmCourseUnits) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmCourseUnits) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmCourseUnits) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmCourseUnits) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmCourseUnits) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmCourseUnits) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmCourseUnits) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.CourseUnits)
	for i, r := range bd.CourseUnits {
		r.ResetIdWithId_()
		r.ReSetProp()
		bd.CourseUnits[i] = r
	}
	return err
}

func (bd *BmCourseUnits) SortByStartDate(increasing bool) error {
	courseUnits := bd.CourseUnits
	if courseUnits == nil {
		return nil
	}
	if increasing {
		sort.Sort(BmCourseUnitsWrapper{courseUnits, func(cu1, cu2 *BmCourseUnit) bool {
			return cu1.StartDate < cu2.StartDate //按开始时间递增排序
		}})
	} else {
		sort.Sort(BmCourseUnitsWrapper{courseUnits, func(cu1, cu2 *BmCourseUnit) bool {
			return cu1.StartDate > cu2.StartDate //按开始时间递减排序
		}})
	}

	return nil
}

func (bd *BmCourseUnits) SortByEndDate(increasing bool) error {
	courseUnits := bd.CourseUnits
	if courseUnits == nil {
		return nil
	}
	if increasing {
		sort.Sort(BmCourseUnitsWrapper{courseUnits, func(cu1, cu2 *BmCourseUnit) bool {
			return cu1.EndDate < cu2.EndDate //按结束时间递增排序
		}})
	} else {
		sort.Sort(BmCourseUnitsWrapper{courseUnits, func(cu1, cu2 *BmCourseUnit) bool {
			return cu1.EndDate > cu2.EndDate //按结束时间递减排序
		}})
	}

	return nil
}

type BmCourseUnitsWrapper struct {
	courseUnits []BmCourseUnit
	sortBy      func(cu1, cu2 *BmCourseUnit) bool
}

func (bd BmCourseUnitsWrapper) Len() int {
	return len(bd.courseUnits)
}

func (bd BmCourseUnitsWrapper) Swap(i, j int) {
	bd.courseUnits[i], bd.courseUnits[j] = bd.courseUnits[j], bd.courseUnits[i]
}

func (bd BmCourseUnitsWrapper) Less(i, j int) bool {
	return bd.sortBy(&bd.courseUnits[i], &bd.courseUnits[j])
}
