package settings

type envVar struct {
	MemberCount int
}

var ENV envVar

const ENVPATH = "./discovery.env"
