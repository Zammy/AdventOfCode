PREAMBLE_SIZE = 25
INPUT_FILE_NAME = "input"

class EncryptedData
  attr_accessor :data

  def initialize(data)
    @data = data
    @weak_number = nil
  end

  def find_weak_number
    (PREAMBLE_SIZE..@data.length - 1).each do |i|
      is_weak = check_if_number_is_weak_at_index(i)
      if is_weak
        @weak_number = @data[i]
        puts "Found weak number #{@weak_number}"
        break
      end
    end
  end

  def find_range_that_sums_to_weak_number
    @data.each_index do |start_i|
      sum = 0
      (start_i..@data.length - 1).each do |i|
        sum += @data[i]
        if sum == @weak_number
          puts "Found range #{start_i..i}"
          return @data.slice(start_i, i - start_i + 1)
        end
      end
    end
  end

  private

  def check_if_number_is_weak_at_index(num_index)
    num = @data[num_index]
    (num_index - PREAMBLE_SIZE..num_index - 1).each do |i|
      (i + 1..num_index - 1).each do |y|
        sum = @data[i] + @data[y]
        if sum == num
          return false
        end
      end
    end
    true
  end
end

encrypted_data = EncryptedData.new(File.open(INPUT_FILE_NAME).map { |line| line.to_i })
encrypted_data.find_weak_number
numbers = encrypted_data.find_range_that_sums_to_weak_number
min = numbers.min
max = numbers.max
puts "Min num:#{min} Max num: #{max} Sum:#{min + max}"
