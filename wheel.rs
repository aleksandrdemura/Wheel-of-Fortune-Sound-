// wheel.rs — Rust версия

use rand::Rng;
use std::io::{self, Write};
use std::process::Command;
use std::thread;
use std::time::Duration;
use colored::*;

const PRIZES: [&str; 8] = ["500", "200", "100", "50", "20", "10", "5", "2"];

fn play_sound(freq: u32, duration: u32) {
    let _ = Command::new("beep")
        .args(&["-f", &freq.to_string(), "-l", &duration.to_string()])
        .status();
}

fn main() {
    println!("{}", "🎡 Wheel of Fortune (Sound) (Rust)".cyan());
    println!("Призы: {}", PRIZES.join(", "));
    print!("\nНажмите Enter, чтобы крутить колесо...");
    io::stdout().flush().unwrap();
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();

    let total = PRIZES.len();
    print!("\nКрутим...");
    for i in 0..20 {
        let idx = i % total;
        play_sound(200 + (idx as u32) * 50, 50);
        thread::sleep(Duration::from_millis(50));
        print!(".");
        io::stdout().flush().unwrap();
    }
    println!();

    let win_idx = rand::thread_rng().gen_range(0..total);
    let prize = PRIZES[win_idx];
    // финал
    play_sound(400, 150);
    thread::sleep(Duration::from_millis(100));
    play_sound(600, 150);
    thread::sleep(Duration::from_millis(100));
    play_sound(800, 150);

    println!("\n{}", format!("🎉 Вы выиграли: {}! 🎉", prize).green());
}
