#!/bin/bash

NODE_ENV=production npx tailwindcss-cli@latest build ./public/main.css -o ./public/main.out.css
