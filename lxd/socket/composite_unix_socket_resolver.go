package socket

type CompositeUnixSocketResolver struct {
	resolvers []UnixSocketResolver
}

func (r CompositeUnixSocketResolver) Resolve() string {
	for _, resolver := range r.resolvers {
		path := resolver.Resolve()
		if path != "" {
			return path
		}
	}
	return ""
}

func NewCompositeUnixSocketResolver(resolvers []UnixSocketResolver) UnixSocketResolver {
	return CompositeUnixSocketResolver{
		resolvers: resolvers,
	}
}
