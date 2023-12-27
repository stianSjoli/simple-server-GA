package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func clone(path string, url string) *Repository {
	repo , err := git.PlainClone(path, false, &git.CloneOptions{
    	URL:      url,
    	Progress: os.Stdout,
	})
	errorCheck(err)
	return repo
} 

func Commit(filePath string, repoUrl string) {
	repo = clone("temp/", repoUrl)
	repo.WorkTree.Add(filePath)
	commit, errCommit := repo.Commit("automatic k8s template update", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "mean mug CI",
			Email: "meanMugCI@doe.org",
			When:  time.Now(),
		},
	})
	errorCheck(errCommit)
	commitObject, errCommitObject := repo.CommitObject(commit)
	errorCheck(errCommitObject)
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}