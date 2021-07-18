class Operation
  attr_accessor :type, :value

  def initialize(type, value)
    @type = type
    @value = value
  end

  def to_s
    "#{type} #{value}"
  end
end

class Procedure
  attr_accessor :acc, :execution_pointer

  def initialize(procedure)
    @procedure = procedure
    @execution_pointer = 0
    @operations_executed = []
    @acc = 0
  end

  def execute
    if @operations_executed.any? @execution_pointer
      puts "Detected infinite loop starting from operation at index:#{execution_pointer}"
      return
    end
    @operations_executed << @execution_pointer
    operation = @procedure[@execution_pointer]
    if operation.nil?
      throw :out_of_bounds
      return
    end
    puts "Executing operation ##{execution_pointer}"
    case operation.type
    when "acc"
      @acc += operation.value
      @execution_pointer += 1
      puts "  Increased accumulator to #{@acc}"
    when "jmp"
      @execution_pointer += operation.value
      puts "  Jumped to ##{@execution_pointer}"
    when "nop"
      #does nothing
      puts "  -"
      @execution_pointer += 1
    else
      throw :this_should_not_happen
    end

    execute()
  end
end

boot_procedure = []
File.open("input").each do |line|
  line.scan(/(\w{3}) ([+|-]\d+)/) do |r|
    boot_procedure << Operation.new(r[0], r[1].to_i)
  end
end

boot = Procedure.new(boot_procedure)
boot.execute

puts "--------------------------------------------------------"
puts "Acc is #{boot.acc} after detecting infinite loop on operation starting from #{boot.execution_pointer} index"
