# turing-bot

A discord bot helps groups with coding practicing and note taking.

## Dependency
- [Go 1.12.1+](https://golang.org/dl/)
- [sqlite](https://www.sqlite.org/download.html)

## DB Schema

**user**(uid, dcid, uname, fname, lname, createts)

**problem**(pid, pname)

**user_problem**(upid, uid, pid, ts, note)

## Usage

```
!solved <pname> [-m <msg>]            Record the problem solved with the option to take notes
!show <uname> [-a]                    Show the problems solved today by given username 
                                      and show all entries when -a specified
TODO:
!count <uname> [-a]
!noteshow <uname> [-p <pname>]
!graph <uname>
!help <pname>
!today
!libraryhours

e.g. !solved LintCode18
     !solved LeetCode200 -m "This is a comment"
     !show Richard -a
```
