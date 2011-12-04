package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("package main")

	goroot := runtime.GOROOT()
	pkgDir := path.Join(goroot, "pkg", runtime.GOOS + "_" + runtime.GOARCH)

	pkgs := make(map[string][]string)
	fmt.Fprint(os.Stderr, "Scanning for packages...")
	readPackages("", pkgDir, pkgs)

	fmt.Println("import (")
	for name, _ := range pkgs {
		iName := "i_" + strings.Replace(name, "/", "_", -1)
		fmt.Printf("\t%s \"%s\"\n", iName, name)
	}
	fmt.Println(")")
}

var isPkg = regexp.MustCompile("\\.a$")

func readPackages(p string, d string, pkgs map[string][]string) {
	ents, err := ioutil.ReadDir(path.Join(d, p))
	if err != nil {
		return
	}
	for _, ent := range ents {
		n := ent.Name()
		if ent.IsDir() {
			readPackages(path.Join(p, n), d, pkgs)
			continue
		}
		if isPkg.MatchString(n) {
			pkgs[path.Join(p, n[0:len(n)-2])] = make([]string, 0)
			readPackage(p, d, n, pkgs)
		}
	}
}

func readPackage(p, d, n string, pkgs map[string][]string) {

}
