class PasswordRuler
  def initialize(rule)
    r = rule.match(/^(\d+)-(\d+) (\w+)/)
    if !r 
      throw :not_valid_rule
    end
    @first_index = r[1].to_i-1
    @second_index = r[2].to_i-1
    @char = r[3]
  end

  def check_password(password)
    return false if password[@first_index] == @char && password[@second_index] == @char
    return false if password[@first_index] != @char && password[@second_index] != @char
    return true
  end
end


valid_password = 0
all_passwords = 0;
File.open("input").each do |line|
  split = line.split(':')
  rule = PasswordRuler.new(split[0])
  password = split[1].lstrip
  valid_password +=1 if rule.check_password(password)
  all_passwords +=1
end

puts "valid: #{valid_password} of total #{all_passwords}"
