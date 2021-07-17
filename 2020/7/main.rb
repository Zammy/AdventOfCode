class Bag
  attr_accessor :bags, :adjectives

  def initialize(bag_adjectives)
    r = bag_adjectives.match(/(\w+ \w+) bags*/)
    unless r
      puts bag_adjectives
      throw :bad_data
    end
    @adjectives = r[1]
    @bags = []
  end

  def add_bag(bag)
    @bags << bag
  end

  def does_it_contain_bag?(bag)
    @bags.any? { |other_bag| other_bag.adjectives == bag.adjectives }
  end

  def to_s
    "#{@adjectives} bag"
  end
end

bags = []
File.open("input").each do |line|
  split = line.split("contain")
  base_bag = Bag.new(split[0])
  split = split[1].split(",")
  split.each do |bag_str|
    bag = Bag.new(bag_str)
    base_bag.add_bag(bag)
  end
  bags << base_bag
end

bags_count = []
my_bag = Bag.new("shiny gold bag")
bags_to_check = [my_bag]
while bags_to_check.length > 0
  bag_to_check = bags_to_check.shift
  puts "Looking for #{bag_to_check}"
  bags.each do |bag|
    if bag.does_it_contain_bag? bag_to_check
      bags_to_check << bag
      puts "  found inside #{bag}"
      bags_count << bag
    end
  end
end

puts "------------------------------------------------"
puts "All bag options count: #{bags_count.uniq.length}"
