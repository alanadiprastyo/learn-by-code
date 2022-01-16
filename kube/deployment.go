package kube

type deployment struct {
	image   string
	version string
}

func New(image, version string) *deployment {

	return &deployment{
		image:   image,
		version: version,
	}
}

func (d *deployment) GetImage() string {
	return d.image
}

func (d *deployment) GetVersion() string {
	return d.version
}
