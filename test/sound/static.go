package main

import (
	"time"

	"github.com/skycoin/cx-game/physics"
	"github.com/skycoin/cx-game/sound"
)

func main() {
	sound.Init()
	sound.LoadSound("boo", "boo.wav")
	soundId := sound.PlaySound("boo", sound.SoundOptions{
		Looped: true,
	})
	go func() {
		time.Sleep(1500 * time.Millisecond)
		sound.StopSound(soundId)
	}()
	pos := physics.Vec2{-5, 0}
	for {
		pos.X += 0.5
		sound.SetListenerPosition(pos)
		sound.Update()
		time.Sleep(50 * time.Millisecond)
	}
}
