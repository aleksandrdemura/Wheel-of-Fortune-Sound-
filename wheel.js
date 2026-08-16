// wheel.js — JavaScript версия

const readline = require('readline');
const { exec } = require('child_process');

const prizes = ["500", "200", "100", "50", "20", "10", "5", "2"];

function playSound(freq, duration) {
    const cmd = `beep -f ${freq} -l ${duration}`;
    exec(cmd, (error) => { if (error) { /* игнорируем */ } });
}

function spinWheel() {
    const rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout
    });

    console.log('\x1b[36m🎡 Wheel of Fortune (Sound) (JavaScript)\x1b[0m');
    console.log(`Призы: ${prizes.join(', ')}`);

    rl.question('\nНажмите Enter, чтобы крутить колесо...', () => {
        const total = prizes.length;
        process.stdout.write('\nКрутим...');
        for (let i = 0; i < 20; i++) {
            const idx = i % total;
            playSound(200 + idx * 50, 50);
            process.stdout.write('.');
        }
        console.log();

        const winIdx = Math.floor(Math.random() * total);
        const prize = prizes[winIdx];
        // финал
        playSound(400, 150);
        setTimeout(() => {
            playSound(600, 150);
            setTimeout(() => {
                playSound(800, 150);
                console.log(`\n\x1b[32m🎉 Вы выиграли: ${prize}! 🎉\x1b[0m`);
                rl.close();
            }, 100);
        }, 100);
    });
}

spinWheel();
