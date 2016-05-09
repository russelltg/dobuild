cmake_minimum_required(VERSION 2.6)
project(dobuild)

add_executable(dobuild main.cpp)

install(TARGETS dobuild RUNTIME DESTINATION bin)
