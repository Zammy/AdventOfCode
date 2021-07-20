STEP = 3
adapters_array = File.open("input").map { |line| line.to_i }
adapters_array.sort!

options = Array.new(adapters_array.length, 0)
adapters_array.unshift 0
adapters_array << adapters_array[-1] + STEP
options.unshift 1
options << 0

adapters_array.each_index do |i|
  next if i == 0
  value = adapters_array[i]

  target_index = [i - STEP, 0].max
  (target_index..i - 1).step 1 do |y|
    prev_value = adapters_array[y]
    if value - prev_value <= STEP
      options[i] += options[y]
    end
  end
end

puts "Possible arrangements: #{options[-1]}"
