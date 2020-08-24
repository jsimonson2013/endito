package document

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
)

// Node is main representation of a document
type Node struct {
	dir      string
	name     string
	children []*Node
}

var docmap = make(map[string]string)

// FromDir returns a pointer to the root of a document tree at a directory
func FromDir(d string) *Node {
	fls, err := ioutil.ReadDir(d)
	if err != nil {
		return nil
	}

	for _, f := range fls {
		if f.IsDir() {
			continue
		}

		if _, ok := docmap[nrmlz(f.Name())]; ok {
			continue
		}

		n := &Node{
			dir:  d,
			name: f.Name(),
		}

		n.Read()

		n.Print()
	}
	return nil
}

var hrf = regexp.MustCompile(`href="(?P<Link>.*)"`)

// FromLink returns a pointer to the root of a document tree at a link
func FromLink(l []byte, n *Node) *Node {
	rpl := fmt.Sprintf("${%s}", hrf.SubexpNames()[1])
	f := string(hrf.ReplaceAll(l, []byte(rpl)))

	if docmap[nrmlz(n.name)] == nrmlz(f) {
		return nil
	}

	c := &Node{
		dir:  n.dir,
		name: f,
	}

	c.Read()

	return c
}

// Read descends the document tree from the current node
func (n *Node) Read() {
	bs := n.bytes()

	for _, l := range hrf.FindAll(bs, -1) {
		c := FromLink(l, n)
		if c != nil {
			docmap[nrmlz(n.name)] = nrmlz(c.name)
			n.children = append(n.children, c)
		}
	}
}

// Print outputs the node's directory tree
// with formatting to visualize tree
func (n *Node) Print(prefix ...string) {
	if len(prefix) < 1 {
		prefix = []string{""}
	}
	fmt.Println(prefix[0] + n.name)
	prefix[0] = prefix[0] + "|_"
	for _, c := range n.children {
		c.Print(prefix[0])
	}
}

func nrmlz(tgt string) string {
	return strings.TrimPrefix(tgt, "./")
}

var files = make(map[string][]byte)

func (n *Node) bytes() []byte {
	bs, ok := files[nrmlz(path.Join(n.dir, n.name))]
	var err error
	if !ok {

		bs, err = ioutil.ReadFile(path.Join(n.dir, n.name))
		if err != nil {
			return nil
		}

		files[nrmlz(path.Join(n.dir, n.name))] = bs
	}

	return bs
}
