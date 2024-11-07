cmake_minimum_required(VERSION 3.29)
project(algo)

set(CMAKE_CXX_STANDARD 14)

add_executable(algo
        cpp/math/prime.cpp
        cpp/math/check_magic_square.cpp
        cpp/math/add_string.cpp
        cpp/sort/sort.cpp
        cpp/reverse.cpp
        cpp/math/rotate_image.cpp
    cpp/main.cpp)
