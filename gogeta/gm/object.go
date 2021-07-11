package gm

import (
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

// objects stored objectData with mapping by objectType and objectName.
// For example. there are 2 "enemy" objects and 1 "player" object, first map
// indexed by "enemy" and "player", meanwhile the second mapping contain the
// address (interface) of "enemy 1", "enemy 2", and "player 1".
type objects map[string]map[Object]objectData

// TODO: #redesignmapobject
// Example 1: all object pointer							map[obj][]objectData
// Example 2: finding by exact object pointer				map[obj::"*objPointer"][0]objectData
// Example 3: finding by object type						map[objT::"objType"][]objectData
// Example 4: finding by exact behaviour pointer (for
//            parent object case)							map[bhv::"*bhvPointer"][0]objectData
// Example 5: finding by z-index							map[zidx][]objectData
// type objects map[string][]objectData

// Set an Instance to Game. It groups Instances based on Object type. What is
// actually stored is the pointer (interface) of the instance.
func (objs objects) setObject(obj Object) {
	objType := reflect.TypeOf(obj).String()
	if _, ok := objs[objType]; !ok {
		objs[objType] = make(map[Object]objectData)
	}
	objs[objType][obj] = objectData{
		object: obj,
		// behaviours: make(behaviours),
	}
}

func (objs objects) getObjectData(obj Object) objectData {
	objType := reflect.TypeOf(obj).String()
	return objs[objType][obj]
}

// Update all Instances.
func (objs objects) Update() {
	for _, val := range objs {
		for _, objData := range val {
			updateBehaviours(objData.object)
			objData.object.Update()
		}
	}
}

// Draw all Instances.
func (objs objects) Draw(screen Screen) {
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
