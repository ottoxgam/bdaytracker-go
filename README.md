# BDayTracker

![Linter passing status](https://github.com/LoDThe/bdaytracker-go/workflows/Golang%20CI%20Lint/badge.svg)
![Test passing status](https://github.com/LoDThe/bdaytracker-go/workflows/Test/badge.svg)

## Description

This is a Telegram bot that helps to not miss your friends' birthdays. 

Add a friend's date of birth, and the bot will remind you about the birthday on the right day.

## Features

Add friends manually or link them by Telegram @username. Birthday reminders are sent automatically on the right day.

## How to find the bot

The bot is available by the following link: [@bdaytracker_bot](https://t.me/bdaytracker_bot)

The bot UI is in English.

## Technicalities

- CI consists of running tests and the linter.

- PostgreSQL is used to store user states. [Migrations](./migrations) are applied automatically when the app starts, so the application can be started with an empty database. 

- The application can be run with [docker-compose](./docker-compose.yml). Check [.env.dist](.env.dist) as an example for envs. 