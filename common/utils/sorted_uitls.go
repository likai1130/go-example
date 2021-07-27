package utils

import (
	"reflect"
	"sort"
)

/**
通用排序
构体排序，必须重写数组Len() Swap() Less()函数
*/
type body_wrapper struct {
	Bodys []interface{}
	by    func(p, q *interface{}) bool //内部Less()函数会用到
}

type SortBodyBy func(p, q *interface{}) bool //定义一个函数类型

//数组长度Len()
func (acw body_wrapper) Len() int {
	return len(acw.Bodys)
}

//元素交换
func (acw body_wrapper) Swap(i, j int) {
	acw.Bodys[i], acw.Bodys[j] = acw.Bodys[j], acw.Bodys[i]
}

//比较函数，使用外部传入的by比较函数
func (acw body_wrapper) Less(i, j int) bool {
	return acw.by(&acw.Bodys[i], &acw.Bodys[j])
}

//自定义排序字段，参考SortBodyByCreateTime中的传入函数
func SortBody(bodys []interface{}, by SortBodyBy) {
	sort.Stable(body_wrapper{bodys, by})
}

//降序排列
func DescSortBodyByFieldName(bodys []interface{}, fieldName string) {
	sort.SliceStable(bodys, func(pre, next int) bool {
		valuePre_ := reflect.ValueOf(bodys[pre]).Elem()
		valueNext_ := reflect.ValueOf(bodys[next]).Elem()

		filedPre := valuePre_.FieldByName(fieldName)
		filedNex := valueNext_.FieldByName(fieldName)

		if filedPre.Kind().String() == "invalid" {
			return filedPre.String() > filedNex.String()
		}

		typeName := filedPre.Type().Name()
		switch typeName {
		case "int":
			return filedPre.Int() > filedNex.Int()
		default:
			return filedPre.String() > filedNex.String()
		}
	})

}

//升序排列
func AscendSortBodyFieldName(bodys []interface{}, fieldName string) {
	sort.SliceStable(bodys, func(pre, next int) bool {
		valuePre_ := reflect.ValueOf(bodys[pre]).Elem()
		valueNext_ := reflect.ValueOf(bodys[next]).Elem()

		filedPre := valuePre_.FieldByName(fieldName)
		filedNex := valueNext_.FieldByName(fieldName)

		if filedPre.Kind().String() == "invalid" {
			return filedPre.String() < filedNex.String()
		}

		typeName := filedPre.Type().Name()
		switch typeName {
		case "int":
			return filedPre.Int() < filedNex.Int()
		default:
			return filedPre.String() < filedNex.String()
		}
	})
}
