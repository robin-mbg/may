#!/bin/bash

if ! [ -x /usr/bin/multitime ] 
then
  echo "Skipping benchmark as multitime is not available."
  exit 0
fi

TEST_COMMAND="may"
TEST_VERSION=`may -V`

TARGET_FILE="benchmarks/results-$TEST_VERSION-$TEST_COMMAND.log"
TEMP_FILE="benchmark_temp.log"

echo "-------------"
echo "Starting test of $TEST_VERSION"
echo "Outputs will be written to $TARGET_FILE"

multitime -n 20 -q $TEST_COMMAND > $TEMP_FILE 2>&1

cat $TEMP_FILE | tail -4 > $TARGET_FILE
rm $TEMP_FILE

echo "Finished testing. Results:"
echo "--------------------------"
cat "$TARGET_FILE"

