package yaml

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

var Author string

var Cmd = &cobra.Command{
	Use:	"yaml",
	Short:	"Read yaml files",
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

func init() {
	Cmd.Flags().StringVarP(&Author, "author", "a", "AUTHOR", "author name for attribution")
}

type Person struct {
	Name string
	Age int
}

func Execute() {
	// First Create a struct
	var p Person
	// Read the current file
	yamlFile, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	// Use the pointer to set the file
	if err := yaml.Unmarshal(yamlFile, &p); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	p.Name = Author
	// Turn the pointer into a yaml
	y, err := yaml.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	// Write the yaml
	err = ioutil.WriteFile("test.yaml", y, 0644)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(Author)
}
