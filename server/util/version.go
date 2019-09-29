package util

var (
	version string
	hash    string
	branch  string
)

func GetVersion() string {
	if version == "" {
		return "bleeding edge"
	} else {
		return "v" + version
	}
}

func GetHash() string {
	if hash == "" {
		return ""
	} else {
		return hash
	}
}

func GetBranch() string {
	if branch == "" {
		return ""
	} else {
		return branch
	}
}

func GetFullVersion() string {
	h := GetHash()
	b := GetBranch()
	if h == "" || b == "" {
		return GetVersion()
	} else {
		return GetVersion() + "-" + h + "-" + b
	}
}
