package dumper

type options struct {
	imageURLs []string
}

type Option func(o *options)

func WithImage(url string) Option {
	return func(o *options) {
		o.imageURLs = append(o.imageURLs, url)
	}
}
