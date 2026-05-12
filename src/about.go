package main

import (
	"time"

	"github.com/skip2/go-qrcode"
)

// showAboutScreen displays the about page with version info, credits, and donation QR code
func (app *MiyooPod) showAboutScreen() {
	dc := app.DC

	// Set background
	dc.SetHexColor(app.CurrentTheme.BG)
	dc.Clear()

	// Title
	dc.SetFontFace(app.FontTitle)
	dc.SetHexColor(app.CurrentTheme.HeaderTxt)
	dc.DrawStringAnchored("MiyooPod 정보", SCREEN_WIDTH/2, 30, 0.5, 0.5)

	// Version and author info
	dc.SetFontFace(app.FontMenu)
	dc.SetHexColor(app.CurrentTheme.ItemTxt)

	yPos := 80
	dc.DrawStringAnchored("버전: "+APP_VERSION, SCREEN_WIDTH/2, float64(yPos), 0.5, 0.5)

	yPos += 30
	dc.DrawStringAnchored("제작자: "+APP_AUTHOR, SCREEN_WIDTH/2, float64(yPos), 0.5, 0.5)

	yPos += 40
	dc.SetFontFace(app.FontSmall)
	dc.SetHexColor(app.CurrentTheme.Dim)
	dc.DrawStringAnchored("Miyoo Mini용 음악 플레이어", SCREEN_WIDTH/2, float64(yPos), 0.5, 0.5)

	yPos += 20
	dc.DrawStringAnchored("클래식 음악 플레이어에서 영감", SCREEN_WIDTH/2, float64(yPos), 0.5, 0.5)

	// Generate QR code
	qr, err := qrcode.New(SUPPORT_URL, qrcode.Medium)
	if err == nil {
		qrSize := 150
		qrImg := qr.Image(qrSize)

		// Draw QR code
		qrX := (SCREEN_WIDTH - qrSize) / 2
		qrY := 220
		app.fastBlitImage(qrImg, qrX, qrY)

		// QR code label
		dc.SetFontFace(app.FontSmall)
		dc.SetHexColor(app.CurrentTheme.ItemTxt)
		dc.DrawStringAnchored("프로젝트 후원", SCREEN_WIDTH/2, float64(qrY+qrSize+20), 0.5, 0.5)
	}

	// Instructions
	dc.SetFontFace(app.FontSmall)
	dc.SetHexColor(app.CurrentTheme.Dim)
	dc.DrawStringAnchored("B 버튼: 돌아가기", SCREEN_WIDTH/2, SCREEN_HEIGHT-20, 0.5, 0.5)

	app.triggerRefresh()

	// Wait for B button to return
	app.waitForAboutExit()
}

// waitForAboutExit waits for the user to press B to exit the about screen
func (app *MiyooPod) waitForAboutExit() {
	for app.Running {
		key := Key(C_GetKeyPress())
		if key == NONE {
			time.Sleep(33 * time.Millisecond)
			continue
		}

		if key == B || key == MENU {
			// Return to menu
			app.setScreen(ScreenMenu)
			app.drawMenuScreen()
			return
		}
	}
}
