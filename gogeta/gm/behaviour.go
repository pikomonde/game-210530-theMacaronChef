package gm

import (
	"errors"
	"log"
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
func SetAndInitBehaviours(obj Object, bhvrArr ...Behaviour) {
	for _, bhvr := range bhvrArr {
		setBehaviour(obj, bhvr)
	}
	for _, bhvr := range bhvrArr {
		bhvr.Init()
	}
}

// Get Object's Behaviour by type.
func GetBehaviour(obj Object, bhvrType Behaviour) (Behaviour, error) {
	if bhvr := gm.objects.getObjectData(obj).behaviours.get(bhvrType); bhvr != nil {
		return bhvr, nil
	}
	return nil, errors.New(ErrBehaviourNotFound)
}

// Get Object's Behaviour by type. Must return, panic if not found.
func MustGetBehaviour(obj Object, bhvrType Behaviour) Behaviour {
	if bhvr := gm.objects.getObjectData(obj).behaviours.get(bhvrType); bhvr != nil {
		return bhvr
	}
	log.Fatalf("[GetBehaviour] Behaviour %T is not found in Object %T", bhvrType, obj)
	return nil
}

// Get relative's Behaviour by type.
func GetBehaviourRel(bhvrThis Behaviour, bhvrType Behaviour) (Behaviour, error) {
	obj := GetObject(bhvrThis)
	if bhvr := gm.objects.getObjectData(obj).behaviours.get(bhvrType); bhvr != nil {
		return bhvr, nil
	}
	return nil, errors.New(ErrBehaviourNotFound)
}

// Get relative's Behaviour by type. Must return, panic if not found.
func MustGetBehaviourRel(bhvrThis Behaviour, bhvrType Behaviour) Behaviour {
	obj := GetObject(bhvrThis)
	if bhvr := gm.objects.getObjectData(obj).behaviours.get(bhvrType); bhvr != nil {
		return bhvr
	}
	log.Fatalf("[GetBehaviourRel] Behaviour %T is not found. It is required by Behaviour %T in Object %T", bhvrType, bhvrThis, obj)
	return nil
}
