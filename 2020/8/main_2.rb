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

  def initialize(operations)
    @procedure = operations
    @execution_pointer = 0
    @call_stack = []
    @acc = 0
    @modified_operation_at_index = -1
  end

  def execute
    if @call_stack.any? @execution_pointer
      puts
      puts "Detected infinite loop at ##{execution_pointer}"
      puts

      if @modified_operation_at_index != -1
        modify_at = find_operation_index_of_jmp_or_nop_before_operation_at(@modified_operation_at_index)
      else
        modify_at = find_operation_index_of_jmp_or_nop_before_operation_at(@call_stack[-2])
      end

      unwind_call_stack_to(modify_at)
      @modified_operation_at_index = modify_at
      puts
    end

    if @execution_pointer == @procedure.length
      puts "Reached end of procedure"
      return
    end

    @call_stack << @execution_pointer

    operation = @procedure[@execution_pointer]
    if operation.nil?
      throw :out_of_bounds
      return
    end
    execute_operation_at(@execution_pointer)

    execute()
  end

  private

  def find_operation_index_of_jmp_or_nop_before_operation_at(operation_index)
    reached_index = false
    @call_stack.reverse_each do |index|
      if reached_index
        if @procedure[index].type == "jmp" || @procedure[index].type == "nop"
          return index
        end
      elsif index == operation_index
        reached_index = true
      end
    end

    throw :should_never_happen
  end

  def execute_operation_at(index)
    puts "Executing operation ##{index}"
    operation = @procedure[index]

    operation_type = operation.type
    if @modified_operation_at_index == index
      if operation_type == "jmp"
        operation_type = "nop"
      elsif operation_type == "nop"
        operation_type = "jmp"
      end
    end

    case operation_type
    when "acc"
      prev_value_of_acc = @acc
      @acc += operation.value
      @execution_pointer += 1
      puts "  Increased accumulator from #{prev_value_of_acc} to #{@acc}"
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
  end

  def reverse_operation_at(index)
    puts "Reversing operation ##{index}"
    operation = @procedure[index]
    if operation.type == "acc"
      prev_value = @acc
      @acc -= operation.value
      puts "  Decreased accumulator to #{@acc} from #{prev_value}"
    else
      puts "  -"
    end
  end

  def unwind_call_stack_to(unwind_index)
    puts "Unwinding to ##{unwind_index} <<<<<<<<<"
    loop do
      index = @call_stack.pop
      reverse_operation_at(index)
      @execution_pointer = index
      break if index == unwind_index
    end
    puts ">>>>>>>>>>>>>>>>>>>>>>>>>>>"
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
puts "Acc is #{boot.acc} after finisning procedure"
