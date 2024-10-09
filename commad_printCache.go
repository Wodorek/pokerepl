package main

// debug method, to be deleted
func commandPrintCache(cfg *config) error {
	cfg.pokeapiClient.PrintCache()
	return nil
}
