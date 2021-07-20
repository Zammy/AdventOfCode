class AdapterArray
  attr_accessor :connections

  def initialize(data)
    @adapters = data.sort
    @connections = {}
  end

  def connect_to_outlet
    prev_joltage = 0
    @adapters.each do |adapter|
      diff = adapter - prev_joltage
      prev_joltage = adapter
      if @connections.has_key? diff
        @connections[diff] += 1
      else
        @connections[diff] = 1
      end
    end
    @connections[3] += 1
  end
end

adapters_array = AdapterArray.new(File.open("input").map { |line| line.to_i })
adapters_array.connect_to_outlet

puts adapters_array.connections
puts "Answer: #{adapters_array.connections[1] * adapters_array.connections[3]}"
