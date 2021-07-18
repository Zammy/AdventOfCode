class Bag
  attr_accessor :bags, :adjectives

  def initialize(bag_adjectives)
    @adjectives = bag_adjectives
    @bags = {}
  end

  def add_bag(bag, count)
    @bags.store(bag, count)
  end

  def does_it_contain_bag?(bag)
    @bags.keys.any? { |other_bag| other_bag.adjectives == bag.adjectives }
  end

  def to_s
    "#{@adjectives} bag"
  end
end

class BagsAggregate
  attr_accessor :bags

  def initialize()
    @bags = {}
  end

  def get_bag_from_string(str)
    bag_adjectives = BagsAggregate.extract_bag_adjectives(str)
    bag = get_bag(bag_adjectives)
    return bag unless bag.nil?
    bag = Bag.new(bag_adjectives)
    @bags.store(bag_adjectives, bag)
    bag
  end

  def get_bag(adjectives)
    @bags[adjectives]
  end

  def count_bags_inside_bag(bag)
    bag_count = 0
    bag.bags.each do |bag_inside, count|
      bag_count += count + count * count_bags_inside_bag(bag_inside)
    end
    puts "Counted #{bag_count} bags inside #{bag}"
    bag_count
  end

  private

  MATCH_BAGS = /(\w+ \w+) bags*/
  def self.extract_bag_adjectives(str)
    r = str.match(MATCH_BAGS)
    unless r
      puts str
      throw :bad_data
    end
    r[1]
  end
end

bags = BagsAggregate.new
File.open("input").each do |line|
  split = line.split("contain")
  base_bag = bags.get_bag_from_string(split[0])
  split = split[1].split(",")
  split.each do |bag_str|
    next if bag_str.include? "no other bag"
    r = bag_str.match(/\d+/)
    bag_count = 0
    bag_count = r[0].to_i if r
    next if bag_count == 0

    bag = bags.get_bag_from_string(bag_str)
    base_bag.add_bag(bag, bag_count)

    puts "#{base_bag} contains #{bag_count} of #{bag}"
  end
end

puts "------------------------------------------------"
my_bag = bags.get_bag("shiny gold")
bags_count = bags.count_bags_inside_bag(my_bag)

puts "------------------------------------------------"
puts "Bags inside my bag: #{bags_count}"
