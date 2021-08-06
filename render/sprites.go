package render

import (
	"log"
)

var sprites = []Sprite{}
var spriteNamesToIDs = map[string]int{}

func addSprite(sprite Sprite) int {
	id := len(sprites)
	sprites = append(sprites, sprite)
	return id
}

func RegisterSprite(sprite Sprite) {
	if sprite.Name=="" {
		log.Fatal("cannot register sprite with empty name")
	}
	spriteNamesToIDs[sprite.Name] = addSprite(sprite)
}
