package main

import "time"

// startPlaybackPoller checks audio state and updates progress display.
// If SDL_mixer cannot report music position correctly, fall back to a local timer.
func (app *MiyooPod) startPlaybackPoller() {
	lastDrawnSecond := -1
	tickCount := 0
	saveTickCount := 0

	lastTick := time.Now()
	lastTrackPath := ""
	lastAudioPosition := 0.0

	for app.Running {
		now := time.Now()
		elapsed := now.Sub(lastTick).Seconds()
		if elapsed <= 0 || elapsed > 2.5 {
			elapsed = 1.0
		}
		lastTick = now

		if app.Playing != nil && app.Playing.State != StateStopped {
			trackPath := ""
			if app.Playing.Track != nil {
				trackPath = app.Playing.Track.Path
			}

			if trackPath != lastTrackPath {
				lastTrackPath = trackPath
				lastAudioPosition = 0
				lastDrawnSecond = -1
				lastTick = now
			}

			state := audioGetState()

			if state.Duration > 0 {
				app.Playing.Duration = state.Duration
				if app.Playing.Track != nil && app.Playing.Track.Duration == 0 {
					app.Playing.Track.Duration = state.Duration
				}
			}

			if state.IsPaused && app.Playing.State != StatePaused {
				app.Playing.State = StatePaused
				app.NPCacheDirty = true
				app.requestRedraw()
			} else if state.IsPlaying && app.Playing.State != StatePlaying {
				app.Playing.State = StatePlaying
				app.NPCacheDirty = true
				app.requestRedraw()
			}

			if state.IsPlaying {
				// Normal path: SDL_mixer reports position.
				if state.Position > 0 && state.Position >= lastAudioPosition {
					app.Playing.Position = state.Position
					lastAudioPosition = state.Position
				} else {
					// Fallback path: SDL_mixer returns 0/-1 forever.
					app.Playing.Position += elapsed
					if app.Playing.Duration > 0 && app.Playing.Position > app.Playing.Duration {
						app.Playing.Position = app.Playing.Duration
					}
				}
			} else if state.Position > 0 {
				app.Playing.Position = state.Position
				lastAudioPosition = state.Position
			}

			if state.Finished {
				lastAudioPosition = 0
				app.handleTrackEnd()
			}

			if app.CurrentScreen == ScreenNowPlaying {
				currentSecond := int(app.Playing.Position)
				if currentSecond != lastDrawnSecond {
					lastDrawnSecond = currentSecond
					app.updateProgressBarOnly()
				}
			}

			if app.CurrentScreen == ScreenLyrics && app.LyricsCachedLRC != nil && app.LastKey == NONE {
				activeLRC := activeLRCIndex(app.LyricsCachedLRC, app.Playing.Position)
				if activeLRC != app.LyricsLastActiveLRC {
					app.LyricsLastActiveLRC = activeLRC
					app.requestRedraw()
				}
			}

			tickCount++
			if tickCount >= 5 {
				audioFlushBuffers()
				tickCount = 0
			}

			saveTickCount++
			if saveTickCount >= 3 {
				app.savePlaybackState()
				saveTickCount = 0
			}
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func (app *MiyooPod) mpvLoadFile(path string) error {
	err := audioLoadFile(path)
	if err != nil {
		return err
	}
	return audioPlay()
}

func (app *MiyooPod) mpvTogglePause() {
	audioTogglePause()
}

func (app *MiyooPod) mpvStop() {
	audioStop()
}

func (app *MiyooPod) mpvSeek(seconds float64) {
	if app.Playing == nil {
		return
	}

	newPos := app.Playing.Position + seconds
	if newPos < 0 {
		newPos = 0
	}
	if newPos > app.Playing.Duration && app.Playing.Duration > 0 {
		newPos = app.Playing.Duration
	}

	audioSeek(newPos)
	app.Playing.Position = newPos
	app.NPCacheDirty = true
	app.requestRedraw()
}
