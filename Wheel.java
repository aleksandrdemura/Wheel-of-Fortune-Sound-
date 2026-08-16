// Wheel.java — Java версия

import java.util.Scanner;
import java.util.Random;

public class Wheel {
    private static final String[] PRIZES = {"500", "200", "100", "50", "20", "10", "5", "2"};

    private static void playSound(int freq, int duration) {
        try {
            Runtime.getRuntime().exec(new String[]{"beep", "-f", String.valueOf(freq), "-l", String.valueOf(duration)});
        } catch (Exception e) {
            // ignored
        }
    }

    public static void main(String[] args) throws InterruptedException {
        Scanner scanner = new Scanner(System.in);
        System.out.println("\u001B[36m🎡 Wheel of Fortune (Sound) (Java)\u001B[0m");
        System.out.print("Призы: ");
        for (String p : PRIZES) System.out.print(p + " ");
        System.out.println();
        System.out.print("\nНажмите Enter, чтобы крутить колесо...");
        scanner.nextLine();

        int total = PRIZES.length;
        System.out.print("\nКрутим...");
        for (int i = 0; i < 20; i++) {
            int idx = i % total;
            playSound(200 + idx * 50, 50);
            Thread.sleep(50);
            System.out.print(".");
        }
        System.out.println();

        Random rand = new Random();
        int winIdx = rand.nextInt(total);
        String prize = PRIZES[winIdx];
        // финал
        playSound(400, 150);
        Thread.sleep(100);
        playSound(600, 150);
        Thread.sleep(100);
        playSound(800, 150);

        System.out.println("\n\u001B[32m🎉 Вы выиграли: " + prize + "! 🎉\u001B[0m");
        scanner.close();
    }
}
