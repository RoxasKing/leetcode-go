package main

/*
  Suppose we have a file system that stores both files and directories. An example of one system is represented in the following picture:

  Here, we have dir as the only directory in the root. dir contains two subdirectories, subdir1 and subdir2. subdir1 contains a file file1.ext and subdirectory subsubdir1. subdir2 contains a subdirectory subsubdir2, which contains a file file2.ext.

  In text form, it looks like this (with ⟶ representing the tab character):

    dir
    ⟶ subdir1
    ⟶ ⟶ file1.ext
    ⟶ ⟶ subsubdir1
    ⟶ subdir2
    ⟶ ⟶ subsubdir2
    ⟶ ⟶ ⟶ file2.ext

  If we were to write this representation in code, it will look like this: "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext". Note that the '\n' and '\t' are the new-line and tab characters.

  Every file and directory has a unique absolute path in the file system, which is the order of directories that must be opened to reach the file/directory itself, all concatenated by '/'s. Using the above example, the absolute path to file2.ext is "dir/subdir2/subsubdir2/file2.ext". Each directory name consists of letters, digits, and/or spaces. Each file name is of the form name.extension, where name and extension consist of letters, digits, and/or spaces.

  Given a string input representing the file system in the explained format, return the length of the longest absolute path to a file in the abstracted file system. If there is no file in the system, return 0.

  Example 1:
    Input: input = "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
    Output: 20
    Explanation: We have only one file, and the absolute path is "dir/subdir2/file.ext" of length 20.

  Example 2:
    Input: input = "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
    Output: 32
    Explanation: We have two files:
      "dir/subdir1/file1.ext" of length 21
      "dir/subdir2/subsubdir2/file2.ext" of length 32.
      We return 32 since it is the longest absolute path to a file.

  Example 3:
    Input: input = "a"
    Output: 0
    Explanation: We do not have any files, just a single directory named "a".

  Example 4:
    Input: input = "file1.txt\nfile2.txt\nlongfile.txt"
    Output: 12
    Explanation: There are 3 files at the root directory.
      Since the absolute path for anything at the root directory is just the name itself, the answer is "longfile.txt" with length 12.

  Constraints:
    1. 1 <= input.length <= 10^4
    2. input may contain lowercase or uppercase English letters, a new line character '\n', a tab character '\t', a dot '.', a space ' ', and digits.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-absolute-file-path
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack

func lengthLongestPath(input string) int {
	n := len(input)
	stk1 := IntStack{} // storage parent dir path length
	stk2 := IntStack{} // storage full dir path length

	out := 0
	for l, r := 0, 0; r < n; l, r = r+1, r+1 {
		for ; r < n && input[r] == '\t'; r++ {
		}
		for stk1.Len() > r-l {
			stk1.Pop()
			stk2.Pop()
		}

		isFile := false
		for l = r; r < n && input[r] != '\n'; r++ {
			isFile = isFile || input[r] == '.'
		}

		if stk1.Len() > 0 {
			stk1.Push(r - l + 1)
		} else {
			stk1.Push(r - l) // root path, not need to add '/'
		}
		if stk2.Len() > 0 {
			stk2.Push(stk2.Top() + stk1.Top()) // acumulate path length
		} else {
			stk2.Push(stk1.Top())
		}

		if isFile && stk2.Top() > out {
			out = stk2.Top()
		}
	}

	return out
}

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s IntStack) Top() int    { return s[s.Len()-1] }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
