#!/bin/bash

echo "-------------------------------"
echo "Starting integration test suite"
echo "-------------------------------"

./test.sh
TEST_RESULT=$?

./benchmark.sh

echo "--------"
echo "Finished."

exit $TEST_RESULT
