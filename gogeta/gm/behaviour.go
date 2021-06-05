package gm

import (
	"reflect"
)

type Behaviour interface {
	Init()
	Update()
	Draw(Screen)
}

type behaviourData struct {
}

type behaviours map[string]Behaviour

// Set a Behaviour to Object.
func (bhvrs behaviours) set(bhvr Behaviour) {
	bhvrType := reflect.TypeOf(bhvr).String()
	bhvrs[bhvrType] = bhvr
}

// Get a Behaviour from Object.
func (bhvrs behaviours) get(bhvr Behaviour) Behaviour {
	bhvrType := reflect.TypeOf(bhvr).String()
	return bhvrs[bhvrType]
}

// Update all Instances.
func (bhvrs behaviours) Update() {
	for _, bhvr := range bhvrs {
		bhvr.Update()
	}
}

// Draw all Instances.
func (bhvrs behaviours) Draw(screen Screen) {
	for _, bhvr := range bhvrs {
		bhvr.Draw(screen)
	}
}

func setBehaviour(obj Object, bhvr Behaviour) {
	// set behaviour to object's behaivour
	gm.objects.getObjectData(obj).behaviours.set(bhvr)

	// set behaviour to top level game
	gm.behaviours.set(obj, bhvr)
}

// Set a Behaviour to Object and initialize it.
func SetBehaviourAndInit(obj Object, bhvr Behaviour) error {
	setBehaviour(obj, bhvr)
	bhvr.Init()
	return nil
}

// Get Object's Behaviour by type
func GetBehaviour(obj Object, bhvrType Behaviour) Behaviour {
	return gm.objects.getObjectData(obj).behaviours.get(bhvrType)
}

// Get relative's Behaviour by type
func GetBehaviourRel(bhvrThis Behaviour, bhvrType Behaviour) Behaviour {
	obj := GetObject(bhvrThis)
	return gm.objects.getObjectData(obj).behaviours.get(bhvrType)
}
