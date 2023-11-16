package main

import (
	"github.com/RealSilentPirate/chipper/pkg/initwirepod"
	stt "github.com/RealSilentPirate/chipper/pkg/wirepod/stt/vosk"
)

func main() {
	initwirepod.StartFromProgramInit(stt.Init, stt.STT, stt.Name)
}
