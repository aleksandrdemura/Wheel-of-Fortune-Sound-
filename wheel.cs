// wheel.cs — C# версия

using System;
using System.Diagnostics;
using System.Threading;

class Wheel
{
    static string[] prizes = { "500", "200", "100", "50", "20", "10", "5", "2" };

    static void PlaySound(int freq, int duration)
    {
        try
        {
            Process.Start("beep", $"-f {freq} -l {duration}");
        }
        catch
        {
            Console.Beep(freq, duration);
        }
    }

    static void Main()
    {
        Console.WriteLine("\u001B[36m🎡 Wheel of Fortune (Sound) (C#)\u001B[0m");
        Console.WriteLine($"Призы: {string.Join(", ", prizes)}");
        Console.Write("\nНажмите Enter, чтобы крутить колесо...");
        Console.ReadLine();

        int total = prizes.Length;
        Console.Write("\nКрутим...");
        for (int i = 0; i < 20; i++)
        {
            int idx = i % total;
            PlaySound(200 + idx * 50, 50);
            Thread.Sleep(50);
            Console.Write(".");
        }
        Console.WriteLine();

        Random rand = new Random();
        int winIdx = rand.Next(total);
        string prize = prizes[winIdx];
        // финал
        PlaySound(400, 150);
        Thread.Sleep(100);
        PlaySound(600, 150);
        Thread.Sleep(100);
        PlaySound(800, 150);

        Console.WriteLine($"\n\u001B[32m🎉 Вы выиграли: {prize}! 🎉\u001B[0m");
    }
}
