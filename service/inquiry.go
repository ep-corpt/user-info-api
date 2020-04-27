package service

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"user-info-api/common"
	"user-info-api/entity"
)

const(
	id = "id"
	idIsRequired = "ID is required."
	idIsInvalid = "ID is invalid."
)

type inquiryTask struct{
	db *gorm.DB
	id string
}

func (h *Handler)Inquiry(w http.ResponseWriter, r *http.Request){
	id := r.Header.Get(id)
	if len(id) == 0{
		common.RespErr(w, idIsRequired)
		return
	}

	var t inquiryTask
	t.initTask(h, id)
	t.execute(w)
}

func (t *inquiryTask)initTask(h *Handler, id string){
	t.db = h.db
	t.id = id
}

func (t *inquiryTask)execute(w http.ResponseWriter){
	rs := t.db.Where("username = ?", t.id).Find(&entity.UserInfo{})
	if rs.Error == nil && rs.Value != nil{
		u := rs.Value.(*entity.UserInfo)
		common.RespSuccess(w, &u)
	}else{
		common.RespErr(w, idIsInvalid)
	}
}


