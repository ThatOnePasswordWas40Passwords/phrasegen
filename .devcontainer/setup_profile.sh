#!/bin/bash

cat << 'EOF' >> ~/.bashrc
alias ls='ls --color -h --group-directories-first'

# Colors
black="\[$(tput setaf 0)\]"
red="\[$(tput setaf 1)\]"
green="\[$(tput setaf 2)\]"
yellow="\[$(tput setaf 3)\]"
blue="\[$(tput setaf 4)\]"
magenta="\[$(tput setaf 5)\]"
cyan="\[$(tput setaf 6)\]"
white="\[$(tput setaf 7)\]"

# Title bar - "user@host: ~"
title="\u@\h: \w"
titlebar="\[\033]0;"$title"\007\]"

# Git branch
git_branch() {
  git branch 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/(\1)\ /';
}

# Clear attributes
clear_attributes="\[$(tput sgr0)\]"

# Custom bash prompt - "➜  ~ (master) "
export PS1="${titlebar}${green}➜  ${blue}\W ${cyan}\$(git_branch)${clear_attributes}"

export CLICOLOR=1
alias lsdir='ls -F | grep /'
export LSCOLORS=ExFxBxDxCxegedabagacad
EOF