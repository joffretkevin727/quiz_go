package structure

type Question struct {
	Text   string
	Answer string
}

type PageData struct {
	Questions    []Question
	CurrentIndex int
	Score        int
	Finished     bool
	LastResult   string
	ShowResult   bool
	ButtonText   string
}
