PREAMBLE_SIZE = 25

class EncryptedData
  attr_accessor :data

  def initialize(data)
    @data = data
  end

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

encrypted_data = EncryptedData.new(File.open("input").map { |line| line.to_i })

(PREAMBLE_SIZE..encrypted_data.data.length - 1).each do |i|
  is_weak = encrypted_data.check_if_number_is_weak_at_index(i)
  puts "#{encrypted_data.data[i]} is weak? #{is_weak ? "yes" : "no"}"
end
