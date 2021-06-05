package gm

type behaviourObjects map[Behaviour]Object

func (bhvro behaviourObjects) set(obj Object, bhvr Behaviour) {
	bhvro[bhvr] = obj
}

func (bhvro behaviourObjects) get(bhvr Behaviour) Object {
	return bhvro[bhvr]
}

// Get Object based on the Behaviour
func GetObject(bhvr Behaviour) Object {
	return gm.behaviours.get(bhvr)
}
