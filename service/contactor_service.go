package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// ContactorHandler 通讯录接口
func ContactorHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method == http.MethodGet {
		// 查询当前通讯录
		contactors, err := dao.Contactor.GetContactors()
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = contactors
		}
	} else if r.Method == http.MethodPost {
		// 新增联系人
		defer r.Body.Close()
		ab, err := ioutil.ReadAll(r.Body)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			contactor := new(model.ContactorModel)
			if err := json.Unmarshal(ab, contactor); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else if err := dao.Contactor.AddContactor(contactor); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			}
		}
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
