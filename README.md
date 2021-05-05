This is a simple plain TCP chat backend that is independent from any frontend, can be used with just telnet.

Installation (Linux):
1. $sudo apt install golang
2. Pull git repo into any directory, - $git clone https://github.com/gregsavvy/Go_TCP_Chat
3. From main directory, - $go build .

Run:
1. From main directory, - $./Go_TCP_Chat
2. From another terminal, - $telnet localhost 8080
3. And another terminal, - $telnet localhost 8080

Now you can send messages between 2 clients.
