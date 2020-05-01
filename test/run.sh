#!/bin/bash

export MAY_BASEPATH=/home
may | sort > actual.txt
diff actual.txt expectation/may.txt
RESULT=$?

if [ $RESULT -eq 1 ]
then
  echo "Actual -----------"
  cat actual.txt
else
  echo "Test case successful"
fi

exit $RESULT
