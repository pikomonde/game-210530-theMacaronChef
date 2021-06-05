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
	object     Object
	behaviours behaviours
}

type objects map[string]map[Object]objectData

// Set an Instance to Game. It groups Instances based on Object type.
// What is actually stored is the pointer (interface) of the instance.
func (objs objects) setObject(obj Object) {
	objType := reflect.TypeOf(obj).String()
	if _, ok := objs[objType]; !ok {
		objs[objType] = make(map[Object]objectData)
	}
	objs[objType][obj] = objectData{
		object:     obj,
		behaviours: make(behaviours),
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
			objData.object.Update()
		}
	}
}

// Draw all Instances.
func (objs objects) Draw(screen Screen) {
	for _, val := range objs {
		for _, objData := range val {
			objData.object.Draw(screen)
		}
	}
}

// Set an Instance to Game and initialize it.
func SetAndInitObject(obj Object) Object {
	gm.objects.setObject(obj)
	obj.Init()
	return obj
}
