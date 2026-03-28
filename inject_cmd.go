package main

import (
	"github.com/digital-foxy/card-parser/character"
	"github.com/digital-foxy/card-parser/png"
)

func handleInject(imageFile string, jsonFile string) error {
	rawCard, err := png.FromFile(imageFile).LastVersion().Get()
	if err != nil {
		return err
	}
	editableCard, err := rawCard.Decode()
	if err != nil {
		return err
	}

	injectedJson, err := character.FromFile(jsonFile)
	if err != nil {
		return err
	}

	editableCard.Sheet = injectedJson

	rawCard, err = editableCard.Encode()
	if err != nil {
		return err
	}

	return rawCard.ToFile(imageFile)
}
