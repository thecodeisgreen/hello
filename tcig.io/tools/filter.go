package tools

import "reflect"

func Filter(items interface{}, filterFunc interface{}) interface{} {
	filterFuncValue := reflect.ValueOf(filterFunc)
	itemsValue := reflect.ValueOf(items)

	itemsType := itemsValue.Type()
	itemsElemType := itemsType.Elem()

	filterFuncType := filterFuncValue.Type()

	if filterFuncType.Kind() != reflect.Func || filterFuncType.NumIn() != 1 || filterFuncType.NumOut() != 1 {
		panic("second argument must be a filter function")
	}

	if !itemsElemType.ConvertibleTo(filterFuncType.In(0)) {
		panic("Map function's argument is not compatible with type of array.")
	}

	resutSliceType := reflect.SliceOf(itemsElemType)
	resultSlice := reflect.MakeSlice(resutSliceType, 0, itemsValue.Len())

	for i := 0; i < itemsValue.Len(); i++ {
		keep := filterFuncValue.Call([]reflect.Value{itemsValue.Index(i)})[0]

		if keep.Bool() {
			resultSlice = reflect.Append(resultSlice, itemsValue.Index(i))
		}
	}

	return resultSlice.Interface()
}
