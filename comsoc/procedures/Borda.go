package comsoc

func BordaSWF(p Profile) (Count, error){
  resultat := make(Count)
	for _, preferences := range(p){
      for i, alt := range(preferences){
        if _, ok := resultat[alt] ; !ok {
    			resultat[alt] = len(preferences) - i
    		}else{
    			resultat[alt] = resultat[alt] + len(preferences) -i
    		}
      }
  }
  return resultat, nil
}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
  resultat_pre,_ := BordaSWF(p)
	bestAlts = maxCount(resultat_pre)
  return bestAlts, nil
}
