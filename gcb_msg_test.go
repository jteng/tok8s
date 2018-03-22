package tok8s

import (
	"fmt"
	"testing"
)

func TestBuildMessage_GetImage(t *testing.T) {
	image := "image"
	sha1 := "sha1"
	bm := BuildMessage{
		Digest: fmt.Sprintf("%s:%s", image, sha1),
	}
	if i := bm.GetImage(); i != image {
		t.Errorf("expect %s got %s", image, i)
	}
}
