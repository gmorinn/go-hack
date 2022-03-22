##
## EPITECH PROJECT, 2022
## groundhog
## File description:
## Makefile
##

CC = go build

TARGET = go-gm

SRC = 	main.go \
		utils.go \
		brut_force.go

all: $(TARGET)

build_all:
	$(CC) $(CFLAGS) -o $(TARGET) $(OBJ)

$(TARGET): build_all

clean:
	rm -f $(TARGET)

fclean: clean

re: fclean all

.PHONY: all re clean fclean
