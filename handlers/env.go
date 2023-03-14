package handlers

import (
	"embed"
	"io/fs"
	"log"
	"text/template"

	"github.com/thisdougb/cleango/pkg/usecase/enablething"
	"github.com/thisdougb/cleango/pkg/usecase/ourpurpose"
)

/*
   The Env struct allows us to pass the datastore pointer around,
   it also allows us to inject mocks in usecase packages.
*/

type Env struct {
	Logger             *log.Logger
	EnableThingService *enablething.Service
	OurPurposeService  *ourpurpose.Service
}

const (
	templatesDir = "templates"
	extension    = "/*.gohtml"
)

var (
	//go:embed templates
	files     embed.FS
	templates map[string]*template.Template
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	loadTemplates()
}

func loadTemplates() {

	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	tmplFiles, err := fs.ReadDir(files, "templates")
	if err != nil {
		log.Println("files error", err.Error())
		return
	}

	funcMap := template.FuncMap{
		"mod": func(a, b int) bool { return a%b == 0 },
		"add": func(a, b int) int { return a + b },
	}

	// loop through files in embed FS and load each as a template[]
	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		// parse the template using the relative dir/pathname
		//log.Println("parsing template:", tmpl.Name())

		// parse file to template, including funcMap
		pt, err := template.New(tmpl.Name()).Funcs(funcMap).ParseFS(files, templatesDir+"/"+tmpl.Name())
		if err != nil {
			log.Println("parse error", err.Error())
			return
		}

		templates[tmpl.Name()] = pt
	}

	// debug to show all files in the embed FS
	all, _ := getAllFilenames(&files)
	for _, f := range all {
		log.Println("embed FS file:", f)
	}
}

func getAllFilenames(efs *embed.FS) (files []string, err error) {
	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			log.Println("dir", d.Name())
			return nil
		}
		log.Println("file:", path)

		files = append(files, path)

		return nil
	}); err != nil {
		log.Println("err:", err.Error())
		return nil, err
	}

	return files, nil
}
