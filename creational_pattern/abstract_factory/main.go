package main

func main() {
	app, _ := new(Application).SetOsFactory(WINDOWS)
	app.CreateUI()
}
