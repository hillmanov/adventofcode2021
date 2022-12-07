def parse_puzzle(lines):
	# Returns a position as a string where the input maps to the following positions:
	#   #############
	#   #89ABCDEFGHI#
	#   ###1#3#5#7###
	#     #0#2#4#6#
	#     #########
	#
	# Thus the puzzle solution is:
	#   "AABBCCDD..........."
	return (lines[3][1] + lines[2][3] + lines[3][3] + lines[2][5] + lines[3][5] +
	        lines[2][7] + lines[3][7] + lines[2][9] + lines[1][1:-1])

T = parse_puzzle([
"#############"
"#...........#",
"###B#C#B#D###",
"  #A#D#C#A#  ",
"  #########  "
])

print(T)
