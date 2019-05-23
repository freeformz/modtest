package modtest

const (
	version = "v1.1.1"
)

func It() string {
	return "modtest " + version
}

func ItNow() string {
	return "now " + It()
}
