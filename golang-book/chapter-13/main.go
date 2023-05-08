/*
func main() {
	fmt.Println(
		// true
		strings.Contains("test", "es"),

		// 2
		strings.Count("test", "t"),

		// true
		strings.HasPrefix("test", "te"),

		// true
		strings.HasSuffix("test", "st"),

		// 1
		strings.Index("test", "e"),

		// "a-b"
		strings.Join([]string{"a", "b"}, "-"),

		// == "aaaaa"
		strings.Repeat("a", 5),

		// "bbaa"
		strings.Replace("aaaa", "a", "b", 2),

		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"),

		// "test"
		strings.ToLower("TEST"),

		// "TEST"
		strings.ToUpper("test"),
	)
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
*/

/*package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
*/

/*package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
*/

/*package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value.(int))
	}
}
