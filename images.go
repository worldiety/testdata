package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(1234)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var resolution = []int{2592, 1944, 4032, 3024, 4912, 3264}
var cameras = []string{"iPhone 7", "Sony A57", "Sony RX100M5"}
var text = []string{"Horse", "Child", "Monkey", "Dog", "Cat", "Tree", "Apple", "Sun", "Colorful", "Bike", "Car"}
var face = []string{"Mr X.", "Mr Y.", "Gopher", "Ant-Man", "Iron Man", "Tony Stark", "Frodo", "Beutlin", "Gandalf"}

func CreateString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Camera() string {
	return cameras[rand.Intn(len(cameras))]
}

func Resolution() (int, int) {
	return resolution[rand.Intn(len(resolution))], resolution[rand.Intn(len(resolution))]
}

func Time(month int, hour int) int64 {
	return time.Date(2019, time.Month(month%12), 19, hour%24, 42, 13, month, time.Local).Unix()
}

func Caption(i int) string {
	return fmt.Sprintf("I took the %d nth image of %s", i, text[rand.Intn(len(text))])
}

func FaceMe() string {
	return face[rand.Intn(len(face))]
}

// A List to hold images
type List struct {
	Images    []*Image
	Timestamp int64
}

// Image is a generic image data
type Image struct {
	Width    int
	Height   int
	Id       string
	Caption  string
	TakenAt  int64
	LastMod  int64
	Sha256   string
	Name     string
	XMP      *XMP
	FileSize int64
}

// XMP meta data
type XMP struct {
	// Owner is the person who made this image
	Owner string
	// Faces declares who is visible in the image
	Faces []*Face

	Camera string
}

// A Face
type Face struct {
	// Name declares the persons name
	Name string
	// X is the relative x coordinate in the image normalized between 0 and 1
	X float64
	// Y is the relative y coordinate in the image normalized between 0 and 1
	Y float64

	// R is the relative radius around the x|y point, normalized between 0 and 1
	R float64
}

// GenerateImageMetaData creates some random image data
func GenerateImageMetaData(howMany int) *List {
	list := &List{}
	for i := 0; i < howMany; i++ {
		img := &Image{}
		img.Name = fmt.Sprintf("IMG_%06d.jpg", i)
		img.Id = CreateString(16)
		w, h := Resolution()
		img.Width = w
		img.Height = h
		img.TakenAt = Time(i, 12)
		img.LastMod = Time(i, 13)
		img.Sha256 = CreateString(16)
		img.FileSize = rand.Int63()
		img.Caption = Caption(i)
		img.XMP = &XMP{}
		img.XMP.Camera = Camera()
		img.XMP.Owner = FaceMe()
		for f := 0; f < rand.Intn(5); i++ {
			face := &Face{}
			face.Name = FaceMe()
			face.R = rand.NormFloat64()
			face.X = rand.NormFloat64()
			face.Y = rand.NormFloat64()
			img.XMP.Faces = append(img.XMP.Faces, face)
		}
		list.Images = append(list.Images, img)
	}
	list.Timestamp = Time(1, 1)
	return list
}
