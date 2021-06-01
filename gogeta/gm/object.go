package gm

import (
	"reflect"
)

type Object interface {
	Init()
	Update()
	Draw(Screen)
	D() *ObjectData
}

type ObjectData struct {
	Objects    Objects
	Behaviours Behaviours
}

type Objects map[string][]Object

// Set an Instance to Game. It groups Instances based on Object type.
func (objs Objects) Set(obj Object) {
	objType := reflect.TypeOf(obj).String()
	if _, ok := objs[objType]; !ok {
		gm.Objects[objType] = make([]Object, 0)
	}
	gm.Objects[objType] = append(gm.Objects[objType], obj)
}

// Update all Instances.
func (objs Objects) Update() {
	for _, val := range objs {
		for _, obj := range val {
			obj.Update()
		}
	}
}

// Draw all Instances.
func (objs Objects) Draw(screen Screen) {
	for _, val := range objs {
		for _, obj := range val {
			obj.Draw(screen)
		}
	}
}

// Set an Instance to Game and initialize it.
func SetObjectAndInit(obj Object) error {
	gm.Objects.Set(obj)
	obj.Init()
	return nil
}
