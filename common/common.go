package common

import (
	"fmt"
	"reflect"
)

// type every arr to []interface
func ArrToObjs(src interface{}) (des []interface{}) {
	data := reflect.ValueOf(src)
	num := data.Len()
	for i := 0; i < num; i++ {
		des = append(des, data.Index(i))
	}
	return
}

func FindElemInObjs(objs interface{}, item interface{}) int {

	// 空值
	if objs == nil {
		return -1
	}

	// 不为切片
	objsValues := reflect.ValueOf(objs)
	objsKind := objsValues.Kind()

	if objsKind != reflect.Slice {
		return -1
	}

	// nil查找
	for i := 0; i < objsValues.Len(); i++ {
		elemData := objsValues.Index(i)

		if elemData.CanInterface() {
			if reflect.DeepEqual(item, elemData.Interface()) {
				return i
			}
		} else {
			fmt.Println("No CanInterface")
		}
	}

	return -1
}

// // find item in objs, return index
// // -1 is not find  // 一般不比较nil
// func FindElemInObjs(objs interface{}, item interface{}) int {

// 	itemData := reflect.ValueOf(item)
// 	itemType := itemData.Type()
// 	itemTypeStr := itemType.String()

// 	data := reflect.ValueOf(objs)
// 	dataType := data.Type()

// 	// 非inteface 先判断类型是否一致
// 	if "[]interface {}" != dataType.String() {
// 		if itemTypeStr != dataType.Elem().String() {
// 			pt.Ptln("类型不一致")
// 			return -1
// 		}
// 	}

// 	// 接下来判断内存是否一致
// 	//  pt.Ptln("type2 equal?:", dataType.Elem() == itemType) 不能用此判断，有可能objects
// 	num := data.Len()
// 	for i := 0; i < num; i++ {

// 		tmpData := data.Index(i)

// 		if mem.MemCmp(tmpData, itemData) {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func FindElemInObjs(item interface{}, objs interface{}) int {
// 	arr := ArrToObjs(objs)
// 	for i, elem := range arr {
// 		pt.Ptln("item:", item, ";elem:", elem)
// 		pt.Ptln("item.type:")
// 		if elem == item {
// 			return i
// 		}
// 	}
// 	return -1
// }
