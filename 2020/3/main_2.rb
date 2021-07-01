class Point
  attr_accessor :x,:y

  def initialize(x = 0, y = 0)
    @x = x
    @y = y
  end
  
  def +(anotherPoint)
    Point.new(@x + anotherPoint.x, @y + anotherPoint.y)
  end

  def to_s
    "(#{@x},#{@y}"
  end
end

class MountainSlope
  def initialize(filepath)
    @data = [];
    File.open(filepath).each_with_index do |line, x|
      line.each_char.each_with_index do |c, y|
        if @data.length <= x 
          @data[x] = []
        end
        @data[x][y] = :tree if c == "#"
        @data[x][y] = :free if c == '.'
      end 
    end
  end

  def is_tree(x, y)
    y = convert_y_coordinate(y)
    @data[x][y] == :tree
  end

  def is_on_slope(x)
    @data.length > x;
  end

  def to_s
    @data.each do |line|
      line.each do |char|
        print "#" if char == :tree
        print "." if char == :free
      end
      print "\n"
    end
  end

  private
  def convert_y_coordinate(y)
    y % @data[0].length;
  end
end

def count_trees_with_delta(slope, delta)
  point = Point.new
  trees = 0
  loop do
    point += delta;
    break unless slope.is_on_slope(point.x)
    if slope.is_tree(point.x, point.y)
      trees += 1 
      # puts "#{point} is tree"
    end
  end
  trees
end

slope = MountainSlope.new("input");
deltas = [Point.new(1, 1),Point.new(1, 3),Point.new(1, 5),Point.new(1, 7),Point.new(2, 1)]

product = 1
deltas.each do |delta|
  product *= count_trees_with_delta(slope, delta)
end
puts product