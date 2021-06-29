class PasswordRuler
  def initialize(rule)
    r = rule.match(/^(\d+)-(\d+) (\w+)/)
    if !r 
      throw :not_valid_rule
    end
    @at_least = r[1].to_i
    @at_most = r[2].to_i
    @char = r[3]
  end

  def check_password(password)
    char_count = password.each_char.reduce(0) do |count, char|
      count += 1 if char == @char
      count
    end
    
    return char_count >= @at_least && char_count <= @at_most
  end
end


valid_password = 0
all_passwords = 0;
File.open("input").each do |line|
  split = line.split(':')
  rule = PasswordRuler.new(split[0])
  password = split[1]
  valid_password +=1 if rule.check_password(password)
  all_passwords +=1
end

puts "valid: #{valid_password} of total #{all_passwords}"
