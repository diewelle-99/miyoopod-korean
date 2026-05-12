package main

import (
	"encoding/binary"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func estimateAudioDuration(path string) float64 {
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".mp3":
		return estimateMP3Duration(path)
	case ".flac":
		return estimateFLACDurationLight(path)
	default:
		// Avoid full-file reads during library scan.
		// OGG duration fallback is intentionally disabled for scan stability.
		return 0
	}
}

func estimateMP3Duration(path string) float64 {
	info, err := os.Stat(path)
	if err != nil || info.Size() <= 0 {
		return 0
	}

	f, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer f.Close()

	buf := make([]byte, 256*1024)
	n, _ := f.Read(buf)
	buf = buf[:n]

	start := 0
	if len(buf) >= 10 && string(buf[:3]) == "ID3" {
		tagSize := int(buf[6]&0x7F)<<21 |
			int(buf[7]&0x7F)<<14 |
			int(buf[8]&0x7F)<<7 |
			int(buf[9]&0x7F)
		start = 10 + tagSize
		if start >= len(buf) {
			start = 10
		}
	}

	for i := start; i+4 <= len(buf); i++ {
		h := binary.BigEndian.Uint32(buf[i : i+4])

		if h&0xFFE00000 != 0xFFE00000 {
			continue
		}

		versionID := int((h >> 19) & 0x3)
		layerID := int((h >> 17) & 0x3)
		bitrateIndex := int((h >> 12) & 0xF)
		sampleRateIndex := int((h >> 10) & 0x3)

		if versionID == 1 || layerID == 0 || bitrateIndex == 0 || bitrateIndex == 15 || sampleRateIndex == 3 {
			continue
		}

		kbps := mp3BitrateKbps(versionID, layerID, bitrateIndex)
		if kbps <= 0 {
			continue
		}

		audioBytes := info.Size() - int64(i)
		if audioBytes <= 0 {
			return 0
		}

		return float64(audioBytes*8) / float64(kbps*1000)
	}

	return 0
}

func mp3BitrateKbps(versionID, layerID, idx int) int {
	if versionID == 3 {
		switch layerID {
		case 3:
			return []int{0, 32, 64, 96, 128, 160, 192, 224, 256, 288, 320, 352, 384, 416, 448}[idx]
		case 2:
			return []int{0, 32, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, 384}[idx]
		case 1:
			return []int{0, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320}[idx]
		}
	} else {
		switch layerID {
		case 3:
			return []int{0, 32, 48, 56, 64, 80, 96, 112, 128, 144, 160, 176, 192, 224, 256}[idx]
		case 2, 1:
			return []int{0, 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160}[idx]
		}
	}
	return 0
}

func estimateFLACDurationLight(path string) float64 {
	f, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer f.Close()

	header := make([]byte, 4)
	if _, err := io.ReadFull(f, header); err != nil {
		return 0
	}
	if string(header) != "fLaC" {
		return 0
	}

	for i := 0; i < 16; i++ {
		blockHeader := make([]byte, 4)
		if _, err := io.ReadFull(f, blockHeader); err != nil {
			return 0
		}

		blockType := blockHeader[0] & 0x7F
		length := int(blockHeader[1])<<16 | int(blockHeader[2])<<8 | int(blockHeader[3])

		if blockType == 0 {
			block := make([]byte, length)
			if _, err := io.ReadFull(f, block); err != nil {
				return 0
			}
			if len(block) < 18 {
				return 0
			}

			x := binary.BigEndian.Uint64(block[10:18])
			sampleRate := int((x >> 44) & 0xFFFFF)
			totalSamples := x & 0xFFFFFFFFF

			if sampleRate <= 0 || totalSamples == 0 {
				return 0
			}
			return float64(totalSamples) / float64(sampleRate)
		}

		if _, err := f.Seek(int64(length), io.SeekCurrent); err != nil {
			return 0
		}

		if blockHeader[0]&0x80 != 0 {
			break
		}
	}

	return 0
}
