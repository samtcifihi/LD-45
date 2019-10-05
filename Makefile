compile:
	go build -o "Nothing_to_Pebbles"

test: compile
	./Nothing_to_Pebbles

clean:
	rm Nothing_to_Pebbles