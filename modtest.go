package modtest

const (
	version = "v2.0.0"
)

func It() (string, string) {
	return "modtest", version
}

func ItNow() (string, string) {
	m, v := It()
	return "now " + m, v
}
