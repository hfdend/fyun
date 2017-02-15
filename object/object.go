package object

import (
	_ "github.com/hfdend/fyun/init"
	"github.com/hfdend/fyun/g"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"strings"
)

var Bucket *oss.Bucket

type Tree struct {
	List      []*Object
	pathToObj map[string]*Object
}

func (t *Tree) GetObject(path string) (*Object, bool) {
	o, ok := t.pathToObj[path]
	return o, ok
}

type Object struct {
	oss.ObjectProperties
	Name  string
	Path  string
	Url   string
	IsDir bool
	List  []*Object
}

func GetTree() (*Tree, error) {
	lsRes, err := Bucket.ListObjects()
	if err != nil {
		return nil, err
	}
	var tree = new(Tree)
	tree.pathToObj = map[string]*Object{}
	for _, object := range lsRes.Objects {
		if strings.HasSuffix(object.Key, "/") {
			continue
		}
		pathAry := strings.Split(object.Key, "/")
		var path string
		for i, name := range pathAry {
			prePath := path
			path += "/" + name
			var o *Object
			var ok bool
			if o, ok = tree.pathToObj[path]; !ok {
				o = new(Object)
				o.ObjectProperties = object
				o.Name = name
				o.Path = path
				o.IsDir = true
				tree.pathToObj[path] = o
				if i == 0 {
					tree.List = append(tree.List, o)
				}
			}
			if i != 0 {
				tree.pathToObj[prePath].List = append(tree.pathToObj[prePath].List, o)
			}
		}
	}
	return tree, nil
}

func init() {
	client, err := oss.New(g.Endpoint, g.AccessKeyID, g.AccessKeySecret)
	if err != nil {
		log.Fatalln(err)
	}
	Bucket, err = client.Bucket(g.Bucket)
	if err != nil {
		log.Fatalln(err)
	}
}