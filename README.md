# Bus Ticket Puzzle

When you get into a bus and pay your fare, you receive a ticket with an arbitrary 6-digit number. Your trip may be quite long, so you can use your time for some mental exercise. Your goal is to put certain mathematical operations and brackets among the digits of your ticket, so that you get a result of 100.

### Rules

- All digits must be presented in their original order, none added, none removed.
- You can join digits into several-digit numbers.
- You can only put parentheses (, ) and arithmetic operations (set of allowed operations can vary; addition +, subtraction -, multiplication *, division /, exponentiation ^, factorial !, square root, unary minus).
- The result must be a valid arithmetical expression, which can be successfully calculated. 

### Examples

	000004  (0!+0!+0!+0!)*(0!+4!)   == 100
	458392  (sqrt(4)*5)!*(8-3)/9!*2 == 100
	957129  (-sqrt(9)+5)^7)+1-29    == 100


### Results

Valid arithmetical expression found for 999858 tickets. Only 142 tickets remains [unsolved](./unsolved.txt)