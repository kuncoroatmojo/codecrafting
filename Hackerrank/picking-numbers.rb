#!/bin/ruby
#https://www.hackerrank.com/challenges/picking-numbers
def pickingNumbers(a)
    max = 0
    a.each_with_index do |val, i|
        sum = 0
        sumup = 0
        sumdown = 0
        a.each_with_index do |valx, j|
            if valx - val == 1
                sumup += 1
            elsif valx - val == -1
                sumdown += 1
            elsif valx == val
                sum += 1
            end
        end
        if sumup > sumdown
            sum += sumup
        else
            sum += sumdown
        end
        if sum > max
            max = sum
        end
    end
    return max
end

n = gets.strip.to_i
a = gets.strip
a = a.split(' ').map(&:to_i)
result = pickingNumbers(a)
puts result
