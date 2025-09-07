package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

// Game - основная структура игры
type Game struct{}

// Update вызывается каждый кадр для обновления игровой логики
func (g *Game) Update() error {
	return nil
}

// Draw вызывается каждый кадр для отрисовки графики
func (g *Game) Draw(screen *ebiten.Image) {
	// Рисуем фон (просто заливаем цветом)
	screen.Fill(color.RGBA{R: 20, G: 20, B: 40, A: 255})

	// Рисуем заголовок
	ebitenutil.DebugPrintAt(screen, "Меню Игр", 250, 100)

	// Рисуем кнопки
	drawButton(screen, 250, 200, "Начать Игру")
	drawButton(screen, 250, 250, "Настройки")
	drawButton(screen, 250, 300, "Выход")
}

// Layout - используется для масштабирования
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

// drawButton - вспомогательная функция для отрисовки кнопки
func drawButton(screen *ebiten.Image, x, y int, text string) {
	// Создаем прямоугольник для кнопки
	ebitenutil.DrawRect(screen, float64(x), float64(y), 150, 40, color.RGBA{R: 50, G: 100, B: 200, A: 255})
	// Рисуем текст на кнопке
	ebitenutil.DebugPrintAt(screen, text, x+10, y+10)
}

func main() {
	ebiten.SetWindowTitle("Меню Игр")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
