package gm

type behaviourObjects map[Behaviour]Object

func (bhvro behaviourObjects) set(obj Object, bhvr Behaviour) {
	bhvro[bhvr] = obj
}

// get is getting parent object of a behaviour
func (bhvro behaviourObjects) get(bhvr Behaviour) Object {
	return bhvro[bhvr]
}

// Get Parent Object based on the Behaviour
func GetObject(bhvr Behaviour) Object {
	return gm.parentObjects.get(bhvr)
}
