class Group
  def initialize(answers)
    @symbolized_answers = answers
      .map { |l| l.split("") }
      .flatten
      .reject { |s| s.nil? || s.strip.empty? }
      .uniq
      .map { |a| a.to_sym }
  end

  def count_answers
    @symbolized_answers.length
  end
end

count_answers = 0
answers = []
File.open("input").each do |line|
  unless line.strip.empty?
    answers << line
  else
    group = Group.new(answers)
    count_answers += group.count_answers
    answers.clear()
  end
end

puts count_answers