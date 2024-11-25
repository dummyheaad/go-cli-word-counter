
# Word Counter CLI App
A simple app that can count the number of words from STDIN. This App can also be used to count the number of lines and bytes from STDIN.  
## How to install
### Clone the repository
    git clone https://github.com/dummyheaad/go-cli-word-counter.git
### Build the executable
    go build
## Examples
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
