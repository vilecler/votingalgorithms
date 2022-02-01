func maxCount(counts map[Alternative]int) (bestAlts []Alternative){
		var BestA []Alternative
		max := 0
		for len(counts) != 0{
			for alternat, score := range(counts){
				if score > max {
					max = score
					BestA = []Alternative{alternat}
					delete(counts, alternat)
				}else if score == max{
					BestA = append(BestA, alternat)
					delete(counts, alternat)
				}else{
					delete(counts, alternat)
				}
			}
		}
    return BestA
}
