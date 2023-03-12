package main

import (
	"robanohashi/api"
	_ "robanohashi/docs"
	"robanohashi/internal/config"
)

// @title Roba no hashi API
// @description Query Kanji, Vocabulary, and Radicals with Mnemonics
// @version 1.0.0
// @host api.robanohashi.org
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {
	cfg := config.New()
	r := api.Create(cfg)

	r.Run(":4000")
}
