package modtest

const (
	version = "v1.0.0"
)

func It() string {
	return "modtest " + version
}

func ItNow() string {
	return "now " + It()
}
