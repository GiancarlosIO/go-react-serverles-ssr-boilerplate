package nrouter

// min finds the smallest value
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func longestCommonPrefix(a, b string) int {
	i := 0
	// we call it max because it is the value where we'll stop to loop
	// because at that point we have found a common prefix between a and b
	// Ex: "/sales" and "/save" will return 3 because "/s" is the common path
	max := min(len(a), len(b))
	for i < max && a[i] == b[i] {
		i++
	}
	return i
}

type nodeType uint8

const (
	static nodeType = iota // default
	root
	param
	catchAll
)

type node struct {
	path string
	handler Handler
	indices string
	ntype nodeType
	priority uint32
	children []*node
	// wildChild is used to verify if the current node is a node with a ":name" param name
	wildChild bool
}


func (n *node) addRoute(path string, handler Handler) {
	fullpath := path
	n.priority++

	// Empty tree
	if n.path == "" {
		n.insertChild(path, fullpath, handler)
		n.ntype = root
		return
	}
// we will use this label to end the loop
walk:
	for {
		// Find the longest common prefix.
		// This also implies that the common prefix contains no ':' or '*'
		// since the existing key can't contain those chars.
		// Ex: "/search" and "/save" will return 0
		i := longestCommonPrefix(path, n.path)

		// Split edge
		// for example, with n.path == "/search" and path = "/save"
		// the current `n` node will get n.path: "/s"
		// and the its two child nodes will be child1.path: "earch" and child2.path: "ave"
		// this node `node`
		if i < len(n.path) {
			child := node{
				// child1.path: "earch"
				path: n.path[i:],
				wildChild: n.wildChild,
				ntype: static,
				indices: n.indices,
				children: n.children,
				handler: n.handler,
				priority: n.priority - 1,
			}
			n.children = []*node{&child}
			// with the example above (search and save), indices will be "e" from n.path
			n.indices = string([]byte{n.path[i]})
			// n.path: "/s"
			n.path = path[:i]
			n.handler = nil
			n.wildChild = false
		}
		// make a new child node of this node
		// eg. path = "/save"
		if i < len(path) {
			// "ave"
			path := path[i:]

			if n.wildChild {
				n = n.children[0]
				n.priority++
				// TODO: Implement logic to s|upport named params "/:user"
			}

			// "a"
			idxc := path[0]
			if n.ntype == param && idxc == '/' && len(n.children) == 1 {
				n = n.children[0]
				n.priority ++
				continue walk
			}

			// check if a child with the next part exists
			for i, c := range []byte(n.indices) {
				// each === save
				//  i: 0, e == s
				// i: 1, a == a
				if c == idxc {
					// i = n.incrementChildPrio(i)
					n = n.children[i]
					continue walk
				}
			}
			// otherwise insert it
			if idxc != ':' && idxc != '*' {
				// []byte for proper unicode char conversion
				// it will be n.indices: "ea"
				n.indices += string([]byte{idxc})
				child := &node{}
				n.children = append(n.children, child)
				// n.incrementChildPrio(len(n.indices) - 1)
				n = child
			}
			n.insertChild(path, fullpath, handler)
			return
		}
		// otherwise add handler to the current node
		if n.handler != nil {
			panic("a handler is already registered for path '" + fullpath + "'")
		}
		n.handler = handler
		return
	}
}

func (n *node) insertChild(path string, _ string, handler Handler) {
	// TODO: Implement logic to support wildcards ('*' and ':') in urls
	n.path = path
	n.handler = handler
}

