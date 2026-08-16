// wheel.go — Go версия

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var prizes = []string{"500", "200", "100", "50", "20", "10", "5", "2"}

func playSound(freq, duration int) {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("powershell", "-Command", fmt.Sprintf("[System.Console]::Beep(%d, %d)", freq, duration))
		cmd.Run()
	default:
		cmd := exec.Command("beep", "-f", fmt.Sprintf("%d", freq), "-l", fmt.Sprintf("%d", duration))
		cmd.Run()
	}
}

func spinWheel() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\x1b[36m🎡 Wheel of Fortune (Sound) (Go)\x1b[0m")
	fmt.Printf("Призы: %v\n", prizes)
	fmt.Print("\nНажмите Enter, чтобы крутить колесо...")
	reader.ReadString('\n')

	total := len(prizes)
	fmt.Print("\nКрутим...")
	for i := 0; i < 20; i++ {
		idx := i % total
		playSound(200+idx*50, 50)
		time.Sleep(50 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()

	rand.Seed(time.Now().UnixNano())
	winIdx := rand.Intn(total)
	prize := prizes[winIdx]
	// финал
	playSound(400, 150)
	time.Sleep(100 * time.Millisecond)
	playSound(600, 150)
	time.Sleep(100 * time.Millisecond)
	playSound(800, 150)

	fmt.Printf("\n\x1b[32m🎉 Вы выиграли: %s! 🎉\x1b[0m\n", prize)
}

func main() {
	spinWheel()
}
