package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func (env *Env) HomePage(w http.ResponseWriter, r *http.Request) {

	type HeaderData struct {
	}
	type FooterData struct {
	}

	type HomePageData struct {
		HeaderData
		Purpose string
		FooterData
	}

	// data is passed into the html template
	data := HomePageData{}

	// get our purpose message
	data.Purpose = env.OurPurposeService.OurPurpose()

	// funcMap is passed into the html template, define helper functions here
	funcMap := template.FuncMap{
		"mod": func(a, b int) bool { return a%b == 0 },
		"add": func(a, b int) int { return a + b },
	}

	tpl := template.New("index.gohtml").Funcs(funcMap)

	// must use relative paths here, matching how files are loaded into embed FS
	_, err := tpl.ParseFS(files,
		"templates/index.gohtml",
		"templates/header.gohtml",
		"templates/footer.gohtml")
	if err != nil {
		log.Println(err.Error())
	}

	// print loaded template names, note these are name strings not file paths.
	// template names must match 'define' in the template file
	log.Println("templates:", tpl.DefinedTemplates())

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	tpl.ExecuteTemplate(w, "index.gohtml", data)

}
