class Passport
  # byr (Birth Year)
  # iyr (Issue Year)
  # eyr (Expiration Year)
  # hgt (Height)
  # hcl (Hair Color)
  # ecl (Eye Color)
  # pid (Passport ID)
  # cid (Country ID)

  attr_reader :data

  def initialize(dataString)
    @data = {}
    dataString.scan(/(\w+):(#*\w+)/).each do |tuple|
      @data[tuple[0].to_sym] = tuple[1]
    end
  end

  def is_valid?
    validate_birth_year &&
    validate_issue_year &&
    validate_expiration_year &&
    validate_height &&
    validate_hair_color &&
    validate_eye_color &&
    validate_passport_id
  end

  def validate_birth_year
    return false unless @data.has_key? :byr
    birth_year = @data[:byr].to_i
    birth_year >= 1920 && birth_year <= 2002
  end

  def validate_issue_year
    return false unless @data.has_key? :iyr
    issue_year = @data[:iyr].to_i
    issue_year >= 2010 && issue_year <= 2020
  end

  def validate_expiration_year
    return false unless @data.has_key? :eyr
    expiration_year = @data[:eyr].to_i
    expiration_year >= 2020 && expiration_year <= 2030
  end

  def validate_height
    return false unless @data.has_key? :hgt
    height = @data[:hgt]
    result = /(\d+)(cm|in)+/.match(height)
    return false unless result
    height_num = result[1].to_i
    height_measure = result[2]
    return height_num >= 150 && height_num <= 193 if height_measure == "cm"
    return height_num >= 59 && height_num <= 76 if height_measure == "in"
  end

  def validate_hair_color
    return false unless @data.has_key? :hcl
    hair_color = @data[:hcl]
    /#[\dabcdef]{6}/.match(hair_color)
  end

  def validate_eye_color
    return false unless @data.has_key? :ecl
    eye_color = @data[:ecl]
    /(?:amb)|(?:blu)|(?:brn)|(?:gry)|(?:grn)|(?:hzl)|(?:oth)/.match(eye_color)
  end

  def validate_passport_id
    return false unless @data.has_key? :pid
    passport_id = @data[:pid]
    /\d{9}/.match(passport_id)
  end
end

passports = []
passportStr = ""
File.open("input").each do |line|
  passportStr += line
  if line.strip.empty?
    passports << Passport.new(passportStr)
    passportStr = ""
  end
end

total_passports_count = passports.length
valid_passports_count = passports.each.reduce(0) do |valid_count, passport|
  if passport.is_valid?
    valid_count += 1 
    puts passport.data
    puts ""
  end
  valid_count
end
puts "#{valid_passports_count} valid passports of #{total_passports_count}"