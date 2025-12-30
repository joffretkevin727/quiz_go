package controller

import (
	"bytes"
	"net/http"
	"proj3/structure"
	"strconv"
	"strings"
	"text/template"
)

var QuizQuestions = []structure.Question{
	{Text: "Quel mot-clé définit une fonction en Go ?", Answer: "func"},
	{Text: "Quel package gère l'affichage formaté ?", Answer: "fmt"},
	{Text: "Quelle est la valeur par défaut d'un type int ?", Answer: "0"},
	{Text: "Comment déclare-t-on une variable courte ?", Answer: ":="},
	{Text: "Quel mot-clé est utilisé pour gérer la concurrence ?", Answer: "go"},
	{Text: "Quel type de données gère les listes dynamiques ?", Answer: "slice"},
	{Text: "Comment s'appelle l'outil de gestion de dépendances ?", Answer: "go mod"},
	{Text: "Quel mot-clé retarde l'exécution d'une fonction ?", Answer: "defer"},
	{Text: "Quelle boucle existe exclusivement en Go ?", Answer: "for"},
	{Text: "Quel mot-clé est utilisé pour les structures ?", Answer: "struct"},
	{Text: "Quelle commande initialise un dépôt ?", Answer: "git init"},
	{Text: "Quelle commande prépare les fichiers (staging) ?", Answer: "git add"},
	{Text: "Quelle commande enregistre les modifications ?", Answer: "git commit"},
	{Text: "Quelle commande affiche l'état des fichiers ?", Answer: "git status"},
	{Text: "Quelle commande envoie vers un dépôt distant ?", Answer: "git push"},
	{Text: "Quelle commande crée une nouvelle branche ?", Answer: "git branch"},
	{Text: "Quelle commande fusionne deux branches ?", Answer: "git merge"},
	{Text: "Quelle commande récupère les changements distants ?", Answer: "git pull"},
	{Text: "Quelle commande annule les changements locaux ?", Answer: "git restore"},
	{Text: "Quel fichier liste les documents à ignorer ?", Answer: ".gitignore"},
}

func RenderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	// Analyse et rend le template spécifié
	tmpl := template.Must(template.ParseFiles("template/" + filename))
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := structure.PageData{
		Questions:    QuizQuestions,
		CurrentIndex: 0,
		Score:        0,
		Finished:     false,
		ButtonText:   "Valider",
	}
	RenderTemplate(w, "home.html", data)
}
func Quiz(w http.ResponseWriter, r *http.Request) {
	idx, _ := strconv.Atoi(r.FormValue("index"))
	score, _ := strconv.Atoi(r.FormValue("score"))
	showResult := r.FormValue("show_result") == "true"

	if showResult {
		idx++
		data := structure.PageData{
			Questions:    QuizQuestions,
			CurrentIndex: idx,
			Score:        score,
			Finished:     idx >= len(QuizQuestions),
			ShowResult:   false,
			ButtonText:   "Valider",
		}
		RenderTemplate(w, "home.html", data)
		return
	}

	userAns := strings.TrimSpace(r.FormValue("answer"))
	result := "Incorrect..."
	if strings.EqualFold(userAns, QuizQuestions[idx].Answer) {
		score++
		result = "Correct !"
	}
	btnText := "Prochaine question"
	if idx+1 == len(QuizQuestions) {
		btnText = "Voir les résultats"
	}
	data := structure.PageData{
		Questions:    QuizQuestions,
		CurrentIndex: idx,
		Score:        score,
		LastResult:   result,
		ShowResult:   true,
		ButtonText:   btnText,
	}
	RenderTemplate(w, "home.html", data)
}
