data = []
File.open("input").each do |line|
  data << line.to_i
end

def find_nums_that_sum_to_2020(data)
  data.each_with_index do |l1, i|
    data.drop(i + 1).each do |l2|
      if (l1 + l2) == 2020
        return l1 * l2
      end
    end
  end
end

result = find_nums_that_sum_to_2020(data)

puts result
