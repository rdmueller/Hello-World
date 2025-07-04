CXX = g++
CXXFLAGS = -std=c++11 -Wall -Wextra
TARGET = hello-world-cpp
SOURCE = main.cpp

$(TARGET): $(SOURCE)
	$(CXX) $(CXXFLAGS) -o $(TARGET) $(SOURCE)

clean:
	rm -f $(TARGET)

.PHONY: clean