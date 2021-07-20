class AdapterArray
  attr_accessor :itterations

  def initialize(data)
    @adapters = data.sort

    @adapters.unshift 0
    @adapters << @adapters[-1] + 3

    @itterations = 1
  end

  def find_all_arrangements(index)
    first_option = true

    (index + 1..@adapters.length - 1).step 1 do |i|
      if @adapters[i] - @adapters[index] <= 3
        if first_option
          first_option = false
        else
          @itterations += 1
        end
        find_all_arrangements i
      else
        break
      end
    end
  end
end

adapters_array = AdapterArray.new(File.open("input").map { |line| line.to_i })
adapters_array.find_all_arrangements 0

puts "Possible arrangments: #{adapters_array.itterations}"
