package GetData

import (
	"testing"
)

func TestSum(t *testing.T) {
	const PATH = "../images"
	const port = "5000"
	Start(port, PATH)
}
