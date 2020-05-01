#!/bin/bash

HAS_FAILED=0

function run_snapshot_test {
  COMMAND=$1
  EXPECTATION_FILENAME=$2
  TESTCASE_NAME=$3

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

function run_exit_code_test {
  COMMAND=$1
  MIN_OUTPUT_LENGTH=$2
  TESTCASE_NAME=$3

  $COMMAND > output.txt

  RESULT=$?
  OUTPUT_LENGTH=`cat output.txt | wc -l`

  if [ $RESULT -eq 1 ] || [ $OUTPUT_LENGTH -lt $MIN_OUTPUT_LENGTH ]
  then
    echo "Test case: $TESTCASE_NAME - FAILURE"
    echo "Output -----------"
    cat output.txt

    HAS_FAILED=1
  else
    echo "Test case: $TESTCASE_NAME - SUCCESS"
  fi
}

run_snapshot_test "may" "may" "may (show)"
run_snapshot_test "may -f may" "may_filtered" "may (show, filtered)"
run_snapshot_test "may -I" "inspect" "may (inspect)"

run_exit_code_test "may -U" 12 "may (update)"
run_exit_code_test "may -V" 1 "may (version)"
run_exit_code_test "may -S" 12 "may (status)"
run_exit_code_test "may -v" 15 "may (show, verbose)"

exit $HAS_FAILED
