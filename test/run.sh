#!/bin/bash

echo "Starting test run"

HAS_FAILED=0

function run_test {
  COMMAND=$1
  EXPECTATION_FILENAME=$2
  TESTCASE_NAME=$3

  echo "Test case: $TESTCASE_NAME - STARTED"

  export MAY_BASEPATH=/home
  $COMMAND | sort > actual.txt
  diff actual.txt expectation/$EXPECTATION_FILENAME.txt
  RESULT=$?

  if [ $RESULT -eq 1 ]
  then
    echo "Test case: $TESTCASE_NAME - FAILURE"
    echo "Actual -----------"
    cat actual.txt

    HAS_FAILED=1
  else
    echo "Test case: $TESTCASE_NAME - SUCCESS"
  fi

  rm actual.txt

  return $RESULT
}

run_test "may" "may" "may (show)"
run_test "may -f may" "may_filtered" "may (show, filtered)"
run_test "may -I" "inspect" "may (inspect)"

echo "Finished."
exit $HAS_FAILED
