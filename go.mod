module github.com/lareplog/ops-sop-search

go 1.14

require (
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20200521065016-b05f73fe0dcf
	github.com/google/go-github/v32 v32.0.0
	github.com/jasonlvhit/gocron v0.0.0-20200423141508-ab84337f7963
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	k8s.io/api v0.18.5
	k8s.io/apimachinery v0.18.5
	k8s.io/client-go v0.18.2
	sigs.k8s.io/controller-runtime v0.6.0
	sigs.k8s.io/structured-merge-diff v0.0.0-20190525122527-15d366b2352e // indirect
)
