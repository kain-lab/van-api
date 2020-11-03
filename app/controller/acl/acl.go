package acl

import (
	curd "github.com/kainonly/iris-curd"
	"github.com/kainonly/iris-helper/res"
	"github.com/kainonly/iris-helper/validate"
	"gorm.io/gorm"
	"van-api/app/cache"
	"van-api/app/model"
	"van-api/types"
)

type Controller struct {
}

type originListsBody struct {
	curd.OriginListsBody
}

func (c *Controller) PostOriginlists(body *originListsBody, mode *curd.Curd) interface{} {
	return mode.
		Originlists(model.Acl{}, body.OriginListsBody).
		OrderBy([]string{"create_time desc"}).
		Exec()
}

type listsBody struct {
	curd.ListsBody
}

func (c *Controller) PostLists(body *listsBody, mode *curd.Curd) interface{} {
	return mode.
		Lists(model.Acl{}, body.ListsBody).
		OrderBy([]string{"create_time desc"}).
		Exec()
}

type getBody struct {
	curd.GetBody
}

func (c *Controller) PostGet(body *getBody, mode *curd.Curd) interface{} {
	return mode.
		Get(model.Acl{}, body.GetBody).
		Exec()
}

type addBody struct {
	Keyid  string     `validate:"required"`
	Name   types.JSON `validate:"required"`
	Read   string
	Write  string
	Status bool
}

func (c *Controller) PostAdd(body *addBody, mode *curd.Curd, cache *cache.Model) interface{} {
	errs := validate.Make(body, nil)
	if errs != nil {
		return res.Error(errs)
	}
	data := model.Acl{
		Keyid:  body.Keyid,
		Name:   body.Name,
		Read:   body.Read,
		Write:  body.Write,
		Status: body.Status,
	}
	return mode.
		Add().
		Exec(&data)
}

type editBody struct {
	curd.EditBody
	Keyid  string     `validate:"required_if=switch false"`
	Name   types.JSON `validate:"required_if=switch false"`
	Read   string
	Write  string
	Status bool
}

func (c *Controller) PostEdit(body *editBody, mode *curd.Curd, cache *cache.Model) interface{} {
	errs := validate.Make(body, nil)
	if errs != nil {
		return res.Error(errs)
	}
	data := model.Acl{
		Keyid:  body.Keyid,
		Name:   body.Name,
		Read:   body.Read,
		Write:  body.Write,
		Status: body.Status,
	}
	return mode.
		Edit(model.Acl{}, body.EditBody).
		After(func(tx *gorm.DB) error {
			clearcache(cache)
			return nil
		}).
		Exec(data)
}

type deleteBody struct {
	curd.DeleteBody
}

func (c *Controller) PostDelete(body *deleteBody, mode *curd.Curd, cache *cache.Model) interface{} {
	return mode.
		Delete(model.Acl{}, body.DeleteBody).
		After(func(tx *gorm.DB) error {
			clearcache(cache)
			return nil
		}).
		Exec()
}

type validedkeyBody struct {
	Keyid string `validate:"required"`
}

func (c *Controller) PostValidedkey(body *validedkeyBody, db *gorm.DB) interface{} {
	errs := validate.Make(body, nil)
	if errs != nil {
		return res.Error(errs)
	}
	var count int64
	db.Model(&model.Acl{}).
		Where("keyid = ?", body.Keyid).
		Count(&count)

	return res.Data(count != 0)
}

func clearcache(cache *cache.Model) {
	cache.AclClear()
	cache.RoleClear()
}