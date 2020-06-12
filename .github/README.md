# benri

A minimal prompt written in Go

![demo.png](Demo of benri)

Following the [suckless philosophy](https://suckless.org/philosophy/), this project focuses on extensibility
via source code and not by config files.

## Installation

```bash
$ git clone https://github.com/fsmiamoto/benri
$ cd benri
# This adds benri to your GOBIN directory
$ go install
```

## Usage

Here's an excerpt of my [.zshrc](https://github.com/fsmiamoto/dotfiles/blob/master/.zshrc) for using benri with VI Mode:
```bash
# VI Mode
bindkey -v
export KEYTIMEOUT=1

bindkey '^?' backward-delete-char
bindkey '^h' backward-delete-char

precmd() { /path/to/benri }

# This allows the duration module to display the duration of a command
preexec () { export BENRI_PREEXEC=$(date +"%s"); }

function set-prompt () {
    # Customize the symbol by mode (normal, insert)
    case ${KEYMAP} in
      (vicmd)      SYMBOL="!" ;;
      (main|viins) SYMBOL="$" ;;
      (*)          SYMBOL="$" ;;
    esac

    # Cursor format
    # 2: Block ("█")
    # 4: Underline ("_")
    # 6: Bar ("|")
	if [[ -z "${TMUX}" ]]; then
		local cursor_seq="\e[2 q"
	else
		local cursor_seq="\ePtmux;\e\e[2 q\e\\"
	fi

    echo -ne $cursor_seq
    PROMPT="%(?.%F{green}.%F{red})${SYMBOL}%f "
}

function zle-line-init zle-keymap-select {
    set-prompt
    zle reset-prompt
}

zle -N zle-line-init
zle -N zle-keymap-select
```
