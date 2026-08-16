

### 1. `wheel.py` (Python)

```python
# wheel.py — Python версия

import sys
import time
import random
import platform
from colorama import init, Fore, Style

init(autoreset=True)

PRIZES = ["500", "200", "100", "50", "20", "10", "5", "2"]
SOUND_ENABLED = True

def play_sound(freq=440, duration=200):
    if not SOUND_ENABLED:
        return
    system = platform.system()
    if system == 'Windows':
        import winsound
        winsound.Beep(freq, duration)
    else:
        try:
            import os
            os.system(f'beep -f {freq} -l {duration}')
        except:
            pass

def spin_wheel():
    print(f"\n{Fore.CYAN}🎡 Wheel of Fortune (Sound) (Python){Style.RESET_ALL}")
    print(f"Призы: {', '.join(PRIZES)}")
    input("\nНажмите Enter, чтобы крутить колесо...")

    total = len(PRIZES)
    # Анимация вращения
    print("\nКрутим...", end="", flush=True)
    for i in range(20):
        idx = i % total
        play_sound(200 + idx * 50, 50)
        time.sleep(0.05)
        print(".", end="", flush=True)
    print()

    # Выбор случайного приза
    win_idx = random.randint(0, total - 1)
    prize = PRIZES[win_idx]
    # Финальный звук
    for f in [400, 600, 800]:
        play_sound(f, 150)
        time.sleep(0.1)
    print(f"\n{Fore.GREEN}🎉 Вы выиграли: {prize}! 🎉{Style.RESET_ALL}")

def main():
    try:
        spin_wheel()
    except KeyboardInterrupt:
        print("\nВыход...")

if __name__ == "__main__":
    main()
