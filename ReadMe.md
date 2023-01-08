### Objectives

This project is meant to solve and build a digital version of an ant farm.

The program read from a file provided that contains information about the number of ants, the rooms details (including their coordinates on the x and y plane), path connections. If the number of ant provided is greater than 0, it will solve and print the ant farm.
-   At the beginning of the game, all the ants are in the room  `##start`. The goal is to bring them to the room  `##end`  with as few moves as possible.

If you want to see a minimalistic movement of the ant, use the flag  `-v` to see how the ants move and the number of steps and how long it will take the ants to go from start room to end room
-   The shortest path is not necessarily the simplest.
- We only use the path that uses the minimum number of steps to moves the ant from start to end. If more than one path returns the same steps, then we use the path with minimum number of moves. If there are more than one path still, we can use any of these paths.

The project is written in Golang as part of 01 edu school projects for Gritlab Coding School, Åland.

**Oluwatosin I**
