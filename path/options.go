package path

type Options func(path *Path)

func Data(data string) Options {
	return func(path *Path) {
		path.SetAttribute("d", data)
	}
}

func Fill(fill string) Options {
	return func(path *Path) {
		path.SetAttribute("fill", fill)
	}
}
