package help

func Reverse(input string) string {
	output:=""
	for i:=len(input)-1;i>=0;i--{
		output+=string(input[i])
	}
	return output
}