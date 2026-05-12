package main

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"github.com/dhowden/tag"
	"github.com/fogleman/gg"
)

func (app *MiyooPod) getNowPlayingTrackCover(size int) image.Image {
	if app == nil || app.Playing == nil || app.Playing.Track == nil {
		return nil
	}

	track := app.Playing.Track
	if track.Path == "" {
		return nil
	}

	if app.Coverflow == nil {
		return nil
	}
	if app.Coverflow.CoverCache == nil {
		app.Coverflow.CoverCache = make(map[string]image.Image)
	}

	info, statErr := os.Stat(track.Path)
	cacheStamp := ""
	if statErr == nil {
		cacheStamp = fmt.Sprintf("%d_%d", info.Size(), info.ModTime().Unix())
	}

	key := fmt.Sprintf("nowplaying_track_%s_%s_%d", track.Path, cacheStamp, size)
	if cached, ok := app.Coverflow.CoverCache[key]; ok {
		return cached
	}

	f, err := os.Open(track.Path)
	if err != nil {
		logMsg(fmt.Sprintf("[TRACK_ART] open failed: %v", err))
		return nil
	}
	defer f.Close()

	m, err := tag.ReadFrom(f)
	if err != nil {
		logMsg(fmt.Sprintf("[TRACK_ART] tag read failed: %v", err))
		return nil
	}

	pic := m.Picture()
	if pic == nil || len(pic.Data) == 0 {
		return nil
	}

	img, _, err := image.Decode(bytes.NewReader(pic.Data))
	if err != nil {
		logMsg(fmt.Sprintf("[TRACK_ART] decode failed: %v", err))
		return nil
	}

	b := img.Bounds()
	if b.Dx() <= 0 || b.Dy() <= 0 {
		return nil
	}

	dc := gg.NewContext(size, size)

	srcW := float64(b.Dx())
	srcH := float64(b.Dy())

	scale := float64(size) / srcW
	if srcH*scale < float64(size) {
		scale = float64(size) / srcH
	}

	newW := srcW * scale
	newH := srcH * scale

	x := (float64(size) - newW) / 2
	y := (float64(size) - newH) / 2

	dc.Push()
	dc.Translate(x, y)
	dc.Scale(scale, scale)
	dc.DrawImage(img, 0, 0)
	dc.Pop()

	result := dc.Image()
	app.Coverflow.CoverCache[key] = result

	return result
}
