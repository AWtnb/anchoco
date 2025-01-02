package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/d-tsuji/clipboard"
	"github.com/ktr0731/go-fuzzyfinder"
	"gopkg.in/yaml.v2"
)

type Entry struct {
	Name  string
	Lines []string
}

type Anchoco struct {
	entries   []Entry
	clipboard string
}

func (a *Anchoco) Init(src string) error {
	buf, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	entries := []Entry{}
	if err = yaml.Unmarshal(buf, &entries); err != nil {
		return err
	}
	a.entries = entries
	if c, err := clipboard.Get(); err == nil {
		a.clipboard = c
	} else {
		a.clipboard = ""
	}
	return nil
}

func (a Anchoco) names() (names []string) {
	for _, e := range a.entries {
		names = append(names, e.Name)
	}
	return
}

func (a Anchoco) fromName(name string) string {
	for _, e := range a.entries {
		if e.Name == name {
			return strings.Join(e.Lines, "\n")
		}
	}
	return ""
}

func (a Anchoco) applyClipboard(s string) string {
	return strings.ReplaceAll(s, "__CLIPBOARD__", a.clipboard)
}

func (a Anchoco) Select() (string, error) {
	names := a.names()
	idx, err := fuzzyfinder.Find(names, func(i int) string {
		return fmt.Sprintf("%02d %s", i, names[i])
	}, fuzzyfinder.WithPreviewWindow(func(i, _, _ int) string {
		if i == -1 {
			return ""
		}
		s := a.fromName(names[i])
		return a.applyClipboard(s)
	}))
	if err != nil {
		return "", err
	}
	s := a.fromName(names[idx])
	return a.applyClipboard(s), nil
}

func run(src string) int {
	var a Anchoco
	if err := a.Init(src); err != nil {
		fmt.Println(err)
		return 1
	}
	s, err := a.Select()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println(s)
	return 0
}

func main() {
	var (
		src string
	)
	flag.StringVar(&src, "src", "", "src yaml path")
	flag.Parse()
	os.Exit(run(src))
}
