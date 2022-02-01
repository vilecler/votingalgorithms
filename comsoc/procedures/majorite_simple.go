package comsoc

func MajoritySWF(p Profile) (count Count, err error){
	// p est un tableau de tableau d'Alternative, les premiers indexes représentent les agents, le reste représente les preférences
	// on doit encore gérer les erreurs
	resultat := make(Count)
	for _ , preferences := range(p){
		//vote a un tour donc uniquement preferences[0] est considéré
		if _, ok := resultat[preferences[0]] ; !ok {
			resultat[preferences[0]] = 1
		}else{
			resultat[preferences[0]] = resultat[preferences[0]] + 1
		}
	}
	return resultat, nil
}


func MajoritySCF(p Profile) (bestAlts []Alternative, err error){
	resultat, err := MajoritySWF(p)
	bestAlts = maxCount(resultat)
	return bestAlts, nil
 	// gérer le cas des alternatives mutliples
}
