package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type SecretData map[string]string

type Secret struct {
	ApiVersion string      `yaml:"apiVersion"`
	Data       SecretData  `yaml:"data"`
	Kind       string      `yaml:"kind"`
	Metadata   interface{} `yaml:"metadata"`
	Type       string      `yaml:"type"`
}

var (
	encode = flag.Bool("encode", false, "Encode values, rather than decoding")
)

func main() {
	flag.Usage = func() {
		_, exe := filepath.Split(os.Args[0])
		fmt.Fprint(os.Stderr, "Easily encode/decode Kubernetes \"Secrets\" yaml.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\n  %s [options]\n\nOptions:\n\n", exe)
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()

	var reader io.Reader
	var err error

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "reading from stdin...")
		reader = os.Stdin
	} else {
		reader, err = os.Open(args[0])
		if err != nil {
			panic(err)
		}
	}

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	secret := Secret{}
	err = yaml.Unmarshal(bytes, &secret)
	if err != nil {
		panic(err)
	}

	for k, v := range secret.Data {
		if *encode {
			secret.Data[k] = base64.StdEncoding.EncodeToString([]byte(v))
		} else {
			tv, err := base64.StdEncoding.DecodeString(v)
			if err != nil {
				panic(err)
			}
			secret.Data[k] = string(tv)
		}
	}

	out, err := yaml.Marshal(&secret)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, string(out))
}
