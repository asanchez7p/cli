package commands

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"libs.altipla.consulting/errors"
)

var ctx = context.Background()
var org string
var pattern string

var cmdClone = &cobra.Command{
	Use:     "clone",
	Short:   "Clona uno o varios proyectos en local",
	Example: "sites clone",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := github.NewClient(nil)

		opt := &github.RepositoryListByOrgOptions{Type: "public"}
		repos, _, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			return errors.Trace(err)
		}

		var repositoryNames []string
		for _, repo := range repos {
			repositoryNames = append(repositoryNames, *repo.Name)
		}

		path, err := os.Getwd()
		if err != nil {
			return errors.Trace(err)
		}

		files, err := ioutil.ReadDir(path + "/" + org)
		if err != nil {
			err = os.MkdirAll(path+"/"+org, 0755)
			if err != nil {
				return errors.Trace(err)
			}
		}

		var dir []string
		for _, f := range files {
			dir = append(dir, f.Name())
		}

		var existe bool
		for _, x := range repositoryNames {
			existe = false
			if strings.HasPrefix(x, pattern) {
				for _, y := range dir {
					if x == y {
						existe = true
						fmt.Println("[" + x + "] OMITIENDO...")
						break
					}
				}
			}

			if !existe && strings.HasPrefix(x, pattern) {
				fmt.Println("[" + x + "] CLONANDO...")
				var path2 string = path + "/" + org + "/" + x
				x = "https://github.com/" + org + "/" + x
				com := exec.Command("git", "clone", x, path2)
				if err := com.Run(); err != nil {
					return errors.Trace(err)
				}
			}
		}

		return nil
	},
}

func init() {
	cmdClone.Flags().StringVarP(&org, "org", "o", "", "Organizacion de Github (requerido)")
	cmdClone.MarkFlagRequired("org")
	cmdClone.Flags().StringVarP(&pattern, "pattern", "p", "", "Prefijo de los proyectos a clonar (requerido)")
	cmdClone.MarkFlagRequired("pattern")
}
