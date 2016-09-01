package hschema

func (ll *LinkList) SetParent(s *HyperSchema) {
	for _, l := range *ll {
		l.SetParent(s)
	}
}

func (l *Link) SetParent(s *HyperSchema) {
	l.parent = s
}

func (l Link) Parent() *HyperSchema {
	return l.parent
}

func (l *Link) Path() string {
	if l.parent != nil {
		return l.parent.PathStart + l.Href
	}

	return l.Href
}
