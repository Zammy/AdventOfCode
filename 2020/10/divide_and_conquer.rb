class AdapterArray
  attr_accessor :adapters

  def initialize(data)
    @adapters = data.sort

    @adapters.unshift 0
    @adapters << @adapters[-1] + 3
  end

  def find_all_arrangements
    first_index = 0
    max_arrangements = 1
    while @adapters.length > first_index
      (first_index..@adapters.length - 1).step 1 do |i|
        if i + 1 == @adapters.length || @adapters[i + 1] - @adapters[i] == 3
          arrangements = 1
          arrangements = find_all_arrangements_between(first_index, i) if first_index != i
          puts "From #{first_index} to #{i} -> #{arrangements}"
          max_arrangements = max_arrangements * arrangements
          first_index = i + 1
        end
      end
    end
    max_arrangements
  end

  private

  def find_all_arrangements_between(start_index, end_index)
    options = 0
    (start_index + 1..end_index).step 1 do |i|
      if @adapters[i] - @adapters[start_index] <= 3
        options += 1
        if end_index != i
          deep_options = find_all_arrangements_between(i, end_index) - 1
          options += [0, deep_options].max
        end
      else
        break
      end
    end
    options
  end
end

adapters_array = AdapterArray.new(File.open("input").map { |line| line.to_i })
puts "Possible arrangements: #{adapters_array.find_all_arrangements}"
