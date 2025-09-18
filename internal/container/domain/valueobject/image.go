package valueobject

type Image struct {
	imageName string
	tag       string
}

func (i Image) ImageName(image string, tag string) Image {
	return Image{
		imageName: image,
		tag:       tag,
	}
}
