package gm

type Behaviour interface {
	Init()
	Update()
	Draw(Screen)
	D() *BehaviourData
}

type BehaviourData struct {
}

type Behaviours map[string]Behaviour

// // Set an Behaviour to Object.
// func (o *Behaviours) Set(bhvr Behaviour) {
// 	bhvrType := reflect.TypeOf(bhvr).String()
// 	if _, ok := gm.Behaviours[objType]; !ok {
// 		gm.Behaviours[objType] = make([]Object, 0)
// 	}
// 	gm.Behaviours[objType] = append(gm.Behaviours[objType], obj)
// }

// // Update all Instances.
// func (o *Behaviours) Update() {
// 	for _, val := range *o {
// 		for _, obj := range val {
// 			obj.Update()
// 		}
// 	}
// }

// // Draw all Instances.
// func (o *Behaviours) Draw(screen Screen) {
// 	for _, val := range *o {
// 		for _, obj := range val {
// 			obj.Draw(screen)
// 		}
// 	}
// }

// // Set an Instance to Game and initialize it.
// func SetObjectAndInit(obj Object) error {
// 	gm.Behaviours.Set(obj)
// 	obj.Init()
// 	return nil
// }
