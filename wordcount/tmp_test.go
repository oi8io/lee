package tmp_test

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestScanDir(t *testing.T) {
	var m = make(map[string]int)
	root := "/Users/dymo/Develop/go/go/src/"
	//root = "/Users/dymo/Develop/go/go/src/sort"
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.Contains(info.Name(), ".go") {
			err := wordCount(path, m)
			if err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	open, err := os.OpenFile("/Users/dymo/Develop/go/crm/library/tmp/cnt.txt", 777, fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	arr := make([]string, 0)
	for s, i := range m {
		s0 := fmt.Sprintf("%d %s", i, s)
		arr = append(arr, s0)
	}
	open.WriteString(strings.Join(arr, "\n"))

}

func wordCount(path string, m map[string]int) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(file)
	compile, _ := regexp.Compile("//(.*)")
	compile1, _ := regexp.Compile(`([a-z_-]+)`)

	split := strings.Split(content, "\n")
	for i := 0; i < len(split); i++ {
		matchString := compile.FindAllStringSubmatch(split[i], -1)
		if len(matchString) == 0 {
			continue
		}
		str := matchString[0][1]
		fields := strings.Fields(str)
		for j := 0; j < len(fields); j++ {
			matchs := compile1.FindAllStringSubmatch(fields[j], -1)
			if len(matchs) == 0 || len(matchs[0][1]) == 1 {
				continue
			}
			//trimFunc := strings.TrimFunc(fields[j], func(r rune) bool {
			//	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' || r == '-' {
			//		return false
			//	} else {
			//		return true
			//	}
			//})
			w := strings.ToLower(matchs[0][1])
			m[w]++
		}
	}
	return err
}
