class PlaneSeat
  def initialize(line_data)
    @row_data = line_data[0, 7].split("").map do |char|
      next :low if char == 'F'
      next :high if char == 'B'
      throw :bad_data_in_row
    end

    @seat_data = line_data[7, 3].split("").map do |char|
      next :low if char == 'L'
      next :high if char == 'R'
      throw :bad_data_in_seat
    end
  end

  def get_row_number
    PlaneSeat.binary_search([0, 127], @row_data.clone)[0]
  end

  def get_seat_number
    PlaneSeat.binary_search([0, 7], @seat_data.clone)[0]
  end

  def get_seat_id
    get_row_number * 8 + get_seat_number
  end

  def self.binary_search(range, search_pointers)
    low_high = search_pointers.shift
    return range unless low_high != nil
    diff = (range[1] - range[0])/2.0
    new_limit = range[0] + diff
    if low_high == :low
      range[1] = new_limit.to_i
    elsif low_high == :high
      range[0] = new_limit.to_i + 1
    end
    PlaneSeat.binary_search(range, search_pointers)
  end
end

ids = []
File.open("input").each do |line|
  seat = PlaneSeat.new(line)
  ids << seat.get_seat_id
end

ids.sort!
lowest_id = ids[0]
highest_id = ids[-1]

for i in lowest_id..highest_id
  if ids.none? i 
    puts "My seat index should be #{i}"
  end
end
