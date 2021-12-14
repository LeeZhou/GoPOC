Common Sequence will return the first 100 repeatable three word sentence. 

To Build Application

go build -o commonsequence

To Run Application 

You can pipe file output to program
cat filename| ./commonsequence

To prase one or more file
.\commonsequence filename

Or both Piped and file arguments
cat filename1|.\commonsequence filename2 filename3 filename4 

Test files can be found in test folder
test1: 6 phrases
test2: 100 phrase (131 actual phrases total), where first 6 are greater than 1
test3: 100 phrases (194922 actual phrases total), where all values are greater than 14

To build docker immage, use:
docker build  --tag commonsequence .

Currently application does not support unicode characters (characters greater than a byte) 

ToDo
- To add argument flag for Stat Size (default 100)
- To add argument flag for phrase size (default 3)
- Build merge sort function (desc) as current sort implementation is 20 
- Automated test
- Docker container

