package models

import (
	"FileReport/common"
	"FileReport/conf"
	"FileReport/entity"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gohouse/t"
	"log"
	"reflect"
	"strings"
	. "xorm.io/builder"
)

func InsertDoc(doc entity.Doc) (docId int32, funcErr error) {

	common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})

	num, err_InsertOne := conf.Engine.InsertOne(doc)
	common.ErrorHandler(err_InsertOne)
	if num < 1 {
		logs.Error("Doc表存入数据失败")
		log.Panicln()
	}
	var resultQuery []map[string][]byte
	resultQuery, err_Query := conf.Engine.Query(`select max(doc_id) as id from Doc`)
	if err_Query != nil {
		logs.Error("Doc表获取最大DocId数据失败")
		log.Panicln()
	}
	fmt.Println(resultQuery)

	return t.New(resultQuery[0]["id"]).Int32(), nil

}

func UpdateDoc(doc entity.Doc) (funcErr error) {

	common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})

	num, err_Update := conf.Engine.ID(doc.DocId).Update(doc)
	common.ErrorHandler(err_Update)
	if num < 1 {
		logs.Error("Doc表更新数据失败")
		log.Panicln("Doc表更新数据失败")
	}
	return nil
}

func InsertFile(file entity.File) (funcErr error) {

	defer common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})
	num, err_InsertOne := conf.Engine.InsertOne(file)
	common.ErrorHandler(err_InsertOne)
	if num < 1 {
		logs.Error("Doc表存入数据失败")
		log.Panicln()
	}
	return nil

}

func GetFileVersion(docId string) (version string, funcErr error) {

	defer common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})

	var mapArray = make([]map[string]any, 0)
	err := conf.Engine.SQL(`select  max(version) as version from File where doc_id = ?`, t.New(docId).Int32()).Find(&mapArray)
	common.ErrorHandler(err)

	if len(mapArray) == 0 {
		log.Panicln("no record")
	}

	return mapArray[0]["version"].(string), nil

}

type PagingInfo struct {
	ProId    int32             `json:"proId,omitempty"`
	Page     int32             `json:"page,omitempty"`
	PageSize int32             `json:"pageSize,omitempty"`
	SortCol  map[string]string `json:"sortCol,omitempty"`
	Search   map[string]string `json:"search,omitempty"`
}

func GetDocFileListByCondition(paging PagingInfo, userInfo *entity.User) (docFiles []DocFile, funcErr error) {

	defer common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})

	var docFileMapArray = make([]map[string]any, 0)
	querySql, queryParam := makeSql(paging, userInfo)
	err_Find := conf.Engine.SQL(querySql, queryParam...).Find(&docFileMapArray)
	fmt.Printf("%+v", err_Find)
	common.ErrorHandler(err_Find)

	docFiles = make([]DocFile, 0)
	for _, docFileMap := range docFileMapArray {
		docFile := DocFile{}
		for k, v := range docFileMap {
			docFileType := reflect.TypeOf(docFile)
			docFileValue := reflect.ValueOf(&docFile)
			fieldName := common.ConvertToCaseAlpha(k)
			if _, ok := docFileType.FieldByName(fieldName); ok {
				docFileValue.Elem().FieldByName(fieldName).SetString(t.New(v).String())
			}
		}
		// 追加 updateContentList 处理
		docId := t.New(docFileMap["doc_id"]).String()

		updateContentList := make([]map[string]string, 0)
		err_updateContentList := conf.Engine.SQL("select version_show as versionShow, update_content as updateContent from File where doc_id = ? order by update_date desc limit 5",
			docId).Find(&updateContentList)
		if err_updateContentList != nil {
			common.ErrorHandler(err_updateContentList)
		}

		docFile.UpdateContentList = updateContentList
		docFiles = append(docFiles, docFile)
	}

	return docFiles, nil

}

func makeSql(paging PagingInfo, user *entity.User) (query any, args []any) {
	proId := paging.ProId

	var searchSql string = ""
	var searchParam []any = make([]any, 0)
	for k, v := range paging.Search {
		// name like '%tom%'

		kv := ""
		if strings.Contains(k, "docName") {
			kv = "doc_name"
		} else if strings.Contains(k, "updateUser") {
			kv = "update_user"
		} else if strings.Contains(k, "owner") {
			kv = "owner"
		}

		searchSql = searchSql + kv + " like ? " + " OR "
		if kv != "" {
			searchParam = append(searchParam, "%"+v+"%")
		}
	}
	if searchSql != "" {
		searchSql = " AND ( " + searchSql[:len(searchSql)-4] + " ) "
	}

	var orderSql string = ""
	for k, v := range paging.SortCol {
		// column like createDate
		// "ascending"  "descending"
		// orderby不支持？，需要自己处理sql注入，可以通过白名单解决，
		// 前端预定符号代表字段名及排序类型，后端根据符号拼接，避免注入
		orderSql = "ORDER BY "

		kv := "createDate"
		if strings.Contains(k, "updateDate") {
			kv = "updateDate"
		}

		vv := "asc"
		if strings.Contains(v, "desc") {
			vv = "desc"
		}
		orderSql = orderSql + common.ConvertHumpNameToSnakeCase(kv) + " " + vv + " "
	}
	var limitSql string = ""
	var limitParam = []any{}
	if paging.Page != 0 {
		//limit (curPage-1)*pageSize,pageSize
		limit1 := (paging.Page - 1) * paging.PageSize
		limit2 := paging.PageSize
		limitSql = "LIMIT ? , ? "
		limitParam = []any{limit1, limit2}
	}
	sb := strings.Builder{}
	sb.WriteString(`
						SELECT
							* ,u1.user_name as owner,u2.user_name as update_user
						FROM
							Doc doc
							INNER JOIN ( SELECT MAX( version ) AS version,* FROM File GROUP BY doc_id ) f ON doc.doc_id = f.doc_id 
							LEFT  JOIN User u1 ON doc.owner_id = u1.phone_number 
							LEFT  JOIN User u2 ON f.update_user_id = u2.phone_number 
						WHERE
							doc.pro_id = ? 
					`)
	// 游客的场合
	if user == nil {
		sb.WriteString(` AND doc.is_discard = 'false' AND doc.is_release = 'true' `)
	} else {
		// 社内用户的场合
		sb.WriteString(` AND ( (doc.is_discard = 'true' and doc.owner_id = '` + user.PhoneNumber + `') OR doc.is_discard = 'false'  ) `)
	}

	sb.WriteString(searchSql)
	sb.WriteString(orderSql)
	sb.WriteString(limitSql)

	paramArray := make([]any, 0)
	paramArray = append(paramArray, proId)
	if len(searchParam) > 0 {
		paramArray = append(paramArray, searchParam...)
	}

	if len(limitParam) > 0 {
		paramArray = append(paramArray, limitParam...)
	}

	querySql, queryParam, queryError := ToSQL(Expr(sb.String(), paramArray...))
	common.ErrorHandler(queryError)

	return querySql, queryParam
}

func GetDocCountByProId(paging PagingInfo, userInfo *entity.User) (count int, funcErr error) {

	defer common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})
	// 清空limit属性，做到全搜索
	paging.Page = 0
	querySql, queryParam := makeSql(paging, userInfo)
	var docFileMapArray = make([]map[string]any, 0)
	err_Count := conf.Engine.SQL(querySql, queryParam...).Find(&docFileMapArray)
	common.ErrorHandler(err_Count)

	return len(docFileMapArray), nil

}
