package alidns

type GetAllDomains struct {
	Base       *SignatureBase
	Action     string
	DomainName string
}

func (getalldomain *GetAllDomains) ToStringSignMap() map[string]string {
	sMap := getalldomain.Base.ToStringSignMap()
	sMap["Action"] = getalldomain.Action
	sMap["DomainName"] = getalldomain.DomainName
	return sMap
}
