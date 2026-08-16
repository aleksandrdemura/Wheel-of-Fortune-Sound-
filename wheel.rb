# wheel.rb — Ruby версия

PRIZES = ["500", "200", "100", "50", "20", "10", "5", "2"]

def play_sound(freq, duration)
  system("beep -f #{freq} -l #{duration}")
end

def spin_wheel
  puts "\e[36m🎡 Wheel of Fortune (Sound) (Ruby)\e[0m"
  puts "Призы: #{PRIZES.join(', ')}"
  print "\nНажмите Enter, чтобы крутить колесо..."
  gets

  total = PRIZES.size
  print "\nКрутим..."
  20.times do |i|
    idx = i % total
    play_sound(200 + idx * 50, 50)
    sleep 0.05
    print "."
  end
  puts

  win_idx = rand(total)
  prize = PRIZES[win_idx]
  # финал
  play_sound(400, 150)
  sleep 0.1
  play_sound(600, 150)
  sleep 0.1
  play_sound(800, 150)

  puts "\e[32m🎉 Вы выиграли: #{prize}! 🎉\e[0m"
end

spin_wheel if __FILE__ == $0
