package comsoc

import (
  "testing";
  "reflect";
)

func TestApprovalSWF(t *testing.T){
  //func BordaSWF(p Profile) (Count, error){
  var prof [][]Alternative
  voter1 := []Alternative{1,2,3}
  voter2 := []Alternative{2,1}
  voter3 := []Alternative{1}
  prof = append(prof, voter1, voter2, voter3)
  got, _ := ApprovalSWF(prof)
  want := map[Alternative]int{
    1: 3,
    2: 2,
    3: 1,
}

  for alt, val := range(got){
    if val != want[alt]{
      t.Errorf("On a %d alors que l'on veut %d etant donne, %v", got, want, prof)
    }
  }
}


func TestApprovalSCF(t *testing.T){
  var prof [][]Alternative
  voter1 := []Alternative{1,2,3}
  voter2 := []Alternative{2,1}
  voter3 := []Alternative{1}
  prof = append(prof, voter1, voter2, voter3)
  got, _ := ApprovalSCF(prof)
  want := []Alternative{1}
  if !( reflect.DeepEqual(got, want)) {
    t.Errorf("On a %d  alors que l'on veut %d etant donne, %v", got, want,   prof)
  }
}
