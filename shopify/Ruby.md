## Printing

```ruby
puts 'Hello'
puts 1 # print with \n
print 'World'
print 'World'
```

## Variables and Data Types

```ruby
=begin
mutiple lines comments
mutiple lines comments
mutiple lines comments
=end

# one line comment
# one line comment

name = "Mike"
age = 30
gpa = 3.5
is_good = true

puts "Your name is #{name}"
puts "Your name is " + name
```

## Casting and Converting

```ruby
puts 3.14.to_i
puts 3.to_f
puts "3.0".to_s

puts 100 + "50".to_i
puts 100 + "50.99".to_f
```

## Strings

```ruby
greeting = "HELLO"

puts greeting.length # 5
puts greeting[0] # H
puts greeting.include? "llo" # true
puts greeting.include? "z" # false
puts greeting[1,3] # ell
```
## Numbers

```ruby
puts 2 * 3 
puts 2 ** 3
puts 10 % 3
puts 1 + 2 * 3
puts 10 / 3

num = 10
num += 100
puts num

num = -36.4
puts num.abs() # 36.4
puts num.round() # -36

puts Math.sqrt(144) # 12.0
puts Math.log(0) # -Infinity
```

## User Input

```ruby
puts "Enter your name: "
name = gets                             #.chomp
puts "Hello #{name}, how are you"

puts "Enter first num: "
num1 = gets.chomp  # 10
puts "enter second num: "
num2 = gets.chomp # 20
puts num1.to_f + num2.to_f # 30.0
```

## Arrays

```ruby
lucky_numbers = [4, 8, "fifteen", 16, 23, 42.0]

lucky_numbers[0] = 90
puts lucky_numbers[0] # 90
puts lucky_numbers[1] # 8
puts lucky_numbers[-1] # 42.0

puts lucky_numbers[2,3] # fifteen 16 23

puts lucky_numbers[2..4] # fifteen 16 23

puts lucky_numbers.length # 6
```

## N Dimensional Arrays

```ruby
#number_grid = [[],[]]
number_grid = [ [1, 2], [3, 4] ]
number_grid[0][0] = 99

puts number_grid[0][0] # 99
puts number_grid[0][1] # 2
```

## Array Methods

```ruby

friends = []
friends.push("Oscar")
friends.push("Angela")
friends.push("Kevin")

# friends.pop
puts  friends
puts "\n"

puts friends.reverse
puts "\n"

puts friends.sort
puts "\n"

puts  friends.include? "Oscar"

=begin
Oscar
Angela
Kevin

Kevin
Angela
Oscar

Angela
Kevin
Oscar

true
=end
```

## Methods

```ruby
def add_numbers(num1, num2=99)
     return num1 + num2
end

sum = add_numbers(4, 3)
puts sum
```

## If Statements

```ruby
is_student = true
is_smart = true

if is_student and is_smart # or
	puts "You are a student"
elsif is_student and !is_smart
	puts "You are not a smart student"
else
	puts "You are not a student and not smart"
end

# >, <, >=, <=, !=, ==, String.equals()
if 1 > 3
	puts "number comparison was true"
end

if "c" > "b"
     puts "string comparison was true"
end
```

## Switch Statements

```ruby

my_grade = "A"
case my_grade
     when "A"
		puts "You Pass"
     when "F"
     	puts "You fail"
     else
     	puts "Invalid grade"
end
```

## Dictionaries

```ruby

test_grades = {
    "Andy" => "B+",
    :Stan_ley => "C",
    "Ryan" => "A",
    3 => 95.2
}

test_grades["Andy"] = "B-"
puts test_grades["Andy"] # B-
puts test_grades[:Stan_ley] # C
puts test_grades[3] # 95.2
```

## While Loops

```ruby
index = 1
while index <= 5
	puts index
	index += 1
end
```

## For Loops

```ruby

for index in 0..5
    puts index
end

5.times do |index|
    puts index
end

lucky_nums = [4, 8, 15, 16, 23, 42]
for lucky_num in lucky_nums
    puts lucky_num
end

lucky_nums.each do |lucky_num|
     puts lucky_num
end

```

## Exception Catching

```ruby
begin
     # puts bad_variable
     num = 10/0
rescue ZeroDivisionError
     puts "Error"
rescue
     puts "All other errors"
end

raise "Made Up Exception"
```

## Classes and Objects

```ruby

```
