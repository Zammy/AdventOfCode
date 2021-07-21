class Seat
  attr_accessor :state

  def initialize(state)
    set_state state
  end

  def set_state(state)
    @state = state
  end

  def ==(other_object)
    @state == other_object.state
  end

  def self.parse_state(char)
    case char
    when "#"
      :occupied
    when "L"
      :empty
    when "."
      :no_seat
    end
  end

  def self.state_to_s(state)
    case state
    when :occupied
      "#"
    when :empty
      "L"
    when :no_seat
      "."
    end
  end
end

class Plane
  attr_accessor :seats

  def initialize(seats = nil)
    if seats
      @seats = seats
    else
      @seats = []
    end
  end

  def add_row(seats)
    @seats << seats
  end

  def get_at(x, y)
    @seats[y][x]
  end

  def apply_rule(current_plane)
    changes = false
    (0..@seats.length - 1).step 1 do |y|
      row = @seats[y]
      (0..row.length - 1).step 1 do |x|
        occupied_around = current_plane.count_occupied_seats_around(x, y)
        # puts "#{occupied_around} occupied_around  (#{x},#{y}) "
        current_seat = @seats[y][x]
        if current_seat.state == :empty && occupied_around == 0
          changes = true
          # puts "  occuping"
          current_seat.set_state :occupied
        elsif current_seat.state == :occupied && occupied_around >= 4
          changes = true
          # puts "  emptying"
          current_seat.set_state :empty
        end
      end
    end
    changes
  end

  def clone
    Plane.new (Marshal.load(Marshal.dump(@seats)))
  end

  def to_s
    result = ""
    @seats.each do |row|
      row.each do |seat|
        result += "#{Seat.state_to_s seat.state}"
      end
      result += "\n"
    end
    result
  end

  def count_occupied_seats_around(x, y)
    occupied = 0

    to_check = [[-1, -1], [0, -1], [+1, -1],
                [-1, 0], [+1, 0],
                [-1, +1], [0, +1], [+1, +1]]
    to_check.each do |pair|
      check_x = x + pair[0]
      check_y = y + pair[1]
      if check_y < 0 || @seats.length <= check_y
        next
      end
      if check_x < 0 || @seats[check_y].length <= check_x
        next
      end
      if @seats[check_y][check_x].state == :occupied
        occupied += 1
      end
    end

    occupied
  end

  def count_occupied_seats
    occupied = 0

    @seats.each do |row|
      row.each do |seat|
        occupied += 1 if seat.state == :occupied
      end
    end

    occupied
  end
end

plane = Plane.new
File.open("input").map do |line|
  row = line
    .split("")
    .filter { |char| char != "\n" }
    .map { |char| Seat.new (Seat.parse_state char) }
  plane.add_row row
end

loop do
  # puts plane
  # puts "--------------"
  next_plane = plane.clone
  changes = next_plane.apply_rule plane
  if !changes
    break
  end
  plane = next_plane
end

puts "Occupied: #{plane.count_occupied_seats}"
