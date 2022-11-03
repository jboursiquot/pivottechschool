package main

import (
	"math"
	"reflect"
	"sort"
	"testing"
)

var c = Client{}

func TestGetFirst25(t *testing.T) {
	expected := []string{"3-D Man", "A-Bomb (HAS)", "A.I.M.", "Aaron Stack", "Abomination (Emil Blonsky)", "Abomination (Ultimate)", "Absorbing Man", "Abyss", "Abyss (Age of Apocalypse)", "Adam Destine", "Adam Warlock", "Aegis (Trey Rollins)", "Aero (Aero)", "Agatha Harkness", "Agent Brand", "Agent X (Nijo)", "Agent Zero", "Agents of Atlas", "Aginar", "Air-Walker (Gabriel Lan)", "Ajak", "Ajaxis", "Akemi", "Alain", "Albert Cleary"}
	res := c.GetFirst25()
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("incorrect result; want %v got %v", expected, res)
	}
}

func TestGetNameByID(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		id := 1009742
		expected := "Zzzax"
		res, err := c.GetNameByID(id)
		assertNoErr(t, err)
		assertEqual(t, expected, res)
	})

	t.Run("not found", func(t *testing.T) {
		res, err := c.GetNameByID(99999999999999)
		assertEqual(t, ErrNotFound, err)
		assertEqual(t, "", res)
	})
}

func TestGetCharacterDetail(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		res, err := c.GetCharacterDetail("Captain America")
		assertNoErr(t, err)
		assertEqual(t, CharacterDetail{
			ID:           1009220,
			Description:  "Vowing to serve his country any way he could, young Steve Rogers took the super soldier serum to become America's one-man army. Fighting for the red, white and blue for over 60 years, Captain America is the living, breathing symbol of freedom and liberty.",
			ThumbnailURL: "http://i.annihil.us/u/prod/marvel/i/mg/3/50/537ba56d31087.jpg",
			ComicCount:   2425,
		}, res)
	})
	t.Run("Not found", func(t *testing.T) {
		res, err := c.GetCharacterDetail("kwerwerkljawerlkj")
		assertEqual(t, ErrNotFound, err)
		assertEqual(t, CharacterDetail{}, res)
	})
}

func TestGetCharactersInComic(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		id := 1994
		expected := []string{
			"Apocalypse",
			"Blink",
			"Colossus",
			"Gambit",
			"Holocaust (Age of Apocalypse)",
			"Magneto",
			"Mister Sinister",
			"Rogue",
			"Sabretooth (Age of Apocalypse)",
			"Shadowcat (Age of Apocalypse)",
			"Silver Samurai (Age of Apocalypse)",
			"Storm (Age of Apocalypse)",
			"Sunfire",
			"Wolverine",
		}
		res, err := c.GetCharactersInComic(id)
		sort.Strings(expected)
		sort.Strings(res)
		assertNoErr(t, err)
		assertEqual(t, expected, res)
	})
	t.Run("Not found", func(t *testing.T) {
		_, err := c.GetCharactersInComic(math.MaxInt)
		assertEqual(t, ErrNotFound, err)
	})
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("want %v, got %v", expected, actual)
	}
}

func assertNoErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
}
