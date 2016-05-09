# dobuild Specification

## Abstract

Ever tried to build a C/C++ program? Maybe a hello world? Maybe something more complex? It's complex. There are hundreds of existing compiler flags, dependency tracking, etc etc etc. This is a build system (Similar to CMake) that can generate UNIX Makefiles (and can totally generate anything else of you build it)

## What's wrong with current tools?

CMake is the current standard for build systems, but I don't like it, for all the following reasons:

* Variables and arrays are treated as the same thing, which leads to confusion
* It can be hard just to setup easy scripts to import a "target" into your program/library
* Very bloated
* Doesn't create standalone Makefiles
* Doens't have similar syntax to C, which is pointless: we are C/C++ devs.

This project tries to solve all of these problems.

## Tutorial 

## Specification

### Statement seperation
Each statemant must be seperated with a semicolon, which leads to more legible code, and avoids some weird compile errors.

### Variables

Variables store a single value that can be referred to later. The general syntax for declaring a variable is

```
var <variable_name> []<variable_type>] [ = <init_value>]
```

dobuild is optionally typed, which means that you can declare the type if you want to, and it is generally considered good practice to. But when writing generic-ish code, it is important to have the option to leave it out.

#### Build-in types

Any language needs to have some primitive types that all else is based off of.

##### `int`

A `int` is an integer, specifically 64-bit signed integer.

EXAMPLES:

```
var a = 3;
var b int = 5;
```

In order for a literal to be an int, it has to be just numbers, and no decimal point.

##### `float`

A `float` is a basic 64-bit floating point number.

There are all the basic operators on strings: addition (`+`), subtraction (`-`), multiplication (`*`), division (`/`), power (`^`).

EXAMPLES:
```
var a = 3.0;
var b float = 3.1;
```

To define a float literal, you must have a number, then a decimal point, _then at least one more digit_

##### `char`

The char represents a UTF-8 encoded character. This is the basis of strings, etc.

### Arrays

One major failure of CMake is their arrays: it is sometimes unclear what they are, where each element starts and ends, etc.

You can define an array of any given type with the syntax:

```
<type>[]
```

So to define an array of integers, you can use `var foobar int[]`. This will initialze an empty array.

#### `append`


#### `array_length`

A built-in function that gets the length of an array. 

EXAMPLE:

