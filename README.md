
# Word Counter CLI App
A simple app that can count the number of words from STDIN. This App can also be used to count the number of lines and bytes from STDIN.  
## How to install
### Clone the repository
    git clone https://github.com/dummyheaad/go-cli-word-counter.git
### Build the executable
    go build
## Examples using STDIN
### Word Counter
    echo "lorem ipsum" | ./wc
    Output: 2
### Line Counter
    cat<< EOF | ./wc -l
    lorem
    ipsum
    sit
    EOF
    Output: 3
### Byte Counter
    echo -n "lorem" | ./wc -b
    Output: 5
## Examples using single or multiple files
### Word Counter
    ./wc -file test.txt
    Output:
    Word counting result for each file:
    filename: test.txt, Count: 4
### Line Counter
    ./wc -l -file test.txt
    Output:
    Line counting result for each file:
    filename: test.txt, Count: 2
### Byte Counter
    ./wc -b -file test.txt test1.txt test2.txt
    Output:
    Byte counting result for each file:
    filename: test.txt, Count: 14
    filename: test1.txt, Count: 33
    filename: test2.txt, Count: 11
