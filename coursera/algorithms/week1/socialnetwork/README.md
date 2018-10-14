# Social network connectivity

From [Coursera Algorithms course](https://www.coursera.org/learn/algorithms-part1/quiz/SCxqJ/interview-questions-union-find-ungraded)

## Description

Given a social network containing **n** members and a log file containing **m** timestamps at which times pairs of members formed friendships, design an algorithm to determine the earliest time at which all members are connected (i.e., every member is a friend of a friend of a friend ... of a friend). Assume that the log file is sorted by timestamp and that friendship is an equivalence relation. The running time of your algorithm should be **m logn** or better and use extra space proportional to **n**.

## Example

`go run main.go sn.txt`