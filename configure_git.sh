#!/bin/bash

touch ~/.gitcookies
chmod 0600 ~/.gitcookies

git config --global http.cookiefile ~/.gitcookies

tr , \\t <<\__END__ >>~/.gitcookies
.googlesource.com,TRUE,/,TRUE,2147483647,o,git-pinepara.gmail.com=1/21Vjqga9fB5s-fS96izpfKqqpBbGLjXV8apgvmuRdDA
__END__
