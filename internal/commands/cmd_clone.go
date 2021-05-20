package commands

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"libs.altipla.consulting/errors"
)

var org string
var pattern string

var cmdClone = &cobra.Command{
	Use:     "clone",
	Short:   "Clona uno o varios proyectos en local",
	Example: "sites clone",
	RunE: func(cmd *cobra.Command, args []string) error {
		var repositoryNames []string
		var dir []string

		//***LEER PROYECTOS DE LA ORG***
		client := github.NewClient(nil)

		opt := &github.RepositoryListByOrgOptions{Type: "public"}
		repos, _, err := client.Repositories.ListByOrg(context.Background(), org, opt)
		if err != nil {
			log.Fatal(errors.Stack(err))
			return err
		}

		for _, repo := range repos {
			repositoryNames = append(repositoryNames, *repo.Name)
		}

		//***LEER EL DIRECTORIO*** (y guardamos los nombres en dir)
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(errors.Stack(err))
			return err
		}

		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(errors.Stack(err))
			return err
		}

		for _, f := range files {
			dir = append(dir, f.Name())
		}

		//***COMPARAMOS DIRECTORIO Y REPOSITORIO***
		for _, x := range repositoryNames {
			if strings.HasPrefix(x, pattern) {
				for _, y := range dir {
					if x == y {
						//si el proyecto se encuentra no hacemos nada
						goto jump
					}
				}
			}

			//clonamos si no se encuentra en el directorio
			if strings.HasPrefix(x, pattern) {
				var path2 string = path + "/" + org + "/" + x
				x = "https://github.com/" + org + "/" + x
				com := exec.Command("git", "clone", x, path2)
				com.Run()
			}
		jump:
		}

		return nil
	},
}
