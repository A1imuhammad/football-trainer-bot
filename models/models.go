package models

type Token struct {
	TGtoken string `env:"BOT_TOKEN"`
}

var UsersID []int64

type TacticsRand struct {
	Tactics struct {
		Defense     []string `json:"defense"`
		Attack      []string `json:"attack"`
		Positioning []string `json:"positioning"`
	} `json:"tactics"`
}

type FactsRand []struct {
	Fact string `json:"fact"`
}
type Exercise struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Focus       string `json:"focus"`
}

type TrainingExercises struct {
	General             []Exercise `json:"General"`
	Forward             []Exercise `json:"Forward"`
	Winger              []Exercise `json:"Winger"`
	DefensiveMidfielder []Exercise `json:"DefensiveMidfielder"`
	CentralMidfielder   []Exercise `json:"CentralMidfielder"`
	AttackingMidfielder []Exercise `json:"AttackingMidfielder"`
	Defender            []Exercise `json:"Defender"`
	Laterale            []Exercise `json:"Laterale"`
	Goalkeeper          []Exercise `json:"Goalkeeper"`
}


