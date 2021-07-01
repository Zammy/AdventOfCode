class Passport
  # byr (Birth Year)
  # iyr (Issue Year)
  # eyr (Expiration Year)
  # hgt (Height)
  # hcl (Hair Color)
  # ecl (Eye Color)
  # pid (Passport ID)
  # cid (Country ID)

  def initialize(dataString)
    @data = {}
    dataString.scan(/(\w+):(#*\w+)/).each do |tuple|
      @data[tuple[0].to_sym] = tuple[1]
    end
  end

  def is_valid?
    keys_needed = %i[ecl pid eyr hcl byr iyr hgt]
    keys_needed.all? do |key| @data.has_key? key end
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
  valid_count += 1 if passport.is_valid?
  valid_count
end
puts "#{valid_passports_count} valid passports of #{total_passports_count}"