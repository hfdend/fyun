package init

import (
	"flag"
	"github.com/hfdend/fyun/g"
	"log"
	"errors"
)

func init() {
	endPoint := flag.String("end-point", "http://oss-cn-shanghai.aliyuncs.com", "AccessKeyID")
	addr := flag.String("addr", "127.0.0.1:8383", "ListenAndServe listens on the TCP network address addr")
	cdnUrl := flag.String("cdn-url", "", "cdn url")
	id := flag.String("id", "", "AccessKeyID")
	sec := flag.String("sec", "", "AccessKeySecret")
	bucket := flag.String("bucket", "", "Bucket")
	flag.Parse()
	g.Endpoint = *endPoint
	g.AccessKeyID = *id
	g.AccessKeySecret = *sec
	g.Bucket = *bucket
	g.CdnUrl = *cdnUrl
	g.Addr = *addr
	if g.CdnUrl == "" {
		log.Fatal(errors.New("please set cdn url"))
	}
}
