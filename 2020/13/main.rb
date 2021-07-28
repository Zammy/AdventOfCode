input = File.open("input").readlines
depart_time = input[0].to_i
bus_ids = input[1].split(",").filter { |id| id != "x" }.map { |id| id.to_i }

smallest_wait_time = 9999999
bus_id_to_take = 0
bus_ids.each do |bus_id|
  parts = depart_time / bus_id
  earliest = (parts + 1) * bus_id
  wait_time = earliest - depart_time
  if wait_time < smallest_wait_time
    smallest_wait_time = wait_time
    bus_id_to_take = bus_id
  end
end

puts "Bus id to take is #{bus_id_to_take} with wait of #{smallest_wait_time}"
puts "Answer: #{bus_id_to_take * smallest_wait_time}"
