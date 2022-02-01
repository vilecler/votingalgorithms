package comsoc

import "fmt"
import "errors"

/*
 * Fonction rank
 *
 * Retourne le classement de l'alternative dans l'esemble des préférences (de 0 à n)
 * Retourne -1 si l'alternative n'est pas
 */
func rank(alt Alternative, prefs []Alternative) int {
	for i := 0; i < len(prefs); i++ {
		if prefs[i] == alt {
			return i
		}
	}
	return -1
}

/*
 * Fonction maxCount
 *
 * Retourne un tableau d'alternatives classés en fonction de leur score dans counts
 */
func maxCount(counts Count) (bestAlts []Alternative) {
	bestAlts = []Alternative{}
	for alt := range counts { //on itère la map
		if len(bestAlts) == 0 { //premier cas
			bestAlts = append(bestAlts, alt)
			continue
		}

		for i := 0; i < len(bestAlts); i++ {
			if len(bestAlts) > 2 {
				if counts[bestAlts[i]] <= counts[alt] && counts[alt] <= counts[bestAlts[i+1]] { //on doit insérer alt à la (i+1) position de bestAlts
					prefix := bestAlts[0 : i+1]
					suffix := bestAlts[i+1 : len(bestAlts)]
					tempSuffix := make([]Alternative, len(bestAlts)-i-1)
					copy(tempSuffix, suffix)
					result := append(prefix, []Alternative{alt}...)
					result = append(result, tempSuffix...)
					bestAlts = result
					break
				}
			}
			if counts[alt] < counts[bestAlts[0]] { //prepend
				bestAlts = append([]Alternative{alt}, bestAlts...)
			} else { //append
				bestAlts = append(bestAlts, alt)
			}
		}
	}
	return bestAlts
}

/*
 * Fonction checkPrefs
 *
 *
 */
func checkProfile(prefs Profile) error {
	if len(prefs) <= 0 {
		return errors.New("Profile is empty")
	}
	n := len(prefs[0])
	for i := 0; i < len(prefs); i++ {
		if len(prefs[i]) != n {
			return errors.New("Preferences sizes are not equals")
		}

		for j := 0; j < len(prefs[i]); j++ { //on vérifie si des alternatives sont présentes deux fois
			c := 0
			for k := 0; k < len(prefs[i]); k++ {
				if i != k && prefs[i][j] == prefs[i][k] {
					c++
					if c >= 2 {
						return errors.New("Alternative must be chosen once at maximum by agent")
					}
				}
			}
		}
	}
	return nil
}

func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	if len(prefs) <= 0 {
		return errors.New("Profile is empty")
	}

	if len(prefs[0]) != len(alts) {
		return errors.New("Each alternative must be present once")
	}

	return checkProfile(prefs)
}

func Test() {
	fmt.Println("Test")

	alternatives := []Alternative{2, 6, 5, 71, 156, 42, 17, 26}

	counts := make(Count)

	counts[2] = 0
	counts[156] = 4
	counts[5] = 1
	counts[17] = 15
	counts[26] = 4

	fmt.Println(rank(-1, alternatives))
	fmt.Println(maxCount(counts))
}
