# Green Day Bank

A Go command-line banking application based on a programming task from my Hive Helsinki sprint curriculum.

## Objective

The goal of this project is to build a simple banking application that manages user accounts, money transfers, savings, and investments.

## Features

* Four predefined users
* Login, logout, and exit
* Cash balance management
* Savings account with 1% interest
* Deposits and withdrawals
* Transfers between savings and investment accounts
* Money transfers between users
* Three investment funds:

  * `LOW_RISK` — 2%
  * `MEDIUM_RISK` — 5%
  * `HIGH_RISK` — 10%
* Investment appreciation when viewing the balance
* Withdrawal of all investments
* Input validation and insufficient-funds handling
* Graceful EOF handling

## Implementation

The application uses `bufio.Scanner` for user input and handles EOF when `scanner.Scan()` returns `false`.

Financial values are represented using `decimal.Decimal` from the `shopspring/decimal` package.

The application manages users, savings accounts, investment accounts, and funds while maintaining the user's session until logout or exit.

## What I Practiced

* Go CLI applications
* `bufio.Scanner`
* Structs and methods
* Packages and project structure
* Maps and user management
* `decimal.Decimal`
* Account and transaction logic
* State management
* Error handling
* Input validation
* Login and session management
* Working with financial calculations
