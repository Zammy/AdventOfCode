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

slope = MountainSlope.new("input");
x=0
y=0
trees = 0
loop do
  x += 1
  y += 3
  break unless slope.is_on_slope(x)
  if slope.is_tree(x, y)
    trees += 1 
    puts "(#{x},#{y}) is tree"
  end
end

puts trees