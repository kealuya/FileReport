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

type PagingInfo struct {
	ProId    int32             `json:"proId,omitempty"`
	Page     int32             `json:"page,omitempty"`
	PageSize int32             `json:"pageSize,omitempty"`
	SortCol  map[string]string `json:"sortCol,omitempty"`
	Search   map[string]string `json:"search,omitempty"`
}

func GetDocFileListByCondition(paging PagingInfo) (docFiles []DocFile, funcErr error) {

	defer common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})

	var docFileMapArray = make([]map[string]any, 0)
	querySql, queryParam := makeSql(paging)

	err_Find := conf.Engine.SQL(querySql, queryParam...).Find(&docFileMapArray)
	common.ErrorHandler(err_Find)

	docFiles = make([]DocFile, 0)
	//fmt.Printf("%+v", docFileMapArray[0])
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
		//logs.Info(fmt.Sprintf("%+v", docFile))
		docFiles = append(docFiles, docFile)
	}

	return docFiles, nil

}

func makeSql(paging PagingInfo) (query any, args []any) {
	proId := paging.ProId

	var searchSql string = ""
	var searchParam []any = make([]any, 0)
	for _, v := range paging.Search {
		// name like '%tom%'
		searchSql = searchSql + " AND " + "doc_name" + " like ? "
		searchParam = append(searchParam, "%"+v+"%")

	}

	var orderSql string = ""
	var orderParam any = ""
	for k, v := range paging.SortCol {
		// time desc
		orderSql = "ORDER BY ? "
		orderParam = k + " " + v
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
	sb.WriteString(searchSql)
	sb.WriteString(orderSql)
	sb.WriteString(limitSql)

	paramArray := make([]any, 0)
	paramArray = append(paramArray, proId)
	paramArray = append(paramArray, searchParam...)
	paramArray = append(paramArray, orderParam)
	paramArray = append(paramArray, limitParam...)

	querySql, queryParam, queryError := ToSQL(Expr(sb.String(), paramArray...))
	common.ErrorHandler(queryError)

	return querySql, queryParam
}

func GetDocCountByProId(paging PagingInfo) (count int, funcErr error) {

	defer common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})
	// 清空limit属性，做到全搜索
	paging.Page = 0
	querySql, queryParam := makeSql(paging)
	var docFileMapArray = make([]map[string]any, 0)
	err_Count := conf.Engine.SQL(querySql, queryParam...).Find(&docFileMapArray)
	common.ErrorHandler(err_Count)

	return len(docFileMapArray), nil

}
