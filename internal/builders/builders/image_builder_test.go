package builders

import (
	"reflect"
	"testing"
)

var getImageBuilderTests = []struct {
	wantType   reflect.Type
	imageName  string
	imageInfo  ServiceInfo
	shouldFail bool
}{
	{
		imageName: "php:latest",
		imageInfo: ServiceInfo{
			ImageName: "php",
			Tag:       "latest",
		},
		wantType: reflect.TypeOf(&PhpImageBuilder{}),
	},
}

func TestGetImageBuilderByImageName(t *testing.T) {
	for _, tt := range getImageBuilderTests {
		t.Run(tt.imageName, func(t *testing.T) {
			ib, err := GetImageBuilderByImageName(tt.imageInfo)

			if tt.shouldFail && err == nil {
				t.Errorf("GetImageBuilderByImageName(%v) should have failed", tt.imageName)
				return
			}

			if err != nil {
				t.Fatalf("GetImageBuilderByImageName(%v) failed: %v", tt.imageName, err)
			}

			actualType := reflect.TypeOf(ib)

			if tt.wantType != actualType {
				t.Fatalf("expected type %v, got %v", tt.wantType, actualType)

			}
		})
	}
}

var relatedPorts = []struct {
	name     string
	input    []RelatedPorts
	expected []string
}{
	{
		name: "multiple",
		input: []RelatedPorts{
			{8080, 80},
			{9090, 92},
			{9091, 54},
		},
		expected: []string{
			"8080:80",
			"9090:92",
			"9091:54",
		},
	},
}

func TestConvertRelatedPortsToString(t *testing.T) {
	for _, tt := range relatedPorts {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertRelatedPortsToString(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ConvertRelatedPortsToString(%v): expected %v, actual %v", tt.input, tt.expected, result)
			}
		})
	}
}

var volumes = []struct {
	name     string
	input    []Volume
	expected []string
}{
	{
		name: "multiple",
		input: []Volume{
			{"app", "var/www/test_app"},
			{"Backend", "C:\\Users\\App"},
			{"app\\App", "var/www/test_app/app"},
		},
		expected: []string{
			"app:var/www/test_app",
			"Backend:C:\\Users\\App",
			"app\\App:var/www/test_app/app",
		},
	},
}

func TestConvertVolumeToString(t *testing.T) {
	for _, tt := range volumes {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertVolumeToString(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ConvertVolumeToString(%v): expected %v, actual %v", tt.input, tt.expected, result)
			}
		})
	}
}
