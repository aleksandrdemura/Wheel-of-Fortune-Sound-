<?php
// wheel.php — PHP версия

$prizes = ["500", "200", "100", "50", "20", "10", "5", "2"];

function play_sound($freq, $duration) {
    exec("beep -f $freq -l $duration");
}

echo "\033[36m🎡 Wheel of Fortune (Sound) (PHP)\033[0m\n";
echo "Призы: " . implode(', ', $prizes) . "\n";
echo "\nНажмите Enter, чтобы крутить колесо...";
fgets(STDIN);

$total = count($prizes);
echo "\nКрутим...";
for ($i = 0; $i < 20; $i++) {
    $idx = $i % $total;
    play_sound(200 + $idx * 50, 50);
    usleep(50000);
    echo ".";
}
echo "\n";

$win_idx = rand(0, $total - 1);
$prize = $prizes[$win_idx];
// финал
play_sound(400, 150);
usleep(100000);
play_sound(600, 150);
usleep(100000);
play_sound(800, 150);

echo "\033[32m🎉 Вы выиграли: $prize! 🎉\033[0m\n";
?>
