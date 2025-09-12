package application

func Run(app BaseWidget) {
	context.AppendChild(app)
	notifyPageLoaded()
	select {}
}
