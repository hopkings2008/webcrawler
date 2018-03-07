.PHONY=all clean
target=webcrawler
gflags=-gcflags "-N -l"

all:
	go build -o ${target} ${gflags}

clean:
	rm -f ${target}
