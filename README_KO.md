# MiyooPod Korean Patch

This repository is a Korean-localized modification of MiyooPod by Danilo Fragoso.

## Status

Current target:
- Miyoo Mini / Miyoo Mini Plus style ARM build
- Korean UI localization
- Korean font support
- Korean search initial consonants
- Playback progress fallback
- Audio duration fallback
- Track-specific now playing artwork test

## Main Changes

- Replaced UI font with Korean-capable Noto Sans KR font.
- Localized menus and settings into Korean.
- Added Korean initial-consonant search support.
- Added playback position fallback when SDL_mixer does not report position.
- Added duration fallback for MP3/FLAC when SDL_mixer returns 0.
- Added experimental per-track artwork display in Now Playing.
- Built successfully on Ubuntu 20.04 WSL with arm-linux-gnueabihf-gcc.
- Verified binary requires GLIBC_2.4 only.

## Build

```bash
make go
make updater