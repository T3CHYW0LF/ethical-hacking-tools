#!/bin/bash
cat $1 | while read line
do
status_code=$(curl --write-out %{http_code} --silent --output /dev/null $line)
if [[ "$status_code" -ne 200 ]]; then
  echo "CHANGED-$status_code:$line"
else
  echo "ACTIVE-$status_code:$line"
fi
done
