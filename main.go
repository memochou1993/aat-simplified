package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/memochou1993/mini-aat/model"
	"gopkg.in/yaml.v3"
)

const (
	RootTitle = "Top of the AAT hierarchies"
)

var (
	filename string
)

func main() {
	flag.StringVar(&filename, "f", "vocabulary.xml", "source file")
	flag.Parse()
	if err := convert(); err != nil {
		log.Fatal(err)
	}
}

func convert() error {
	input, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer closeFile(input)
	decoder := xml.NewDecoder(input)
	dst := fmt.Sprintf("%s.yaml", strings.TrimSuffix(filename, filepath.Ext(filename)))
	if err := os.Remove(dst); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	output, err := os.OpenFile(dst, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer closeFile(output)
	prefix := "---\ntitle: Art & Architecture Thesaurus\nsubjects:\n  "
	if _, err := output.Write([]byte(prefix)); err != nil {
		return err
	}
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if token == nil {
			break
		}
		switch element := token.(type) {
		case xml.StartElement:
			switch element.Name.Local {
			case "Subject":
				subject := model.Subject{}
				if err := decoder.DecodeElement(&subject, &element); err != nil {
					return err
				}
				for i, parent := range subject.ParentRelationships.PreferredParent {
					s := parent.ParentString
					if strings.Contains(s, "\n") {
						s = RootTitle
					}
					protectedTerm := findProtectedTerm(s)
					if protectedTerm != "" {
						s = strings.Replace(s, protectedTerm, "_", 1)
					}
					if i := strings.IndexAny(s, "["); i > -1 {
						s = s[:i]
					}
					if i := strings.IndexAny(s, "("); i > -1 {
						s = s[:i]
					}
					if protectedTerm != "" {
						s = strings.Replace(s, "_", protectedTerm, 1)
					}
					subject.ParentRelationships.PreferredParent[i].ParentString = strings.TrimSuffix(s, " ")
				}
				for i, note := range subject.DescriptiveNotes.DescriptiveNote {
					if i < 1 {
						s := note.NoteText
						s = strings.ReplaceAll(s, "\n", " ")
						subject.DescriptiveNotes.DescriptiveNote[i].NoteText = s
						subject.DescriptiveNotes.DescriptiveNote = []model.DescriptiveNote{subject.DescriptiveNotes.DescriptiveNote[i]}
						break
					}
				}
				subject.NewNotes = subject.DescriptiveNotes.DescriptiveNote
				var b bytes.Buffer
				encoder := yaml.NewEncoder(&b)
				encoder.SetIndent(2)
				if err := encoder.Encode(&[]model.Subject{subject}); err != nil {
					return err
				}
				s := fmt.Sprintf("%s", strings.ReplaceAll(b.String(), "\n", "\n  "))
				if _, err := output.Write([]byte(s)); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func findProtectedTerm(s string) string {
	for _, protectedTerm := range model.ProtectedTerms {
		if strings.Contains(s, protectedTerm) {
			return protectedTerm
		}
	}
	return ""
}
