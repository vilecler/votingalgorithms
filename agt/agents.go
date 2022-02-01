package agt

import prenoms "gitlab.utc.fr/vilecler/tp6_ia04/agt/prenoms"

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

type Alternative int
type AgentID string

type AgentI interface {
	Equal(ag AgentI) bool
	DeepEqual(ag AgentI) bool
	Clone() AgentI
	String() string
	Prefers(a Alternative, b Alternative)
}

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []Alternative
}

func NewAgent(id AgentID, name string, prefs []Alternative) Agent {
	return Agent{
		id,
		name,
		prefs,
	}
}

func (ag1 Agent) Equal(ag2 Agent) bool {
	if ag1.ID != ag2.ID {
		return false
	}

	return true
}

func (ag1 Agent) DeepEqual(ag2 Agent) bool {
	if ag1.ID != ag2.ID {
		return false
	}

	if ag1.Name != ag2.Name {
		return false
	}

	if len(ag1.Prefs) != len(ag2.Prefs) {
		return false
	}

	for i := range ag1.Prefs {
		if ag1.Prefs[i] != ag1.Prefs[i] {
			return false
		}
	}

	return true
}

func (a Agent) Clone() Agent {
	return NewAgent(a.ID, a.Name, a.Prefs)
}

func (a Agent) String() string {
	return fmt.Sprintf("%s %s %v", a.ID, a.Name, a.Prefs)
}

func (a Agent) rank(b Alternative) (int, error) {
	for i, v := range a.Prefs {
		if v == b {
			return i, nil
		}
	}
	return -1, errors.New("Alternative not found")
}

// renvoie vrai si ag préfère a à b
func (ag Agent) Prefers(a, b Alternative) bool {
	r1, err1 := ag.rank(a)
	if err1 != nil {
		return false
	}

	r2, err2 := ag.rank(b)
	if err2 != nil {
		return false
	}

	return r1 < r2
}

func RandomPrefs(ids []Alternative) (res []Alternative) {
	res = make([]Alternative, len(ids))
	copy(res, ids)
	rand.Shuffle(len(res), func(i, j int) { res[i], res[j] = res[j], res[i] })
	return
}

/*
 *  Fonction GenerateAgents
 *
 *	Permet de générer les deux groupes d'agents à partir d'une liste de prénom et d'une taille n.
 *
 */
func GenerateAgents(p []prenoms.Prenom, n int) (agents []Agent) {
	randomPrenoms := prenoms.ShufflePrenoms(p, n)

	agents = make([]Agent, 0, n)

	agentPrefix := "a"

	prefsAlternatives := make([]Alternative, n)

	for i := 0; i < n; i++ { //generating alternatives
		prefsAlternatives[i] = Alternative(i)
	}

	for i := 0; i < n; i++ { //generating agent
		prefsAgent := RandomPrefs(prefsAlternatives)
		agent := NewAgent(agentPrefix+AgentID(strconv.Itoa(int(prefsAgent[i]))), randomPrenoms[i].Prenom, prefsAlternatives)
		agents = append(agents, agent)
	}

	return agents
}
