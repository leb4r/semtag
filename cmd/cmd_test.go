package cmd

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
)

// based on https://github.com/spf13/cobra/blob/master/command_test.go

func emptyRun(*cobra.Command, []string) {}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandWithTempRepository(root *cobra.Command, args ...string) (output string, err error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir, err := ioutil.TempDir(os.Getenv("PWD"), "test-repo-")
	if err != nil {
		log.Fatal(err)
	}

	filename := filepath.Join(dir, "hello-world")
	err = ioutil.WriteFile(filename, []byte("hello world!"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir)

	r, err := git.PlainInit(dir, false)
	if err != nil {
		log.Fatal(err)
	}

	// create git config
	err = ioutil.WriteFile(filepath.Join(dir, ".git", "config"), []byte("[user]\nname = John Doe\nemail = john@doe.org"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Add("hello-world")
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Commit("initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	os.Chdir(dir)

	_, output, err = executeCommandC(root, args...)

	os.Chdir(wd)

	return output, err
}

func executeCommandWithContext(ctx context.Context, root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	err = root.ExecuteContext(ctx)

	return buf.String(), err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

func checkStringContains(t *testing.T, got, expected string) {
	if !strings.Contains(got, expected) {
		t.Errorf("Expected to contain: \n %v\nGot:\n %v\n", expected, got)
	}
}

func checkStringOmits(t *testing.T, got, expected string) {
	if strings.Contains(got, expected) {
		t.Errorf("Expected to not contain: \n %v\nGot: %v", expected, got)
	}
}

func TestEmptyRun(t *testing.T) {
	output, err := executeCommand(rootCmd)
	if output == "" {
		t.Error("Unexpected empty output")
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	checkStringContains(t, output, "Tag your repository according to Semantic Versioning")
}
