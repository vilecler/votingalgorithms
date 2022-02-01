package comsoc

func ApprovalSWF(p Profile) (count Count, err error){
  // les profils sont incomplets, les candidats non soutenus sont ceux qui ne sont pas dans les prefs
  resultat := make(Count)
	for _, preferences := range(p){
    for _, alt := range(preferences){
      if _, ok := resultat[alt] ; !ok {
        resultat[alt] = 1
      }else{
        resultat[alt] = resultat[alt] + 1
      }
    }
  }
  return resultat, nil
}

func ApprovalSCF(p Profile) (bestAlts []Alternative, err error){
  resultat_pre, _ := ApprovalSWF(p)
  bestAlts = maxCount(resultat_pre)
  return bestAlts, nil
}
