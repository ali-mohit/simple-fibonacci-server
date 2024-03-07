#!/bin/bash

# If any command is provided, run that
if [[ $# -gt 0 ]]; then
  /bin/bash -c "$*"
# else start http server
else
  echo "starting http server"
  if ! ./fibo_app serve; then
    echo "failed to start http server"
  fi
fi