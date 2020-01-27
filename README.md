# Introduction

One thing that has haunted developers who use multi-threading in their applications is race-conditions. Due to it's randomness, it removes predictability from applications, resulting in unexpected behaviour. But at the same time, generating random numbers has also been a challenge for computers, due to the algorithmic nature of computers.

RCRN is an experimental algorithm to generate true random numbers with race conditions. The main idea is to have 2 seperate threads trying to set a common variable, where each time, they race to set it to their own bit (one thread represents 0 and another represents 1), and the thread that sets it after the other thread wins.

# What is this repo?

This is a implementation of RCRN with Go using GoRoutines.

# Results

The built in `Profile()` method gives out a percentage of near 50%.

# State

RCRN is completely an experiment for now, so it should definitely **not** be used for **any** production applications.