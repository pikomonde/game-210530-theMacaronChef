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

type objectData struct {
	object Object
	// behaviours behaviours
}

// Example 1: all object pointer & finding by exact
//            object pointer								map[obj]				map["*objPointer"]objectData
//
// Example 2: finding by object type						map[objT::"objType"]	map["*objPointer"]objectData
// objects stored objectData with mapping by objectType and objectName.
// For example. there are 2 "enemy" objects and 1 "player" object, first map
// indexed by "enemy" and "player", meanwhile the second mapping contain the
// address (interface) of "enemy 1", "enemy 2", and "player 1".
//
// Example 3: finding by exact behaviour pointer (for
//            parent object case)							map[bhvr::"*bhvPointer"]map["*objPointer"]objectData
//
// Example 4: finding by z-index							map[zidx]				map["*objPointer"]objectData
type objects map[string]map[Object]objectData

const (
	keyByObj     = "obj"
	keyByObjType = "objT::"
	keyByBhvr    = "bhvr::"
	keyByZIdx    = "zidx"
)

// Set an Instance to Game.
func (objs objects) setObject(obj Object) {
	objd := objectData{
		object: obj,
		// behaviours: make(behaviours),
	}
	objType := reflect.TypeOf(obj).String()

	// case: keyByObj
	key := keyByObj
	if _, ok := objs[key]; !ok {
		objs[key] = make(map[Object]objectData)
	}
	objs[key][obj] = objd

	// case: keyByObjType
	key = fmt.Sprintf("%s%s\n", keyByObjType, objType)
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
	key = fmt.Sprintf("%s%s\n", keyByObjType, objType)
	delete(objs[key], obj)
}

func (objs objects) getObjectData(obj Object) objectData {
	objType := reflect.TypeOf(obj).String()
	return objs[objType][obj]
}

// Update all Instances.
func (objs objects) Update() {
	// TODO: behaviours should be updated outside from this loop
	for _, val := range objs {
		for _, objData := range val {
			updateBehaviours(objData.object)
			objData.object.Update()
		}
	}
}

// Draw all Instances.
func (objs objects) Draw(screen Screen) {
	// TODO: consider z-index when draw objects
	// TODO: behaviours should be draw outside from this loop
	for _, val := range objs {
		for _, objData := range val {
			drawBehaviours(objData.object, screen)
			objData.object.Draw(screen)
		}
	}
}

// Set an Instance to Game and initialize it.
func SetAndInitObject(obj Object) Object {
	gm.objects.setObject(obj)
	initBehaviours(obj)
	obj.Init()
	return obj
}
