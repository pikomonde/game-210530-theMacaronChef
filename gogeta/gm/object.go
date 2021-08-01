package gm

import (
	"fmt"
	"reflect"
)

type Object interface {
	Init()
	Update()
	Draw(Screen)
}

// Set an Instance to Game.
func (objs objects) setObject(obj Object) {
	objd := objectData{object: obj}
	objType := reflect.TypeOf(obj).String()

	// case: keyByObj
	key := keyByObj
	if _, ok := objs[key]; !ok {
		objs[key] = make(map[Object]objectData)
	}
	objs[key][obj] = objd

	// case: keyByObjType
	key = fmt.Sprintf("%s%s", keyByObjType, objType)
	if _, ok := objs[key]; !ok {
		objs[key] = make(map[Object]objectData)
	}
	objs[key][obj] = objd
}

// Delete an Instance from Game.
func (objs objects) delObject(obj Object) {
	objType := reflect.TypeOf(obj).String()

	// case: keyByObj
	key := keyByObj
	delete(objs[key], obj)

	// case: keyByObjType
	key = fmt.Sprintf("%s%s", keyByObjType, objType)
	delete(objs[key], obj)
	if len(objs[key]) == 0 {
		delete(objs, key)
	}
}

// Set an Instance to Game and initialize it.
func SetAndInitObject(obj Object) Object {
	gm.objects.setObject(obj)
	initBehaviours(obj)
	obj.Init()
	return obj
}
