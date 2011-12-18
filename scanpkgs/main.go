package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

var importable = regexp.MustCompile(`^(func|const|var|type) @""\.`)
var fM = regexp.MustCompile(`^func @`)
var cM = regexp.MustCompile(`^const @`)
var vM = regexp.MustCompile(`^var @`)
var tM = regexp.MustCompile(`^type @`)
var first = regexp.MustCompile(`( |\().*`)

const (
	FUNC = iota
	CONST
	VAR
	TYPE
)

type item struct {
	kind int
	name string
	full string
}

func main() {
	fmt.Println(`package main

import . "reflect"`)

	// find AR(1)
	findAr()

	goroot := runtime.GOROOT()
	pkgDir := path.Join(goroot, "pkg", runtime.GOOS + "_" + runtime.GOARCH)

	pkgs := make(map[string][]item)
	fmt.Fprint(os.Stderr, "Scanning for packages...")
	readPackages("", pkgDir, pkgs)

	fmt.Println("import (")
	for name, ss := range pkgs {
		if len(ss) == 0 {
			continue
		}
		iName := escapePkgName(name)
		fmt.Printf("\t%s \"%s\"\n", iName, name)
	}
	fmt.Println(")")

	fmt.Println("var _go_imports = map[string]map[string]interface{} {")
	for name, ss := range pkgs {
		iName := escapePkgName(name)
		fmt.Printf("\"%s\": map[string]interface{} {\n", name)
		for _, i := range ss {
			if i.kind == CONST {
				if isInt(i.full) {
					fmt.Printf("%s: int64(%s.%s),\n", strconv.Quote(i.name), iName, i.name)
				} else if isUint(i.full) {
					fmt.Printf("%s: uint64(%s.%s),\n", strconv.Quote(i.name), iName, i.name)
				} else {
					fmt.Printf("%s: %s.%s, // %s\n", strconv.Quote(i.name), iName, i.name, i.full)
				}
			} else if i.kind == TYPE {
				fmt.Printf("%s: TypeOf((*%s.%s)(nil)).Elem(),\n", strconv.Quote("\x00"+i.name), iName, i.name)
			} else {
				fmt.Printf("%s: %s.%s,\n", strconv.Quote(i.name), iName, i.name)
			}
		}
		fmt.Printf("},\n")
	}
	fmt.Println("}")
	fmt.Fprintln(os.Stderr)
}

func isInt(s string) bool {
	r := regexp.MustCompile(".* = ")
	s = r.ReplaceAllString(s, "")
	_, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}
	return true
}

func isUint(s string) bool {
	r := regexp.MustCompile(".* = ")
	s = r.ReplaceAllString(s, "")
	_, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return false
	}
	return true
}

var isPkg = regexp.MustCompile("\\.a$")

func readPackages(p string, d string, pkgs map[string][]item) {
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
			pkgs[path.Join(p, n[0:len(n)-2])] = make([]item, 0)
			fmt.Fprintf(os.Stderr, " %s", path.Join(p, n[0:len(n)-2]))
			readPackage(p, d, n, pkgs)
		}
	}
}

var arpath string
func readPackage(p, d, n string, pkgs map[string][]item) {
	i := path.Join(p, n[0:len(n)-2])
	// TODO do this /without/ shelling out to AR(1)
	fname := path.Join(d, p, n)
	cmd := exec.Command(arpath, "x", fname, "__.PKGDEF")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	// read all lines from __.PKGDEF
	bs, err := ioutil.ReadFile("__.PKGDEF")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bs),"\n")
	for _, line := range lines {
		l := strings.TrimSpace(line)
		if !importable.MatchString(l) {
			continue
		}
		it := item{}
		if fM.MatchString(l) {
			it.kind = FUNC
		} else if cM.MatchString(l) {
			it.kind = CONST
		} else if vM.MatchString(l) {
			it.kind = VAR
		} else if tM.MatchString(l) {
			it.kind = TYPE
		}
		it.full = l
		l = importable.ReplaceAllString(l, "")
		it.name = first.ReplaceAllString(l, "")
		if unicode.IsUpper(getFirst(it.name)) {
			pkgs[i] = append(pkgs[i], it)
		}
	}
	os.Remove("__.PKGDEF")
}

func findAr() {
	ar, err := exec.LookPath("gopack")
	if err != nil {
		panic(err)
	}
	arpath = ar
}

func getFirst(s string) rune {
	r, _, err := strings.NewReader(s).ReadRune()
	if err != nil {
		panic(err)
	}
	return r
}

func escapePkgName(name string) string {
	return "i_" + strings.Map(func(r rune) rune {
		if strings.ContainsRune("-_./", r) {
			return '_'
		}
		return r
	}, name)
}
