package data

import (
	"botik/models"
	"encoding/json"
	"os"
)

func Training() (models.TrainingExercises, models.TrainingExercises, error) {
	var att models.TrainingExercises
	var def models.TrainingExercises
	data, err := os.ReadFile("../attack.json")
	if err != nil {
		return att, def, err
	}

	err = json.Unmarshal(data, &att)
	if err != nil {
		return att, def, err
	}

	data2, err := os.ReadFile("../defense.json")
	if err != nil {
		return att, def, err
	}

	err = json.Unmarshal(data2, &def)
	if err != nil {
		return att, def, err
	}
	return att, def, nil
}

func Facts() (models.FactsRand, error) {
	var facts models.FactsRand
	data, err := os.ReadFile("../facts.json")
	if err != nil {
		return facts, err
	}

	err = json.Unmarshal(data, &facts)
	if err != nil {
		return facts, err
	}
	return facts, nil
}

func Tactics() (models.TacticsRand, error) {
	var tactics models.TacticsRand
	data, err := os.ReadFile("../tactics.json")
	if err != nil {
		return tactics, err
	}

	err = json.Unmarshal(data, &tactics)
	if err != nil {
		return tactics, err
	}
	return tactics, nil
}
