NAME=ct


all:	
	go build -o $(NAME)
	./$(NAME) $(LOG)

build:
	go build -o $(NAME)


