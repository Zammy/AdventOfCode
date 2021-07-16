class Group
  def initialize(answers)
    @symbolized_answers = answers.map {|l| l.split("") }
    @symbolized_answers.each do |person_answers|
      person_answers.pop
      person_answers.flatten!
      person_answers.reject! { |s| s.nil? || s.strip.empty? }
      person_answers.map! { |a| a.to_sym }
    end

    @combined_answers = @symbolized_answers[0]
    for i in 1..@symbolized_answers.length-1 do
      @combined_answers = @combined_answers & @symbolized_answers[i]
    end
  end

  def count_answers
    @symbolized_answers.length
  end

  def count_unique_answers
    @combined_answers.length
  end
end

count_answers = 0
answers = []
File.open("input").each do |line|
  unless line.strip.empty?
    answers << line
  else
    group = Group.new(answers)
    count_answers += group.count_unique_answers
    answers.clear()
  end
end

puts count_answers